package Slice

import (
	"reflect"
)

//通配
func InArr(slice interface{}, v interface{}) int {
	switch reflect.ValueOf(v).Kind() {
	case reflect.String:
		return containsString(slice.([]string), v.(string))
	case reflect.Int:
		return containsInt(slice.([]int), v.(int))
	case reflect.Int64:
		return containsInt64(slice.([]int64), v.(int64))
	case reflect.Bool:
		return containsIntBool(slice.([]bool), v.(bool))
	case reflect.Uint64:
		return containsInUint64(slice.([]uint64), v.(uint64))
	case reflect.Float64:
		return containsInFloat(slice.([]float64), v.(float64))
	case reflect.Complex128:
		return containsInComplex(slice.([]complex128), v.(complex128))
	default:
		return contains(slice, v)
	}
}

//万能方法，牺牲性能
func contains(slice interface{}, v interface{}) (index int) {
	index = -1
	if slice == nil {
		return
	}
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}

//强类型，保证性能
func containsString(slice []string, v string) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInt(slice []int, v int) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInt64(slice []int64, v int64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsIntBool(slice []bool, v bool) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInUint64(slice []uint64, v uint64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInFloat(slice []float64, v float64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInComplex(slice []complex128, v complex128) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}
