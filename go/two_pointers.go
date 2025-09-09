package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
)

func runTwoPointers() {
	fmt.Println("Valid Palindrome")
	// s := [3]string{
	// 	"A man, a plan, a canal: Panama",
	// 	"race a car",
	// 	" ",
	// }
	// for _, sub := range s {
	// 	fmt.Println(isPalindrome((sub)))
	// 	fmt.Println(isPalindromeFast(sub))
	// }

	fmt.Println("\n2 Sum Sorted")
	// numbers := []int{2, 3, 4, 6, 11, 15}
	// target := 9
	// fmt.Println(twoSum2(numbers, target))

	fmt.Println("\n3 Sum")

	start := time.Now()
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
	// fmt.Println(threeSum([]int{-13, 11, 11, 0, -5, -14, 12, -11, -11, -14, -3, 0, -3, 12, -1, -9, -5, -13, 9, -7, -2, 9, -1, 4, -6}))
	// fmt.Println(threeSum([]int{0, 0, 0, 0}))
	// fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
	fmt.Println("Duration Raw:", time.Since(start))

	start = time.Now()
	// fmt.Println(threeSumGoroutines([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSumGoroutines([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
	// fmt.Println(threeSumGoroutines([]int{-13, 11, 11, 0, -5, -14, 12, -11, -11, -14, -3, 0, -3, 12, -1, -9, -5, -13, 9, -7, -2, 9, -1, 4, -6}))
	// fmt.Println(threeSumGoroutines([]int{0, 0, 0, 0}))
	// fmt.Println(threeSumGoroutines([]int{-2, 0, 1, 1, 2}))
	// fmt.Println(threeSumGoroutines([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
	fmt.Println("Duration Chan:", time.Since(start))

	fmt.Println("Container with most water")
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println("Trapping rain water")
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

// trapping rain water
// 57% time, 83% memory
func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	for left < right {
		// given that each position is limited by the shorter height of the tallest heights on its left or right
		// even if there is a height of 10 on the right, a height of 1 on the left of the position would mean 1 block of water
		// treating every x-tick one at a time so area added is just the shortest max height available.

		// check which side is limiting:
		if maxLeft <= maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			res += maxLeft - height[left] // the subtraction is needed since it is possible that the position that left got moved to is not of height 0
		} else {
			right--
			maxRight = max(maxRight, height[right])
			res += maxRight - height[right]
		}
	}
	return res
}

// Container with most water
// 42% time, 48% memory
func maxArea(height []int) int {
	// start pointers from the outside since the largest area is determined by 2 variables:
	// distance between lines (x) and minimum of the two lines' heights (y)
	// maxing the x-axis is guaranteed, maxing the y axis isnt
	res := 0
	for left, right := 0, len(height)-1; left < right; {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 3Sum using goroutines: waitgroup and mutex
func runIth2SumSorted(wg *sync.WaitGroup, ch chan<- []int, nums []int, i int) {
	defer wg.Done()
	// initialize ith two sum sorted
	target := nums[i]
	left, right := i+1, len(nums)-1
	// ith two sum sorted
	for left < right {
		if target+nums[left]+nums[right] == 0 {
			// send a triplet to the write-only channel
			ch <- []int{target, nums[left], nums[right]}
			left++
			for (left < right) && (nums[left-1] == nums[left]) {
				left++
			}
		} else if target+nums[left]+nums[right] > 0 {
			right--
			for (right > left) && (nums[right+1] == nums[right]) {
				right--
			}
		} else {
			left++
			for (left < right) && (nums[left-1] == nums[left]) {
				left++
			}
		}
		// stop dupes for the same target (move right and left to new numbers)
	}
}

func threeSumGoroutines(nums []int) [][]int {
	// best way to think of it is breaking it down into N 2sumSorted problems
	sort.Ints(nums)
	res := make([][]int, 0)
	var wg sync.WaitGroup
	ch := make(chan []int)

	// start the n goroutines to solve the 2sumSorted problems
	for i := range len(nums) - 2 {
		if (1 <= i) && (nums[i-1] == nums[i]) {
			continue
		}
		wg.Add(1)
		go runIth2SumSorted(&wg, ch, nums, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for triplet := range ch {
		res = append(res, triplet)
	}

	return res
}

// 37% time , 37% memory
func threeSum(nums []int) [][]int {
	// best way to think of it is breaking it down into N 2sumSorted problems
	sort.Ints(nums)
	res := make([][]int, 0)
	i := 0
	for i < len(nums)-2 {
		// initialize ith two sum sorted
		// skip
		for 1 <= i && i < len(nums)-2 && nums[i-1] == nums[i] {
			i++
		}
		target := nums[i]
		left, right := i+1, len(nums)-1

		// ith two sum sorted
		for left < right {
			if target+nums[left]+nums[right] == 0 {
				res = append(res, []int{target, nums[left], nums[right]})
				left++
				for (left < right) && (nums[left-1] == nums[left]) {
					left++
				}
			} else if target+nums[left]+nums[right] > 0 {
				right--
				for (right > left) && (nums[right+1] == nums[right]) {
					right--
				}
			} else {
				left++
				for (left < right) && (nums[left-1] == nums[left]) {
					left++
				}
			}
		}
		i++
	}
	return res
}

// 167 Two Sum 2. Input Array is Sorted
// 66% time, 40% memory
func twoSum2(numbers []int, target int) []int {
	// array is sorted so you can just have two pointers start from the min and max
	// then they move inwards until the sum is found
	// O(1) memory
	// O(n) time
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{left, right}
}

// 125. Valid Palindrome
// 43# time, 50% memory
func isPalindrome(s string) bool {
	str := new(strings.Builder)
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			str.WriteRune(char)
		}
	}
	s2 := str.String()
	left, right := 0, 0
	for right < len(s2)-1 {
		left++
		right++
		right++
	}
	if right == len(s2) {
		right--
	}
	start := 0
	for start < left {
		if s2[start] != s2[right] {
			return false
		}
		start++
		right--
	}
	return true
}

// 100% time, 69.9% memory
func isPalindromeFast(s string) bool {
	s = strings.ToLower(s) // lowercase
	l, r := 0, len(s)-1
	for l < r {
		// skipping non alphanumerics:
		for l < r && !(unicode.IsDigit(rune(s[l])) || unicode.IsLetter(rune(s[l]))) {
			l++
		}
		for r > l && !(unicode.IsDigit(rune(s[r])) || unicode.IsLetter(rune(s[r]))) {
			r--
		}
		// check palindrome
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
