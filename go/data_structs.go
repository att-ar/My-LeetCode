package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SegmentTreeNode struct {
	Start int
	End   int
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
}

type MapNode struct {
	Key  int
	Val  int
	Next *MapNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val  int
	Next *Node
}

type AVL interface {
	RR()
	LL()
	RL()
	LR()
}

/*
TrieNode (marker is a bool)
Would technically be more efficient to use byte (int8) instead of rune (int32)
*/
type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), end: false}
}

/* TrieNodeWord (marker is a string) */
type TrieNodeWord struct {
	children map[byte]*TrieNodeWord
	word     string
}

func NewTrieNodeWord() *TrieNodeWord {
	return &TrieNodeWord{children: make(map[byte]*TrieNodeWord), word: ""}
}
func (t *TrieNodeWord) Insert(word string) {
	node := t
	var c byte
	for i := range word {
		c = word[i]
		if _, exists := node.children[c]; !exists {
			node.children[c] = NewTrieNodeWord()
		}
		node = node.children[c]
	}
	node.word = word
}

/* TrieNodeNum (marker is a number) */
type TrieNodeNum struct {
	children map[rune]*TrieNodeNum
	number   int
}

func NewTrieNodeNum() *TrieNodeNum {
	return &TrieNodeNum{children: make(map[rune]*TrieNodeNum), number: 0}
}
func (tn *TrieNodeNum) Insert(word string) {
	curNode := tn
	for _, char := range word {
		if _, exists := curNode.children[char]; !exists {
			curNode.children[char] = NewTrieNodeNum()
		}
		curNode = curNode.children[char]
		// non-trivial part is here:
		// need to increment the number of words passing through this letter.
		curNode.number++
	}
}
func (tn *TrieNodeNum) SearchPrefixes(word string) int {
	curNode := tn
	res := 0
	for _, char := range word {
		if _, exists := curNode.children[char]; exists {
			curNode = curNode.children[char]
			res += curNode.number
		} else {
			break
		}
	}
	return res
}

/*
There is a built in heap interface with methods in "container/heap". Can try this again using that instead

See the IntHeap and PriorityQueue implementations below this.
*/
type Heap[T any] struct {
	Array   []T
	Compare func(a, b T) bool
}

func NewHeap[T any](compare func(a, b T) bool) *Heap[T] {
	return &Heap[T]{Array: []T{}, Compare: compare}
}
func (h *Heap[T]) HeapPush(value T) {
	h.Array = append(h.Array, value)
	// the index of the newly appended value if passed to heapifyUp
	h.heapifyUp(len(h.Array) - 1)
}
func (h *Heap[T]) heapifyUp(index int) {
	for index > 0 {
		// heap is an array representation of a binary tree, so the indices are set like this:
		// integer division:
		parent := (index - 1) / 2
		if h.Compare(h.Array[index], h.Array[parent]) {
			h.Array[index], h.Array[parent] = h.Array[parent], h.Array[index]
			index = parent
		} else {
			break
		}
	}
}
func (h *Heap[T]) HeapPop() (T, bool) {
	if len(h.Array) == 0 {
		var zero T // Return zero value of T if heap is empty
		return zero, false
	}

	min := h.Array[0]
	h.Array[0] = h.Array[len(h.Array)-1]
	h.Array = h.Array[:len(h.Array)-1]
	h.heapifyDown(0)
	return min, true
}
func (h *Heap[T]) heapifyDown(index int) {
	lastIndex := len(h.Array) - 1
	for {
		left, right := 2*index+1, 2*index+2
		smallest := index

		if left <= lastIndex && h.Compare(h.Array[left], h.Array[smallest]) {
			smallest = left
		}
		if right <= lastIndex && h.Compare(h.Array[right], h.Array[smallest]) {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.Array[index], h.Array[smallest] = h.Array[smallest], h.Array[index]
		index = smallest
	}
}

/* Using the std lib container/heap*/
type IntHeap []int

func (ih IntHeap) Len() int           { return len(ih) }
func (ih IntHeap) Less(i, j int) bool { return ih[i] < ih[j] }
func (ih IntHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }
func (ih *IntHeap) Push(x any) {
	*ih = append(*ih, x.(int))
}

/*
Honestly not sure why I can't just reassign the underlying array to a truncation of itself
Answer from ChatGPT: https://chatgpt.com/share/678021ad-aeec-800a-84b2-2a7c40ac19e5
*/
func (ih *IntHeap) Pop() any {
	hold := *ih
	n := len(hold)
	result := hold[n-1]
	*ih = hold[0 : n-1]
	return result
}

/*
Higher Priority means Lower Value

	for Values: [1, 5]
	Priorities: [2, 1] (anything works as long as prio for 1 is higher than that of 5)
	Priorities: [-1, -5] (this is the easiest way for integers)
	Higher prio gets popped first -> 1 gets popped first
*/
type QueueItem struct {
	Value    any
	Index    int
	Priority int
}

type PriorityQueue []*QueueItem

/*
These methods that do not take a pointer value make a copy of the `pq` struct every time pq.Method gets called
Unless you explicitly use (&pq).Method, the struct is copied
*/
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority > pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) {
	x.(*QueueItem).Index = len(*pq)
	*pq = append(*pq, x.(*QueueItem))
}
func (pq *PriorityQueue) Pop() any {
	hold := *pq
	n := len(hold)
	result := hold[n-1]
	*pq = hold[0 : n-1]
	return result
}
