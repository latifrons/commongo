package math

func SmallerInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func BiggerInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func BiggerUint64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func BiggerInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func SmallerUint64(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func SmallerInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
