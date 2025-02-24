// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	metrics "github.com/hxx258456/fabric/common/metrics"
	mock "github.com/stretchr/testify/mock"
)

// MetricsProvider is an autogenerated mock type for the MetricsProvider type
type MetricsProvider struct {
	mock.Mock
}

// NewCounter provides a mock function with given fields: opts
func (_m *MetricsProvider) NewCounter(opts metrics.CounterOpts) metrics.Counter {
	ret := _m.Called(opts)

	var r0 metrics.Counter
	if rf, ok := ret.Get(0).(func(metrics.CounterOpts) metrics.Counter); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Counter)
		}
	}

	return r0
}

// NewGauge provides a mock function with given fields: opts
func (_m *MetricsProvider) NewGauge(opts metrics.GaugeOpts) metrics.Gauge {
	ret := _m.Called(opts)

	var r0 metrics.Gauge
	if rf, ok := ret.Get(0).(func(metrics.GaugeOpts) metrics.Gauge); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Gauge)
		}
	}

	return r0
}

// NewHistogram provides a mock function with given fields: opts
func (_m *MetricsProvider) NewHistogram(opts metrics.HistogramOpts) metrics.Histogram {
	ret := _m.Called(opts)

	var r0 metrics.Histogram
	if rf, ok := ret.Get(0).(func(metrics.HistogramOpts) metrics.Histogram); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metrics.Histogram)
		}
	}

	return r0
}
