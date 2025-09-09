package main

import (
	"fmt"
	"strconv"
)

func runArrays() {
	fmt.Println("topKFrequent")
	nums := []int{1, 6, 4, 4, 3, 3}
	fmt.Println(nums)
	fmt.Println(topKFrequent(nums, 2))

	fmt.Println("productExceptSelf")
	nums = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))

	fmt.Println("Maximum Score After Splitting a String")
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00"))
	fmt.Println(maxScoreSpaceOptimized("011101"))

	fmt.Println("Count Vowel Strings in Ranges")
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	fmt.Println(vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))

	fmt.Println("Number Ways to Split Array")
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))

	fmt.Println("Minimum Number of Operations to Move All Balls to Each Box")
	fmt.Println(minOperations("001011"))
	fmt.Println(minOperations("110"))

	fmt.Println("Max Sum of a Pair With Equal Sum of Digits")
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
	fmt.Println(maximumSum([]int{229, 398, 269, 317, 420, 464, 491, 218, 439, 153, 482, 169, 411, 93, 147, 50, 347, 210, 251, 366, 401}))
}

/*
2342. Max Sum of a Pair With Equal Sum of Digits
74% time, 5% memory (using hashmap)
98% time, 89% memory (using array)

The largest digitSum you can have is 81 because the largest integer available is 10^9 which means nine 9s is the max sum
*/
func maximumSum(nums []int) int {
	// will store the current largest integer with that digit sum
	digitSumMap := [82]int{} // see docstring
	result := -1
	/* Will go through each number and upate result everytime a larger pair is found for a given digitSum */
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var digitSum, tmp int
	for _, num := range nums {
		tmp = num
		for tmp > 0 {
			digitSum += tmp % 10
			tmp /= 10
		}
		if digitSumMap[digitSum] > 0 {
			result = max(result, digitSumMap[digitSum]+num)
		}
		digitSumMap[digitSum] = max(digitSumMap[digitSum], num)
		// clear the hash
		digitSum = 0
	}
	return result
}

/* 18% time, 45% memory */
func maximumSumSlow(nums []int) int {
	// will store the 2 largest numbers per digitSum key
	digitSumMap := make(map[int][2]int)

	getDigitSum := func(num int) int {
		sNum := strconv.Itoa(num)
		res := 0
		var tmp int
		for i := range sNum {
			tmp, _ = strconv.Atoi(string(sNum[i]))
			res += tmp
		}
		return res
	}

	var digitSum int
	var exists bool
	pairExists := false
	for i := range nums {
		digitSum = getDigitSum(nums[i])
		if _, exists = digitSumMap[digitSum]; exists {
			pairExists = true
			// need to keep the two largest entries
			// the largest number will be in index 1, the second lagest in index 0
			if nums[i] >= digitSumMap[digitSum][1] {
				digitSumMap[digitSum] = [2]int{digitSumMap[digitSum][1], nums[i]}
			} else if nums[i] > digitSumMap[digitSum][0] {
				digitSumMap[digitSum] = [2]int{nums[i], digitSumMap[digitSum][1]}
			}
		} else {
			digitSumMap[digitSum] = [2]int{0, nums[i]}
		}
	}
	if !pairExists {
		// edge case, there is a case where there are no pairs at all
		return -1
	}
	result := 0
	var tempSum int
	for _, pair := range digitSumMap {
		if pair[0] == 0 {
			// edge case, if there is only one occurrence of the digitSum, it isn't a pair so skip
			continue
		}
		tempSum = pair[0] + pair[1]
		if tempSum > result {
			result = tempSum
		}
	}
	// fmt.Printf("%+v\n", digitSumMap)
	return result
}

/*
1769. Minimum Number of Operations to Move All Balls to Each Box
100% time, 95% memory

I'm thinking you do one pass to get the total number of '1' and the total distance for all of them from the zeroth index

[1,0,0,1,0,1] -> 3 '1's with 0 + 3 + 5 = 8 total distance from index 0

	rightCount = 3
	leftCount = 0
	totalDist = 8

then start iter:

i = 0

	res.append(totalDist)
	{
		rightCount = 3 - 1
		leftCount = 0 + 1
	}
	{
		totalDist -= rightCount - leftCount
	}
*/
func minOperations(boxes string) []int {
	leftcount, rightCount := 0, 0
	totalDist := 0
	result := make([]int, len(boxes))
	for i := range boxes {
		if boxes[i] == '1' {
			rightCount++
			totalDist += i
		}
	}
	for i := range boxes {
		result[i] = totalDist

		if boxes[i] == '1' {
			leftcount++
			rightCount--
		}

		totalDist -= (rightCount - leftcount)
	}

	return result
}

