package hw_1

// order is important
func equal1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, el := range a {
		if el != b[idx] {
			return false
		}
	}
	return true
}

// order is not important
func equal2(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	aMap := make(map[int]int, len(a))
	bMap := make(map[int]int, len(b))

	for _, val := range a {
		aMap[val]++
	}

	for _, val := range b {
		bMap[val]++
	}

	for key, val := range aMap {
		if bMap[key] != val {
			return false
		}
	}

	return true
}
