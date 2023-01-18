package utils

func IsEqual(a interface{}, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

type Comparator func(value any) bool
