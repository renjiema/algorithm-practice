package modify_array_in_place

func twoSum(nums []int, target int) []int {
	//两个map,一个记录num值，一个记录target-num值
	subValues, n := make(map[int]int), len(nums)
	for i := 0; i < n; i++ {
		if j, ok := subValues[i]; ok {
			return []int{i, j}
		}
		sub := target - nums[i]
		subValues[sub] = i
	}
	return []int{-1, -1}
}
