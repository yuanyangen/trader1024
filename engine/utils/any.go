package utils

func AnyToBool(in any) bool {
	if in == nil {
		return false
	}
	r, _ := in.(bool)
	return r
}
