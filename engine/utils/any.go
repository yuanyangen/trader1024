package utils

import "github.com/yuanyangen/trader1024/engine/model"

func AnyToBool(in any) bool {
	if in == nil {
		return false
	}
	r, _ := in.(bool)
	return r
}

func AnyToFloat(in any) (float64, bool) {
	if in == nil {
		return 0, false
	}
	r, ok := in.(float64)
	if !ok {
		panic("should not reach here")
	}
	return r, true
}

func AnySliceToFloat(in []any) []float64 {
	if in == nil {
		return nil
	}
	res := make([]float64, len(in))
	for i, v := range in {
		r, ok := v.(float64)
		if !ok {
			panic("should not reach here")
		}
		res[i] = r
	}

	return res
}

func DataNodeSliceToFloat(in []model.DataNode) []float64 {
	if in == nil {
		return nil
	}
	res := make([]float64, len(in))
	for i, v := range in {
		r := v.GetValue()
		res[i] = r
	}

	return res
}
