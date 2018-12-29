// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "reflect"

// FieldLevel is an autogenerated mock type for the FieldLevel type
type FieldLevel struct {
	mock.Mock
}

// ExtractType provides a mock function with given fields: field
func (_m *FieldLevel) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	ret := _m.Called(field)

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func(reflect.Value) reflect.Value); ok {
		r0 = rf(field)
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	var r1 reflect.Kind
	if rf, ok := ret.Get(1).(func(reflect.Value) reflect.Kind); ok {
		r1 = rf(field)
	} else {
		r1 = ret.Get(1).(reflect.Kind)
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func(reflect.Value) bool); ok {
		r2 = rf(field)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Field provides a mock function with given fields:
func (_m *FieldLevel) Field() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// FieldName provides a mock function with given fields:
func (_m *FieldLevel) FieldName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStructFieldOK provides a mock function with given fields:
func (_m *FieldLevel) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	var r1 reflect.Kind
	if rf, ok := ret.Get(1).(func() reflect.Kind); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(reflect.Kind)
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func() bool); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Param provides a mock function with given fields:
func (_m *FieldLevel) Param() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Parent provides a mock function with given fields:
func (_m *FieldLevel) Parent() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}

// StructFieldName provides a mock function with given fields:
func (_m *FieldLevel) StructFieldName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Top provides a mock function with given fields:
func (_m *FieldLevel) Top() reflect.Value {
	ret := _m.Called()

	var r0 reflect.Value
	if rf, ok := ret.Get(0).(func() reflect.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Value)
	}

	return r0
}
