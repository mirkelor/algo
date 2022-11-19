package common

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func FindIndex[T Number](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
