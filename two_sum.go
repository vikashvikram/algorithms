package algorithms

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if i, ok := m[target-val]; ok {
			return []int{i, index}
		} else {
			m[val] = index
		}
	}
	return []int{-1, -1}
}
