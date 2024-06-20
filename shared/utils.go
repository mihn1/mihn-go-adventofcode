package shared

func Contains[S ~[]E, E comparable](source S, target E) bool {
	for _, v := range source {
		if v == target {
			return true
		}
	}
	return false
}

func IntPow(base, exp int) int {
	result := 1
	for exp != 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
