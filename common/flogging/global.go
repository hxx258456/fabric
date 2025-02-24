/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flogging

import (
	"io"

	"github.com/hxx258456/ccgo/grpc/grpclog"
	"go.uber.org/zap/zapcore"
)

const (
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} %{id:04x} %{level:.4s}%{color:reset} [%{module}] %{color:bold}%{shortfunc}%{color:reset} -> %{message}"
	defaultLevel  = zapcore.InfoLevel
)

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
	grpcLogger := Global.ZapLogger("grpc")
	grpclog.SetLogger(NewGRPCLogger(grpcLogger))
}

// Init initializes logging with the provided config.
func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

// Reset sets logging to the defaults defined in this package.
//
// Used in tests and in the package init
func Reset() {
	Global.Apply(Config{})
}

// LoggerLevel gets the current logging level for the logger with the
// provided name.
func LoggerLevel(loggerName string) string {
	return Global.Level(loggerName).String()
}

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *FabricLogger {
	return Global.Logger(loggerName)
}

// ActivateSpec is used to activate a logging specification.
func ActivateSpec(spec string) {
	err := Global.ActivateSpec(spec)
	if err != nil {
		panic(err)
	}
}

// DefaultLevel returns the default log level.
func DefaultLevel() string {
	return defaultLevel.String()
}

// SetWriter calls SetWriter returning the previous value
// of the writer.
func SetWriter(w io.Writer) io.Writer {
	return Global.SetWriter(w)
}

// SetObserver calls SetObserver returning the previous value
// of the observer.
func SetObserver(observer Observer) Observer {
	return Global.SetObserver(observer)
}
