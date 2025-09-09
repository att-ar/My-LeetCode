package main

import "fmt"

func runSlidingWindow() {
	kWindows := kWindowsNListsDistanceD([][]int{
		{4, 1, 5, 6, 1, 3},
		{1, 9, 5, 2, 1, 3},
		{5, 4, 5, 8, 1, 3},
		{9, 9, 5, 11, 1, 2},
		{9, 1, 5, 11, 12, 122},
	}, 6, 3)
	fmt.Println(kWindows)
}

/*
Google OA from internet question that I did in Python first:

You need to:

- Compute the **k-window** for each list: a set of distinct elements from the first `k` elements.

- Ensure that the `k-window` of list `i` shares **no common elements** with the `k-windows` of the previous `d` lists (`i-d, ..., i-1`).

- Return the resulting `k-windows` for all `n` lists.
*/
func kWindowsNListsDistanceD(
	lists [][]int,
	k int,
	d int,
) [][]int {
	// init new map {int:int} with size n*k
	// (worst case space complexity since I am not deleting entries when they are rolled out)
	n := len(lists)
	rollingSeen := make(map[int]int, n*k)
	kWindows := make([][]int, n)
	// go through every list one time
	for i := 0; i < n; i++ {
		// roll out the out of bounds list from the window
		if i-d-1 > 0 {
			for _, num := range lists[i-d-1] {
				rollingSeen[num] -= 1
			}
		}
		// go through the first k numbers one time
		for _, num := range lists[i][:k] {
			// if the value has not been seen or it has not been rolled out completely, we add it to our result
			if val, exists := rollingSeen[num]; !(exists && (val > 0)) {
				kWindows[i] = append(kWindows[i], num)
			}

			rollingSeen[num] += 1
		}

	}
	return kWindows
}
