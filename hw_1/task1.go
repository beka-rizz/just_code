package hw_1

func TwoSum(nums []int, target int) []int {
	myMap := make(map[int]int)  

	for index, num := range nums {
		if _, isExist := myMap[target-num]; isExist {
			return []int{myMap[target-num], index}
		} else {
			myMap[num] = index
		}
	}
	return []int{}
}

