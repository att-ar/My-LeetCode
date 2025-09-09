package main

import "fmt"

func runBacktrack() {
	fmt.Println("Subsets")
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2}
	fmt.Println(subsets(nums1))
	fmt.Println(subsets(nums2))
}

// 75% time, 96% memory
func subsets(nums []int) [][]int {
	res := &[][]int{} // pointer to slice
	hold := &[]int{}  // pointer to subset being built
	var backtrack func(i int, res *[][]int, hold *[]int)
	backtrack = func(i int, res *[][]int, hold *[]int) {
		if i == len(nums) {
			*res = append(*res, append([]int{}, *hold...))
			return
		}
		*hold = append(*hold, nums[i])
		backtrack(i+1, res, hold)
		*hold = (*hold)[:len(*hold)-1]
		backtrack(i+1, res, hold)
	}

	backtrack(0, res, hold)
	return *res
}
