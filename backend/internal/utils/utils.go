package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func GetFloat(unk interface{}) (float64, error) {
	if unk == nil {
		return math.NaN(), fmt.Errorf("cannot call GetFloat on a nil value")
	}

	var floatType = reflect.TypeOf(float64(0))
	var stringType = reflect.TypeOf("")

	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("can't convert %v to float64", v.Type())
		}
	}
}
