/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoutil

import (
	"bytes"
	"crypto/sha256"
	"encoding/asn1"
	"math/big"

	"github.com/golang/protobuf/proto"
	cb "github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/pkg/errors"
)

// NewBlock constructs a block with no data and no metadata.
func NewBlock(seqNum uint64, previousHash []byte) *cb.Block {
	block := &cb.Block{}
	block.Header = &cb.BlockHeader{}
	block.Header.Number = seqNum
	block.Header.PreviousHash = previousHash
	block.Header.DataHash = []byte{}
	block.Data = &cb.BlockData{}

	var metadataContents [][]byte
	for i := 0; i < len(cb.BlockMetadataIndex_name); i++ {
		metadataContents = append(metadataContents, []byte{})
	}
	block.Metadata = &cb.BlockMetadata{Metadata: metadataContents}

	return block
}

type asn1Header struct {
	Number       *big.Int
	PreviousHash []byte
	DataHash     []byte
}

func BlockHeaderBytes(b *cb.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
		Number:       new(big.Int).SetUint64(b.Number),
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		// Errors should only arise for types which cannot be encoded, since the
		// BlockHeader type is known a-priori to contain only encodable types, an
		// error here is fatal and should not be propagated
		panic(err)
	}
	return result
}

func BlockHeaderHash(b *cb.BlockHeader) []byte {
	sum := sha256.Sum256(BlockHeaderBytes(b))
	return sum[:]
}

func BlockDataHash(b *cb.BlockData) []byte {
	sum := sha256.Sum256(bytes.Join(b.Data, nil))
	return sum[:]
}

// GetChannelIDFromBlockBytes returns channel ID given byte array which represents
// the block
func GetChannelIDFromBlockBytes(bytes []byte) (string, error) {
	block, err := UnmarshalBlock(bytes)
	if err != nil {
		return "", err
	}

	return GetChannelIDFromBlock(block)
}

// GetChannelIDFromBlock returns channel ID in the block
func GetChannelIDFromBlock(block *cb.Block) (string, error) {
	if block == nil || block.Data == nil || block.Data.Data == nil || len(block.Data.Data) == 0 {
		return "", errors.New("failed to retrieve channel id - block is empty")
	}
	var err error
	envelope, err := GetEnvelopeFromBlock(block.Data.Data[0])
	if err != nil {
		return "", err
	}
	payload, err := UnmarshalPayload(envelope.Payload)
	if err != nil {
		return "", err
	}

	if payload.Header == nil {
		return "", errors.New("failed to retrieve channel id - payload header is empty")
	}
	chdr, err := UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return "", err
	}

	return chdr.ChannelId, nil
}

// GetMetadataFromBlock retrieves metadata at the specified index.
func GetMetadataFromBlock(block *cb.Block, index cb.BlockMetadataIndex) (*cb.Metadata, error) {
	if block.Metadata == nil {
		return nil, errors.New("no metadata in block")
	}

	if len(block.Metadata.Metadata) <= int(index) {
		return nil, errors.Errorf("no metadata at index [%s]", index)
	}

	md := &cb.Metadata{}
	err := proto.Unmarshal(block.Metadata.Metadata[index], md)
	if err != nil {
		return nil, errors.Wrapf(err, "error unmarshalling metadata at index [%s]", index)
	}
	return md, nil
}

// GetMetadataFromBlockOrPanic retrieves metadata at the specified index, or
// panics on error
func GetMetadataFromBlockOrPanic(block *cb.Block, index cb.BlockMetadataIndex) *cb.Metadata {
	md, err := GetMetadataFromBlock(block, index)
	if err != nil {
		panic(err)
	}
	return md
}

// GetConsenterMetadataFromBlock attempts to retrieve consenter metadata from the value
// stored in block metadata at index SIGNATURES (first field). If no consenter metadata
// is found there, it falls back to index ORDERER (third field).
func GetConsenterMetadataFromBlock(block *cb.Block) (*cb.Metadata, error) {
	m, err := GetMetadataFromBlock(block, cb.BlockMetadataIndex_SIGNATURES)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to retrieve metadata")
	}

	// TODO FAB-15864 Remove this fallback when we can stop supporting upgrade from pre-1.4.1 orderer
	if len(m.Value) == 0 {
		return GetMetadataFromBlock(block, cb.BlockMetadataIndex_ORDERER)
	}

	obm := &cb.OrdererBlockMetadata{}
	err = proto.Unmarshal(m.Value, obm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal orderer block metadata")
	}

	res := &cb.Metadata{}
	err = proto.Unmarshal(obm.ConsenterMetadata, res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal consenter metadata")
	}

	return res, nil
}

// GetLastConfigIndexFromBlock retrieves the index of the last config block as
// encoded in the block metadata
func GetLastConfigIndexFromBlock(block *cb.Block) (uint64, error) {
	m, err := GetMetadataFromBlock(block, cb.BlockMetadataIndex_SIGNATURES)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to retrieve metadata")
	}
	// TODO FAB-15864 Remove this fallback when we can stop supporting upgrade from pre-1.4.1 orderer
	if len(m.Value) == 0 {
		m, err := GetMetadataFromBlock(block, cb.BlockMetadataIndex_LAST_CONFIG)
		if err != nil {
			return 0, errors.WithMessage(err, "failed to retrieve metadata")
		}
		lc := &cb.LastConfig{}
		err = proto.Unmarshal(m.Value, lc)
		if err != nil {
			return 0, errors.Wrap(err, "error unmarshalling LastConfig")
		}
		return lc.Index, nil
	}

	obm := &cb.OrdererBlockMetadata{}
	err = proto.Unmarshal(m.Value, obm)
	if err != nil {
		return 0, errors.Wrap(err, "failed to unmarshal orderer block metadata")
	}
	return obm.LastConfig.Index, nil
}

// GetLastConfigIndexFromBlockOrPanic retrieves the index of the last config
// block as encoded in the block metadata, or panics on error
func GetLastConfigIndexFromBlockOrPanic(block *cb.Block) uint64 {
	index, err := GetLastConfigIndexFromBlock(block)
	if err != nil {
		panic(err)
	}
	return index
}

// CopyBlockMetadata copies metadata from one block into another
func CopyBlockMetadata(src *cb.Block, dst *cb.Block) {
	dst.Metadata = src.Metadata
	// Once copied initialize with rest of the
	// required metadata positions.
	InitBlockMetadata(dst)
}

// InitBlockMetadata initializes metadata structure
func InitBlockMetadata(block *cb.Block) {
	if block.Metadata == nil {
		block.Metadata = &cb.BlockMetadata{Metadata: [][]byte{{}, {}, {}, {}, {}}}
	} else if len(block.Metadata.Metadata) < int(cb.BlockMetadataIndex_COMMIT_HASH+1) {
		for i := int(len(block.Metadata.Metadata)); i <= int(cb.BlockMetadataIndex_COMMIT_HASH); i++ {
			block.Metadata.Metadata = append(block.Metadata.Metadata, []byte{})
		}
	}
}
