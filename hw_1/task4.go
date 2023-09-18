package hw_1

type MySlice struct {
	arr []int
}

func (ms *MySlice) sort() []int {
	return mergeSort(ms.arr, 0, len(ms.arr)-1)
}

func mergeSort(a []int, left, right int) []int {
	if left == right {
		result := []int{a[left]}
		return result
	}
	middle := (left + right) / 2
	leftPart := mergeSort(a, left, middle)
	rightPart := mergeSort(a, middle+1, right)
	return merge(leftPart, rightPart)
}

func merge(a, b []int) []int {
	res := []int{}
	left, right := 0, 0
	for left < len(a) && right < len(b) {
		if a[left] <= b[right] {
			res = append(res, a[left])
			left++
		} else {
			res = append(res, b[right])
			right++
		}
	}
	for left < len(a) {
		res = append(res, a[left])
		left++
	}
	for right < len(b) {
		res = append(res, b[right])
		right++
	}
	return res
}
