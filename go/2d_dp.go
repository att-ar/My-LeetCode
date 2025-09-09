package main

import (
	"fmt"
	"math"
)

func run2dDp() {
	fmt.Printf("\n2D-DP\n-----\n")
	fmt.Println("Max of 3 subarray sums")
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
	fmt.Println(maxSumOfThreeSubarrays([]int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}, 3))

	fmt.Println("Number of Ways to Form a Target String Given a Dictionary")
	fmt.Println(numWays([]string{"acca", "bbbb", "caca"}, "aba"))
	fmt.Println(numWays([]string{"abba", "baab"}, "bab"))

	fmt.Println("Unique Length-3 Palindromic Subsequences")
	fmt.Println(countPalindromicSubsequence("aabca"))
	fmt.Println(countPalindromicSubsequence("adc"))
	fmt.Println(countPalindromicSubsequence("bbcbaba"))
}

/*
1930. Unique Length-3 Palindromic Subsequences
38% time, 62% memory
Length-3 palindrome just means the first and last letters need to match
Build a Counter from s[1:]
then starting from index 0, go through s, while updating the counter by decrementing the value you are processing
*/
func countPalindromicSubsequence(s string) int {
	// counter for characters to the right of the index we are processing
	n := len(s)
	counter := [26]int{}
	seenPrefix := map[[2]byte]bool{}
	leftLetters := make(map[byte]struct{}, 26)
	for i := 1; i < n; i++ {
		// note that runes are int32 which work as indices
		counter[s[i]-'a']++
	}

	leftLetters[s[0]] = struct{}{}
	result := 0

	for i := 1; i < n; i++ {
		// decrement counter for current letter
		counter[s[i]-'a']--

		for leftLetter := range leftLetters {
			// if this prefix hasn't been seen and there is a valid palindrome
			if !seenPrefix[[2]byte{leftLetter, s[i]}] && counter[leftLetter-'a'] > 0 {
				seenPrefix[[2]byte{leftLetter, s[i]}] = true
				result++
			}
		}

		leftLetters[s[i]] = struct{}{}
	}
	return result

}

/*
Much more interesting approach using knowledge of the indices:


func countPalindromicSubsequence(s string) int {
    indexes := make(map[rune][]int)
    for k, v := range s {
        if indexes[v] == nil {
            indexes[v] = make([]int, 0)
        }
        indexes[v] = append(indexes[v], k)
    }

    res := 0
    for _, idxs := range indexes {
        if len(idxs) >= 2 {
            l := idxs[0]
            r := idxs[len(idxs)-1]

            for _, idxs2 := range indexes {
                for _, i := range idxs2 {
                    if i > l && i < r  {
                        res++
                        break
                    }
                }
            }
        }
    }
    return res
} */

/*
1639. Number of Ways to Form a Target String Given a Dictionary
14% time, 42% memory

Reimplementing my python solution (minus the use of `ord` and array indexing instead of a hashmap)

Parameters
  - i index of target
  - k index of letters (inclusive, so we are allowed to use index k at position (i,k))

Given a position (i, k), the number of ways you can get to the end of target is the number of ways that taking the letter or skipping the letter can get you to the end of target

	dp[i][k] = dp[i+1][k+1] * int(letters[k][target[i]]) + dp[i][k+1]

The multiplication is to scale the number of ways taking succeeds by the number of occurrences the current position's letters match
  - Can be any integer >= 0, 0 would mean that taking doesn't actually work because the current position's letters don't match
*/
func numWays(words []string, target string) int {
	K := len(words[0])
	n := len(target)
	MOD := int(math.Pow(10.0, 9.0)) + 7
	// turn words into a frequency dictionary `letters`
	letters := make([]map[rune]int, K)
	for _, word := range words {
		for k, letter := range word {
			// Lazily initialize the map if it's nil
			if letters[k] == nil {
				letters[k] = make(map[rune]int)
			}
			letters[k][letter]++
		}
	}

	// dp matrix
	dp := make([][]int, n+1)
	for i := range n + 1 {
		// padding the rightmost column with zeros because if you are out of bounds for words but not target the path has failed
		dp[i] = make([]int, K+1)
	}
	// padding the bottom row with 1s because if you reach out of bounds for `target` then the path has succeeded
	for k := range K + 1 {
		dp[n][k] = 1
	}

	// dp execution
	for i := n - 1; i >= 0; i-- {
		for k := K - 1; k >= 0; k-- {
			dp[i][k] = (dp[i+1][k+1]*int(letters[k][rune(target[i])]) + dp[i][k+1]) % MOD
		}
	}
	return dp[0][0]
}

/*
689. Maximum Sum of 3 Non-Overlapping Subarrays
63% time, 28% memory

# Already solved this in python but I want to try a different way closer to NeetCode's solution

# Recurrence Analysis

Parameters
- `i` index of k_sums (index of starting position in nums that forms the ith k sum)
- `count` number of subarrays used to make the current path's sum

Bounds
- i in [0, n - k + 1)
- count in [0, 3]

Direction
for Bottom Up `i` goes down as `count` goes up

for a position (i, count) the result is the max of taking position i or skipping position i
REMEMBER that we are going bottom up so earlier solutions are at `count-1`

	dp[i][count] = max(dp[i+k][count-1], dp[i+1][count])
*/
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	n_effective := n - k + 1 // the number of iterations we actually have to do in DP

	// preprocessing by calculating all the k_sums
	k_sums := make([]int, n_effective)
	for i := range k {
		k_sums[0] += nums[i]
	}
	for i := 1; i < n_effective; i++ {
		k_sums[i] = k_sums[i-1] - nums[i-1] + nums[i+k-1]
	}

	// DP part, (n_effective, 4) matrix
	// the leftmost column (count == 0) is just zero padding
	// the bottommost row is zero padding for when i goes out of bounds
	dp := make([][4]int, n_effective+1)
	var take, skip int
	for i := n_effective - 1; i >= 0; i-- {
		// calculate max sums of 1 and 2 non-overlapping subarrays
		for count := 1; count < 4; count++ {
			if i > n-k*count {
				// not enough number left so don't bother checking
				break
			}
			// I don't want to add k padding rows at the bottom so I only add one and then take the min(i+k, n_effective) to use the padding
			take = k_sums[i] + dp[min(i+k, n_effective)][count-1]
			skip = dp[i+1][count]
			dp[i][count] = max(take, skip)
		}
	}

	// now get the correct indices
	result := make([]int, 3)
	count := 0
	i := 0
	for count < 3 {
		// check for which indices lead to the earliest max
		take = k_sums[i] + dp[min(i+k, n_effective)][3-count-1]
		skip = dp[i+1][3-count]

		if take >= skip {
			result[count] = i
			count += 1
			i += k
		} else {
			i += 1
		}
	}
	return result
}
