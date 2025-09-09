package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func runIntervals() {
	// new returns a pointer
	// make returns the actual object
	fmt.Println(("Non-overlapping intervals"))
	intervals := [][]int{{1, 2}, {1, 3}, {2, 3}, {2, 5}, {3, 4}}
	fmt.Println(eraseOverlapIntervals((intervals)))
	intervals = [][]int{{1, 4}, {2, 3}, {1, 3}}
	fmt.Println(eraseOverlapIntervals((intervals)))
	intervals = [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}
	fmt.Println(eraseOverlapIntervals((intervals)))
}

/*
98 % time, 93 % memory

My thinking is pretty straightforward: (Greedy algorithm)
You must remove an interval if it overlaps.
So just go through the sorted array and "remove" it.
Need to keep track of removed ones so that we don't double count overlaps.

1st thinking:
I think my first attempt will just be to keep track of the largest right pointer from the removed intervals.
Since I am sorting by the first pointer. Then I can assume that anything smaller than the stored max has been checked

2nd thinking:
keep track of the largest right pointer from the NON-removed intervals lol.
*/
func eraseOverlapIntervals(intervals [][]int) int {
	// sort by the start of the intervals
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	result := 0
	// store the current largest number of the list of intervals (excluding removed intervals)
	// init with the smallest number in the array so that the first iteration works properly
	currentEnd := intervals[0][1]

	for _, interval := range intervals[1:] {
		if interval[0] < currentEnd {
			// overlap
			result += 1
			// update currentEnd to the value with the smallest right pointer
			// (we want to get rid of the larger one because it will affect more intervals)
			currentEnd = min(interval[1], currentEnd)
		} else {
			// move to the next interval if there is no overlap
			currentEnd = interval[1]
		}
	}
	return result
}

/* 1851. Minimum Interval to Include Each Query */
func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	queriesWithIdx := make([][2]int, len(queries))
	for i, q := range queries {
		queriesWithIdx[i] = [2]int{q, i}
	}
	sort.Slice(queriesWithIdx, func(i, j int) bool {
		return queriesWithIdx[i][0] < queriesWithIdx[j][0]
	})

	comparator := func(a, b interface{}) int {
		pair1 := a.([2]int)
		pair2 := b.([2]int)
		if pair1[0] != pair2[0] {
			if pair1[0] < pair2[0] {
				return -1
			}
			return 1
		}
		return 0
	}

	pq := priorityqueue.NewWith(comparator)
	res := make([]int, len(queries))
	i := 0

	for _, qPair := range queriesWithIdx {
		q, originalIdx := qPair[0], qPair[1]

		for i < len(intervals) && intervals[i][0] <= q {
			size := intervals[i][1] - intervals[i][0] + 1
			pq.Enqueue([2]int{size, intervals[i][1]})
			i++
		}

		for !pq.Empty() {
			if top, _ := pq.Peek(); top.([2]int)[1] < q {
				pq.Dequeue()
			} else {
				break
			}
		}

		if !pq.Empty() {
			if top, _ := pq.Peek(); true {
				res[originalIdx] = top.([2]int)[0]
			}
		} else {
			res[originalIdx] = -1
		}
	}

	return res
}
