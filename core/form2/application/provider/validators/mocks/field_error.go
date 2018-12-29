// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import reflect "reflect"
import ut "github.com/go-playground/universal-translator"

// FieldError is an autogenerated mock type for the FieldError type
type FieldError struct {
	mock.Mock
}

// ActualTag provides a mock function with given fields:
func (_m *FieldError) ActualTag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Field provides a mock function with given fields:
func (_m *FieldError) Field() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Kind provides a mock function with given fields:
func (_m *FieldError) Kind() reflect.Kind {
	ret := _m.Called()

	var r0 reflect.Kind
	if rf, ok := ret.Get(0).(func() reflect.Kind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reflect.Kind)
	}

	return r0
}

// Namespace provides a mock function with given fields:
func (_m *FieldError) Namespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Param provides a mock function with given fields:
func (_m *FieldError) Param() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StructField provides a mock function with given fields:
func (_m *FieldError) StructField() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StructNamespace provides a mock function with given fields:
func (_m *FieldError) StructNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Tag provides a mock function with given fields:
func (_m *FieldError) Tag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Translate provides a mock function with given fields: _a0
func (_m *FieldError) Translate(_a0 ut.Translator) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(ut.Translator) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *FieldError) Type() reflect.Type {
	ret := _m.Called()

	var r0 reflect.Type
	if rf, ok := ret.Get(0).(func() reflect.Type); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.Type)
		}
	}

	return r0
}

// Value provides a mock function with given fields:
func (_m *FieldError) Value() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}
