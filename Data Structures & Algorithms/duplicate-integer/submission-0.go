func hasDuplicate(nums []int) bool {
    hash := make(map[int]int)
    
    for key, val := range nums {
        if _, ok := hash[val]; ok {
            return true
        }
        hash[val] = key
    }
    return false
}
