package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func runGreedy() {
	fmt.Printf("\nGreedy\n------\n")

	fmt.Println("Wildcard Matching")
	fmt.Println(isMatch("aa", "a"), isMatchGreedy("aa", "a"))
	fmt.Println(isMatch("", "?"), isMatchGreedy("", "?"))
	fmt.Println(isMatch("cb", "?a"), isMatchGreedy("cb", "?a"))
	fmt.Println(isMatch("aa", "*"), isMatchGreedy("aa", "*"))
	fmt.Println(isMatch("acdcb", "a*c?b"), isMatchGreedy("acdcb", "a*c?b"))
	fmt.Println(isMatch("abcabczzzde", "*abc???de*"), isMatchGreedy("abcabczzzde", "*abc???de*"))
	fmt.Println(isMatch("adceb", "*a*b"), isMatchGreedy("adceb", "*a*b"))

	fmt.Println("Largest Number")
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{10, 2}))
}

/*
179. Largest Number
100% time, 11% memory

Had to look at the solution ngl
*/
func largestNumber(nums []int) string {
	s_nums := make([]string, len(nums))
	countZeros := 0
	for i, n := range nums {
		s_nums[i] = strconv.Itoa(n)
		if n == 0 {
			countZeros++
		}
	}
	if countZeros == len(nums) {
		return "0"
	}
	slices.SortFunc(s_nums, func(i, j string) int {
		// string concatenation then check which is larger
		// you want to reverse sort, because you want the largest number concatenations to come first
		// so I flip the i and j in the Compare call
		// (you can also just iterate in reverse at the end instead of flipping here)
		return strings.Compare(j+i, i+j)
	})
	result := ""
	for i := range len(s_nums) {
		result += s_nums[i]
	}
	return result
}

/*
44. Wildcard Matching
51% time, 35% memory (2D DP non-greedy solution)

LEARNING: bool is way cheaper than int (obviously 1 bit versus 32/64)

# Very similar to regex pattern matching

'?' and '*'

Parameters
  - i index of s
  - j index of p

Bounds
  - i [0, len(s))
  - j [0, len(p))

Base Case (order of priority)
  - if i == n you succeed in the path
  - if j == m you fail the path
  - Note that the case (n,m) is a success because of priority

Recurrence Relations

	if p[j] == '?' {dp[i][j] = dp[i+1][j+1]}

If p[j] == '*' what do you do? Since it can build on itself and any existing path, it should look at (i+1, j), (i+1, j+1) and (i, j+1)

	dp[i][j] = dp[i+1][j] | dp[i+1][j+1] | dp[i][j+1]
	current_position = using_itself | continuing_a_path_by_adding_1_character | continuing_a_path_by_adding_0_characters

	  b a *
	a[0 1 1 0]
	a[0 1 1 0]
	a[0 0 1 0]
	 [0 0 0 1] -> output is dp[0][0] (False)
*/
func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	// explicitly checking some missing inputs
	if n == 0 {
		for _, c := range p {
			if c != '*' {
				return false
			}
		}
		return true
	} else if m == 0 {
		return false
	}

	dp := make([][]bool, n+1)
	for i := range n + 1 {
		dp[i] = make([]bool, m+1)
	}

	dp[n][m] = true
	for j := m - 1; j >= 0; j-- {
		// this loop is needed because I noticed that there's no way for the dp loop to test skipping suffix '*' wildcards
		if p[j] == '*' {
			dp[n][j] = true
		} else {
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if p[j] == '?' || p[j] == s[i] {
				dp[i][j] = dp[i+1][j+1]
			} else if p[j] == '*' {
				dp[i][j] = dp[i+1][j] || dp[i+1][j+1] || dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}

/*
Same question as above but this is the proper greedy solution taken from LeetCode Solutions
100% time, 82% memory

Use two indices sIdx and pIdx to traverse the string s and pattern p.
Keep track of the last seen '*' in the pattern with starIdx.
Use sTmpIdx to remember the position in s corresponding to the last '*' in p.
*/
func isMatchGreedy(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	sIdx := 0
	pIdx := 0
	starIdx := -1
	sTmpIdx := -1

	for sIdx < sLen {
		if pIdx < pLen && (s[sIdx] == p[pIdx] || p[pIdx] == '?') {
			// Case 1: Characters match or '?' matches any single character
			sIdx++
			pIdx++
		} else if pIdx < pLen && p[pIdx] == '*' {
			// Case 2: '*' matches zero or more characters
			starIdx = pIdx
			sTmpIdx = sIdx
			pIdx++
		} else if starIdx != -1 {
			// Case 3: Mismatch and backtrack to last '*'
			pIdx = starIdx + 1
			sTmpIdx++
			sIdx = sTmpIdx
		} else {
			// Case 4: No match possible
			return false
		}
	}

	// Check remaining characters in pattern are '*'
	for pIdx < pLen && p[pIdx] == '*' {
		pIdx++
	}

	return pIdx == pLen
}
