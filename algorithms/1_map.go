package algorithms

func twoSum(nums []int, target int) []int {
	used := make(map[int]int)

	for i, v := range nums {
		need := target - v

		if j, ok := used[need]; ok {
			return []int{j, i}
		}

		used[v] = i
	}

	return nil
}
