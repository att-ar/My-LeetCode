package main

import (
	"container/list"
	"fmt"
)

func runGraph() {
	fmt.Printf("\nGraph\n-----\n")
	// fmt.Println("Number of Provinces")
	// fmt.Println(findCircleNum(
	// 	[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
	// ))
	// fmt.Println(findCircleNum([][]int{
	// 	{1, 1, 1, 1},
	// 	{1, 1, 0, 1},
	// 	{1, 0, 1, 0},
	// 	{1, 1, 0, 1},
	// }))
	// fmt.Println(findCircleNum([][]int{
	// 	{1, 0, 0, 1},
	// 	{0, 1, 1, 0},
	// 	{0, 1, 1, 1},
	// 	{1, 0, 1, 1},
	// }))

	// fmt.Println("Min Height Trees")
	// fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	// fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	// fmt.Println(findMinHeightTrees(1, [][]int{}))

	fmt.Println("Redundant Connection")
	fmt.Println(findRedundantConnection([][]int{{16, 25}, {7, 9}, {3, 24}, {10, 20}, {15, 24}, {2, 8}, {19, 21}, {2, 15}, {13, 20}, {5, 21}, {7, 11}, {6, 23}, {7, 16}, {1, 8}, {17, 20}, {4, 19}, {11, 22}, {5, 11}, {1, 16}, {14, 20}, {1, 4}, {22, 23}, {12, 20}, {15, 18}, {12, 16}}))
}

/*
684. Redundant Connection
100% time, 38% memory

Basically you want to reduce your search space to the nodes in the tree that are part of the cycle.
You can use Kahn's algo to find all nodes that are part of the cycle (there can only be 1 since the graph is connected)

	The graph is connected because it started off as a tree before the extra edge was added.

Once you get the nodes that are part of the cycle you can just iterate through and remove the last edge for which both nodes are in the cycle.

My python solution was union-find.
I think union-find is the faster algo.
*/
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	inDegree := make([]int, n+1) // vertices are 1-indexed
	graph := make([][]int, n+1)

	// build graph and indegree map for Kahn's algorithm
	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)

		inDegree[a]++
		inDegree[b]++
	}

	// make the starting deque (container/list is backed by a doubly-linked list, there are array backed alternatives)
	deque := list.New()
	// technically node == 0 can be skipped since they are 1-indexed
	for node := range inDegree {
		if inDegree[node] == 1 {
			deque.PushBack(node)
			// set it to 0 so that it doesn't get appended in the future by accident.
			inDegree[node]--
		}
	}

	// now do the topological sort until you see the deque is empty.
	// this will occur when all nodes remaining to be processed are nodes that are part of the cycle
	// (these nodes will always have an indegree higher than 1 since there are always at two nodes attached to them in the cycle)
	var node int
	for deque.Len() > 0 {
		// type assertion since it doesn't support generics
		node = deque.Remove(deque.Front()).(int)
		for _, child := range graph[node] {
			inDegree[child]--
			if inDegree[child] == 1 {
				deque.PushBack(child)
			}
		}
	}
	// now all values in inDegree > 0 will be the nodes in the cycle.

	// isCycleNode := make([]bool, n+1)
	// // technically node == 0 can be skipped since they are 1-indexed
	// for node := range inDegree {
	// 	// we need greater than 1 because the nodes leading into the cycle will not be able to get rid of the indegree they get from the cycle node
	// 	if inDegree[node] > 1 {
	// 		isCycleNode[node] = true
	// 	}
	// }

	// now go through the edges in reverse and return when you find an edge that is unecessary
	for i := n - 1; i >= 0; i-- {
		// instead of making another array `isCycleNode` just gonna do the check here
		if inDegree[edges[i][0]] > 1 && inDegree[edges[i][1]] > 1 {
			return edges[i]
		}
	}
	// for the compiler, this code is unreachable given the question's guarantees.
	return []int{0, 0}
}

/*
310. Minimum Height Trees
79% time, 74% memory (using a map (hashset) as the inDegree map and copying the stack on every iteration instead of using `countProcessed`)
95% time, 87% memory (using a slice as the inDegree map and using `countProcessed` to know when to return instead of copying `stack` on every iteration)
Simplest way is to get the height of all the different trees (rooted at all nodes)

Can I do it in O(v + e) instead of O(v * e)?

One of the topics is Topological Sort.
If you use Kahn's algo to get all the leaf nodes, then you BFS up while adding the inDegree==1 nodes (1 since it is undirected)
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		// edge case: no BFS possible
		return []int{0}
	}
	tree := make([][]int, n)
	inDegree := make([]int, n)

	for i := range n {
		tree[i] = []int{}
	}
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
		inDegree[edge[0]]++
		inDegree[edge[1]]++
	}

	stack := []int{}
	// Topological Sort
	// init the Leaf stack
	for node := range n {
		if inDegree[node] == 1 {
			stack = append(stack, node)
			// modify the inDegree to Zero since it is a leaf node
			inDegree[node]--
		}
	}
	// BFS
	// this will store the result (whenever the loop breaks, wtv is in here will be the answer)
	// var prevStack []int
	// will keep track of the number of nodes processed instead of copying the stack every time
	countProcessed := 0
	var node int
	for len(stack) > 0 {
		// make a copy of the current stack
		// prevStack = append([]int{}, stack...)
		m := len(stack)
		countProcessed += m
		if countProcessed == n {
			// all nodes will be processed so just return the stack since it contains the answer root nodes
			return stack
		}
		for range m {
			node = stack[0]
			stack = stack[1:]

			for _, neigh := range tree[node] {
				inDegree[neigh]--
				if inDegree[neigh] == 1 {
					// if these end up as the children of other nodes later on
					// they won't get re-appended because their indegree will be set to zero by the `inDegree[neigh]--` call
					stack = append(stack, neigh)
				}
			}
		}
	}
	// just for the compiler, this code is unreachable
	return []int{}
}

/*
547. Number of Provinces
100% time, 24% memory

  - Union Find and then return the number of distinct parents pre much
  - Go through the adjacency matrix and run union on them
  - init the parents array where each node is a parent of itself
*/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parents := make([][2]int, n)
	for i := range n {
		parents[i] = [2]int{i, 0}
	}

	// only need to read the upper or lower triangles of the adjacency matrix
	for i := range n {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				// if nodes are neighbors, union their parents
				union(parents, i, j)
			}
		}
	}

	uniqueParents := 0
	seenParents := make(map[int]struct{}, n)
	var parent int
	for i := range n {
		// find the parent of the current index
		parent = find(parents, i)[0]
		if _, exists := seenParents[parent]; !exists {
			// if the parent has not been seen before, we found a new connected component
			uniqueParents++
			seenParents[parent] = struct{}{}
		}
	}
	return uniqueParents
}
