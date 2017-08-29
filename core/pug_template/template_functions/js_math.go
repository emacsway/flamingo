package template_functions

import (
	"math"
	"reflect"
)

type (
	// JsMath is exported as a template function
	JsMath struct{}

	// Math is our Javascript's Math equivalent
	Math struct{}
)

// Name alias for use in template
func (ml JsMath) Name() string {
	return "Math"
}

// Func as implementation of debug method
func (ml JsMath) Func() interface{} {
	return func() Math {
		return Math{}
	}
}

// Ceil rounds a value up to the next biggest integer
func (m Math) Ceil(x interface{}) int {
	if reflect.TypeOf(x).Kind() == reflect.Int {
		x = float64(reflect.ValueOf(x).Int())
	} else if reflect.TypeOf(x).Kind() == reflect.Int64 {
		x = float64(reflect.ValueOf(x).Int())
	}
	return int(math.Ceil(x.(float64)))
}

// Min gets the minimum value
func (m Math) Min(x ...interface{}) (res float64) {
	res = float64(math.MaxFloat64)
	for _, v := range x {
		if reflect.TypeOf(v).Kind() == reflect.Int {
			v = float64(reflect.ValueOf(v).Int())
		} else if reflect.TypeOf(v).Kind() == reflect.Int64 {
			v = float64(reflect.ValueOf(v).Int())
		}
		if v.(float64) < res {
			res = v.(float64)
		}
	}
	return
}

// Max gets the maximum value
func (m Math) Max(x ...interface{}) (res float64) {
	res = float64(math.SmallestNonzeroFloat64)
	for _, v := range x {
		if reflect.TypeOf(v).Kind() == reflect.Int {
			v = float64(reflect.ValueOf(v).Int())
		} else if reflect.TypeOf(v).Kind() == reflect.Int64 {
			v = float64(reflect.ValueOf(v).Int())
		}
		if v.(float64) > res {
			res = v.(float64)
		}
	}
	return
}