package main

import "fmt"

func runTrees() {
	fmt.Printf("\nTrees\n-----\n")
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 9},
			Left:  &TreeNode{Val: 90, Right: &TreeNode{Val: 12}},
		},
		Right: &TreeNode{
			Val:  90,
			Left: &TreeNode{Val: -3},
		},
	}
	// fmt.Println("treeSum")
	// fmt.Println(treeSumDFSRecurse(root))
	// fmt.Println(treeSumDFSIter(root))
	// fmt.Println(treeSumBFSIter(root))

	// fmt.Println("invertTree")
	// invertedRoot := invertTree(root)
	// fmt.Println(invertedRoot.Left.Right)

	fmt.Println("maxDepth")
	fmt.Println(maxDepth(root))

	fmt.Println("MyCalendar")
	events := [][2]int{{10, 15}, {15, 25}, {5, 10}}
	calendar := CalendarConstructor()
	possible := true
	for _, event := range events {
		if !calendar.Book(event[0], event[1]) {
			possible = false
			fmt.Println(possible)
			break
		}
	}
	if possible {
		fmt.Println(possible)
	}

	fmt.Println("Find Largest Value in Each Tree Row")
	fmt.Println(largestValues(
		&TreeNode{1,
			&TreeNode{3,
				&TreeNode{5, nil, nil},
				&TreeNode{3, nil, nil},
			},
			&TreeNode{2,
				nil,
				&TreeNode{9, nil, nil},
			}},
	))

	fmt.Println("Maximum Number of K-Divisible Components")
	fmt.Println(maxKDivisibleComponents(
		5, [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 8, 1, 4, 4}, 6,
	))
}

/*
2872. Maximum Number of K-Divisible Components
93% time, 88% memory (array of arrays graph)
37% time, 44% memory (hashmap of arrays graph)

HINT:
  - Set root to node 0
  - If a leaf node is not divisible by k, it must be in the same component as its parent node so we merge it with its parent node.
*/
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	tree := make([][]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	/*
		Will DFS and merge a leaf with its parent if the value of a leaf node is not divisible by k

		Post-Order Processing

		Returns the number of k-divisible components there are in the children of the current node
	*/
	var dfs func(node int, parent int) int
	dfs = func(node int, parent int) int {
		// base case: leaf node
		// the explicit check for != 0 is because we set that as our root node in the last line of the question's function
		// if 0 is a leaf node, this doesn't work
		if node != 0 && len(tree[node]) == 1 {
			// 2 possibilities:
			if values[node]%k == 0 {
				// 1. the leaf is divisible by k and we can get keep it
				return 1
			} else {
				// 2. the leaf is not divisible by k and we need to merge it into its parent
				values[parent] += values[node]
				return 0
			}
		}

		var result int
		// do the dfs
		for _, child := range tree[node] {
			if child != parent {
				result += dfs(child, node)
			}
		}

		// check the same condition POST-ORDER
		if values[node]%k == 0 {
			// 1. the leaf is divisible by k and we can get keep it
			return result + 1
		} else {
			// 2. the leaf is not divisible by k and we need to merge it into its parent
			values[parent] += values[node]
			return result
		}
	}
	// setting the root to 0
	return dfs(0, -1)
}

/*
515. Find Largest Value in Each Tree Row
100% time, 88% memory
it's just a BFS lol
*/
func largestValues(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	stack := []*TreeNode{root}
	var node *TreeNode
	var rowMax int
	for len(stack) > 0 {
		rowMax = stack[0].Val
		for range len(stack) {
			node = stack[0]
			stack = stack[1:]
			if node.Val > rowMax {
				rowMax = node.Val
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, rowMax)
	}

	return result
}

// 729. My Calendar
// 70% time, 54% memory
// 2 solutions, BST (Segment Tree) or Array Indexing
// copy of the SegmentTreeNode in data_structs.go
type MyCalendar struct {
	Start int
	End   int
	Left  *MyCalendar
	Right *MyCalendar
}

func CalendarConstructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	cur := mc
	for {
		if (start >= cur.End) || (end <= cur.Start) {
			// not overlapping
			if start >= cur.End {
				// move down to the right
				if cur.Right != nil {
					cur = cur.Right
				} else {
					cur.Right = &MyCalendar{Start: start, End: end}
					return true
				}
			} else {
				// move down to the left
				if cur.Left != nil {
					cur = cur.Left
				} else {
					cur.Left = &MyCalendar{Start: start, End: end}
					return true
				}
			}
		} else {
			return false
		}
	}
}

// // My Calendar II
// // doing this in python first
// type MyCalendarTwo struct {
// }

// func CalendarTwoConstructor() MyCalendarTwo {

// }

// func (mc *MyCalendarTwo) Book(start int, end int) bool {

// }

// 63% time, 11% memory
// it seems that DFS is faster than BFS because it uses less memory and deals with a smaller overhead
func maxDepth(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	// bfs works better for me mentally because I don't need to check too many paths
	// dfs also works but you need to keep track of a max res
	// as opposed to just incrementing
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		res++
		for range stack {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return res
}

// 74% time, 53% memory
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}
}

func treeSumDFSRecurse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + treeSumDFSRecurse(root.Left) + treeSumDFSRecurse(root.Right)
}

func treeSumDFSIter(root *TreeNode) int {
	res := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// iterative DFS is a last in first out game (pop Right)
		// always add the Left node last and it will only process Left all the way down
		// untl it has to pop back up to a Right node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res += node.Val
		stack = append(stack, node.Right, node.Left)
	}
	return res
}

func treeSumBFSIter(root *TreeNode) int {
	res := 0
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		for range len(stack) {
			node := stack[0]
			stack = stack[1:]
			if node == nil {
				continue
			}
			res += node.Val
			stack = append(stack, node.Left, node.Right)
		}
	}
	return res
}

func isBalanced(root *TreeNode) {

}
