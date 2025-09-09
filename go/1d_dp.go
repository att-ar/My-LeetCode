package main

import (
	"fmt"
	"math"
	"slices"
)

func run1dDp() {
	fmt.Printf("\n1D-DP\n-----\n")
	fmt.Println("Min Cost Climbing Stairs")
	slice := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(slice))
	slice = []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(slice))
	fmt.Println(minCostClimbingStairsConstantSpace(slice))

	fmt.Println("Best Sightseeing Pair")
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))

	fmt.Println("Count Ways to Build Good Strings")
	fmt.Println(countGoodStrings(3, 3, 1, 1))
	fmt.Println(countGoodStrings(1, 100000, 1, 1))
	fmt.Println(countGoodStrings(2, 3, 1, 2))

	fmt.Println("Minimum Cost For Tickets")
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))
	fmt.Println(mincostTicketsBottomUp(days, costs))
}

/*
983. Minimum Cost For Tickets
100% time, 60% memory

I will do it recursively top down first
*/
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	memo := make([]int, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		var j, res int

		// base case
		if i >= n {
			return 0
		} else if memo[i] > 0 {
			return memo[i]
		}

		j = i + 1
		// go to next index in `days` since the days in between you can skip
		res = dfs(j) + costs[0]

		// need to find the next index that is not covered by buying a 7 day pass
		for j < n && days[j] < days[i]+7 {
			j++
		}
		if j > i {
			res = min(res, dfs(j)+costs[1])
		}

		// need to find the next index that is not covered by buying a 30 day pass
		for j < n && days[j] < days[i]+30 {
			j++
		}
		if j > i {
			res = min(res, dfs(j)+costs[2])
		}

		memo[i] = res
		return res
	}

	return dfs(0)
}

/*
983. Minimum Cost For Tickets
100% time, 75% memory

Same question as right above but doing it Bottom Up.

There is a way to avoid computing `j` in every single loop. You would have to swap it to start from n outside the loop.
Then flip the comparison to be

	days[j-1] >= days[i] + (7 or 30) {
	    j--
	}

I tried it but it didn't work on the first try, so no need to keep trying
*/
func mincostTicketsBottomUp(days []int, costs []int) int {
	n := len(days)

	// dp array
	dp := make([]int, n+1)
	// implement base case (past the max day, zero cost)
	dp[n] = 0

	var j, res int
	for i := n - 1; i >= 0; i-- {
		j = i + 1

		// go to next index in `days` since the days in between you can skip
		res = dp[j] + costs[0]

		// need to find the next index that is not covered by buying a 7 day pass
		for j < n && days[j] < days[i]+7 {
			j++
		}
		if j > i {
			res = min(res, dp[j]+costs[1])
		}

		// need to find the next index that is not covered by buying a 30 day pass
		for j < n && days[j] < days[i]+30 {
			j++
		}
		if j > i {
			res = min(res, dp[j]+costs[2])
		}

		dp[i] = res
	}

	return dp[0]
}

/*
2466. Count Ways To Build Good Strings
94% time, 67% memory

The setup is all hand waviness.
  - You are given two numbers `zero` and `one`, those are the stepsizes you are allowed.
  - How many linear combinations of those two integers result in a sum between `low` and `high`

# Thinking there is one parameter `i` which represents the current sum

Parameters
  - i : the current sum

Bounds
  - i : [0, high]

The number of ways to reach a sum between low and high at position `i` is the sum of the ways to get to a target sum by either
  - adding `zero` to `i`
  - adding `one` to `i`

Recurrence Relation in a forward pass (just reverse the logic I describe above)

	dp[i] = dp[i-zero] + dp[i-one]
*/

func countGoodStrings(low int, high int, zero int, one int) int {
	mod := int(math.Pow(10, 9)) + 7
	// max bound that i can reach is high (inclusive)
	dp := make([]int, high+1)
	// init base case
	dp[0] = 1
	var useZero, useOne int
	for i := min(one, zero); i <= high; i++ {
		if i >= zero {
			useZero = dp[i-zero]
		}
		if i >= one {
			useOne = dp[i-one]
		}
		dp[i] = (useZero + useOne) % mod
	}
	result := 0
	for i := low; i <= high; i++ {
		result += dp[i]
		result %= mod
	}
	return result
}

/*
1014. Best Sightseeing Pair

100% time, 31% memory

Parameters
  - i index of values currently being processed

Bounds
  - i [0, len(values)) in reverse direction

What is actually happening though?
  - Store the max effective value you can use at position i
  - This needs to take into account (value and the distance of the value from i)

Example

	[8,1,5,2,6] + [0] (ZERO PADDED)
	@n-1: dp[i] = max(nums[i], dp[i+1] - 1) = max(6, 0-1) = dp[i] = 6
	@n-2: dp[i] = max(nums[i], dp[i+1] - 1) = max(2, 6-1) = dp[i] = 5

Notice that dp[i] is only reliant on dp[i+1] (you only care about the current value and the stored max for comparison).
So I don't need an array, just one variable.

Update the best pair as you iterate.
*/
func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	dp := 0
	result := 0
	for i := n - 1; i >= 0; i-- {
		dp--
		// in Go, the actual verbose if-statement is WAY WAY faster than just doing variable = max(variable, other)
		if values[i]+dp > result {
			result = values[i] + dp
		}
		if values[i] > dp {
			dp = values[i]
		}
	}
	return result
}

// minCost using a dp array
// 100% time, 32% memory
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 3 {
		return slices.Min[[]int](cost)
	}
	dp := make([]int, len(cost))
	// dp[i] = min(dp[i], dp[i-1]) + cost[i]
	// cost[i] is irrelevant to choosing the optimal path up to i though.
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	// can go all the way up the staircase from either of the last two steps
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

// minCost using just two variables to hold i-1 and i-2
// 56% time, 99.8% memory
func minCostClimbingStairsConstantSpace(cost []int) int {
	if len(cost) < 3 {
		return slices.Min[[]int](cost)
	}
	// same logic as dp array, but you don't need entire history
	i_2, i_1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		i_2, i_1 = i_1, min(i_2, i_1)+cost[i]
	}
	// can go all the way up the staircase from either of the last two steps
	return min(i_2, i_1)
}
