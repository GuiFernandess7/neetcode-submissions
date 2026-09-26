func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)

	for pos, n := range nums {
		t := target - n
		if _, ok := hash[n]; ok {
			return []int{hash[n], pos}
		}
		hash[t] = pos
	}
	return []int{}
}
