package main

import (
	"fmt"
	"time"
)

func runAdvancedGraphs() {
	fmt.Printf("\n\nAdvanced Graphs\n--------------\n")
	fmt.Println("Swim in Rising Water")

	start := time.Now()
	fmt.Println(swimInWater([][]int{{0, 2}, {1, 3}}))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(swimInWaterBinarySearch([][]int{{0, 2}, {1, 3}}))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(swimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(swimInWaterBinarySearch([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}))
	fmt.Println(time.Since(start))
}

/*
778. swim in rising water
78% time, 69% memory

Since you can move an infinite distance, the question is just asking to find the path with the minimum max height along a path.
The max height on a path is the amount of time it takes to travel it (rate limiting step)
The result path is the path with the smallest max height (shortest rate limiting step)

Djikstra's where the cost of the path is the maximum height currently seen. The first path that reaches the end is the solution
But how do you handle cycles?
- maintaining the positions in the currently visited path. The only way I can think of right now.

CORRECTION: you actually don't care about specific paths because the heap ensures that IF you get to a position (i,j) you NECESSARILY took the most efficient way there, so you never want to come back to it
instead of a path variable for each store added to the heap, just maintain a global `seen` hashset

NOTE: I declared the Store struct right beneath this function
*/
func swimInWater(grid [][]int) int {

	// nxn grid
	n := len(grid)
	minHeap := NewHeap(func(a, b Store) bool { return a.cost < b.cost })
	// start traversal from top left
	minHeap.HeapPush(Store{
		cost:     grid[0][0],
		position: [2]int{0, 0}},
	)
	hashSet := make(map[[2]int]struct{}, n*n)
	hashSet[[2]int{0, 0}] = struct{}{}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(minHeap.Array) > 0 {
		store, _ := minHeap.HeapPop()
		// fmt.Printf("%+v\n", store)
		i, j := store.position[0], store.position[1]

		if i == n-1 && j == n-1 {
			// found the bottom right so return
			return store.cost
		}

		for _, d := range directions {
			nextPosition := [2]int{i + d[0], j + d[1]}
			if _, seen := hashSet[nextPosition]; seen {
				continue
			} else if 0 <= nextPosition[0] && nextPosition[0] < n && 0 <= nextPosition[1] && nextPosition[1] < n {
				// add the next position to the visited hash set
				hashSet[nextPosition] = struct{}{}
				minHeap.HeapPush(Store{max(store.cost, grid[nextPosition[0]][nextPosition[1]]), nextPosition})
			}
		}
	}

	// unreachable code given the leetcode inputs
	return 0
}

type Store struct {
	cost     int
	position [2]int
}

/*
778. swim in rising water
100% time, 95.6% memory

This is solving the same question as above question by binary searching for the minimum height necessary to reach the bottom right
*/
func swimInWaterBinarySearch(grid [][]int) int {
	n := len(grid)
	// get the minimum and maximum heights in the grid
	minH, maxH := grid[0][0], grid[0][0]
	for row := range n {
		for col := range n {
			if grid[row][col] < minH {
				minH = grid[row][col]
			} else if grid[row][col] > maxH {
				maxH = grid[row][col]
			}
		}
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// need to predeclare since it is recursive
	var dfs func(row, col, target int) bool
	dfs = func(row, col, target int) bool {
		if row < 0 || col < 0 || row >= n || col >= n || visited[row][col] || grid[row][col] > target {
			// position is out of bounds
			// or position has been visited
			// or the height is larger than the current binary search target
			return false
		}
		if row == n-1 && col == n-1 {
			// if we reach the end position, we can return true since it is possible to reach given the constraints
			return true
		}
		// updated visited matrix
		visited[row][col] = true
		// check if any of the 4 directions can find the bottom right
		found := dfs(row+1, col, target) || dfs(row-1, col, target) || dfs(row, col+1, target) || dfs(row, col-1, target)

		return found
	}

	left, right := minH, maxH
	for left < right {
		target := (left + right) / 2
		if dfs(0, 0, target) {
			right = target
		} else {
			// need to look with a looser maximum height on path constraint
			left = target + 1
		}

		// clear the visited matrix
		for i := range visited {
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
	}

	// when the binary search finishes return the right pointer since it is the one that gets updated on success in the loop
	return right
}
