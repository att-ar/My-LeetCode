package main

func runLinkedLists() {
	// 1. Reverse a singly linked list
	// 2. Detect a cycle in a linked list
	// 3. Merge two sorted linked lists
	// 4. Remove Nth node from the end of a linked list
	// 5. Reorder list
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find the middle of the list
	end := head
	middle := head
	for end != nil && end.Next != nil {
		end = end.Next.Next
		middle = middle.Next
	}

	// Step 2: Reverse the second half of the list
	var prev *ListNode = nil
	var hold *ListNode
	for middle != nil {
		hold = middle.Next
		middle.Next = prev
		prev = middle
		middle = hold
	}

	// Step 3: Merge the two halves
	// prev is at the L_n now
	start := head
	second := prev // the second half of the list (which is now reversed)
	var middleNext, startNext *ListNode
	for second.Next != nil {
		middleNext = second.Next
		startNext = start.Next

		start.Next = second
		second.Next = startNext

		start = startNext
		second = middleNext
	}
}
