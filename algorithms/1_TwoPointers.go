package algorithms

import "sort"

type Pair struct {
	Value int
	Index int
}

func twoSum(nums []int, target int) []int {
	arr := make([]Pair, len(nums))

	for i, v := range nums {
		arr[i] = Pair{v, i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value < arr[j].Value
	})

	i, j := 0, len(arr)-1

	for i < j {
		sum := arr[i].Value + arr[j].Value

		if sum == target {
			return []int{arr[i].Index, arr[j].Index}
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}
