// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	ocrkey "github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocrkey"
	mock "github.com/stretchr/testify/mock"
)

// OCR is an autogenerated mock type for the OCR type
type OCR struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *OCR) Add(key ocrkey.KeyV2) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(ocrkey.KeyV2) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *OCR) Create() (ocrkey.KeyV2, error) {
	ret := _m.Called()

	var r0 ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func() ocrkey.KeyV2); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ocrkey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *OCR) Delete(id string) (ocrkey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func(string) ocrkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ocrkey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *OCR) EnsureKey() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *OCR) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *OCR) Get(id string) (ocrkey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func(string) ocrkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ocrkey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *OCR) GetAll() ([]ocrkey.KeyV2, error) {
	ret := _m.Called()

	var r0 []ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func() []ocrkey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocrkey.KeyV2)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetV1KeysAsV2 provides a mock function with given fields:
func (_m *OCR) GetV1KeysAsV2() ([]ocrkey.KeyV2, error) {
	ret := _m.Called()

	var r0 []ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func() []ocrkey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocrkey.KeyV2)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *OCR) Import(keyJSON []byte, password string) (ocrkey.KeyV2, error) {
	ret := _m.Called(keyJSON, password)

	var r0 ocrkey.KeyV2
	if rf, ok := ret.Get(0).(func([]byte, string) ocrkey.KeyV2); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(ocrkey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