/*
2270. Number of Ways to Split Array
SuffixSum problem (Range Query problem essentially)

35% time, 60% memory (2 pass solution with array)
100% time, 75% memory (2 pass solution with single variable)
*/
func waysToSplitArray(nums []int) int {
	n := len(nums)
	// suffixSum := make([]int, n)
	// suffixSum[n-1] = 0
	// for i := n - 2; i >= 0; i-- {
	// 	suffixSum[i] = suffixSum[i+1] + nums[i+1]
	// }
	suffixSum := 0
	for i := range n {
		suffixSum += nums[i]
	}

	result := 0
	curSum := 0
	for i := range n - 1 {
		curSum += nums[i]
		// if curSum >= suffixSum[i] {
		// 	result++
		// }
		suffixSum -= nums[i]
		if curSum >= suffixSum {
			result++
		}
	}
	return result
}

/*
2559. Count Vowel Strings in Ranges
100% time, 62% memory

Topics say Prefix Sum, I was thinking of using a heap similar to the other Alice and Bob query question lol.

But prefix sum makes sense, for every index `i` you compute how many words start and end with a vowel up to and including `i`
  - Should also store a variable for whether or not index `i` itself contributes
*/
func vowelStrings(words []string, queries [][]int) []int {
	checkVowel := func(r byte) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	n := len(words)
	prefixSum := make([][2]int, n)
	if checkVowel(words[0][0]) && checkVowel(words[0][len(words[0])-1]) {
		prefixSum[0] = [2]int{1, 1}
	}
	for i := 1; i < n; i++ {
		if checkVowel(words[i][0]) && checkVowel(words[i][len(words[i])-1]) {
			prefixSum[i] = [2]int{prefixSum[i-1][0] + 1, 1}
		} else {
			prefixSum[i] = [2]int{prefixSum[i-1][0], 0}
		}
	}

	result := make([]int, len(queries))
	for j, query := range queries {
		// adding the 1 or 0 back at the end dependant on whether or not the query[0] position is an appropriate vowel word
		// because the subtraction will remove 1 unnecessarily from the total if that left index word is valid
		result[j] = prefixSum[query[1]][0] - prefixSum[query[0]][0] + prefixSum[query[0]][1]
	}
	return result
}

/*
1422. Maximum Score After Splitting a String
100% time, 24% memory

NOTE: non-empty substring
*/
func maxScore(s string) int {
	suffixSumOnes := make([]int, len(s)+1)
	suffixSumOnes[len(s)] = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			suffixSumOnes[i] = suffixSumOnes[i+1] + 1
		} else {
			suffixSumOnes[i] = suffixSumOnes[i+1]
		}
	}
	resValue := 0
	countZeros := 0
	// the len(s) - 1 is because we require a non-empty substring
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			countZeros++
		}
		if suffixSumOnes[i+1]+countZeros > resValue {
			resValue = suffixSumOnes[i+1] + countZeros
		}
	}

	return resValue
}

/*
Clever way to update the result based on the overall number of ones.
*/
func maxScoreSpaceOptimized(s string) int {
	oneCount := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			oneCount++
		}
	}
	cur := 1
	if s[0] == '1' {
		cur = 0
	}
	maxS := cur + oneCount
	m := maxS
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			m++
		} else {
			m--
		}
		if m > maxS {
			maxS = m
		}
	}
	return maxS
}

func topKFrequent(nums []int, k int) []int {
	hold := make([][]int, len(nums)+1)
	dct := make(map[int]int, len(nums))
	res := []int{}

	for _, n := range nums {
		dct[n]++
	}

	for key, value := range dct {
		hold[value] = append(hold[value], key)
	}
	for i := len(nums); i >= 0; i-- {
		res = append(res, hold[i]...)
		if len(res) == k {
			break
		}
	}
	return res
}

// 238. Product of Array Except Self
// 64% time, 25% memory
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i > -1; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
