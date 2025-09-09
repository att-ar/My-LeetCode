package main

import "fmt"

func runStack() {
	fmt.Println("Valid Parentheses")
	fmt.Println(isValid("()"))
	fmt.Println(isValid("("))

	fmt.Println("Min Stack")
	// too lazy to write tests
}

// MinStack
// 60% time, 51% memory
// can replace the int64 with a Node {value, min} struct
type MinStack struct {
	Stack [][]int64
	Min   int64
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Min = int64(val)
	} else {
		this.Min = min(this.Min, int64(val))
	}
	this.Stack = append(this.Stack, []int64{int64(val), this.Min})
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	if len(this.Stack) == 0 {
		this.Min = 0
	} else {
		this.Min = this.Stack[len(this.Stack)-1][1]
	}
}

func (this *MinStack) Top() int {
	return int(this.Stack[len(this.Stack)-1][0])

}

func (this *MinStack) GetMin() int {
	return int(this.Min)
}

// Valid Parentheses
// 100% time, 73% time
func isValid(s string) bool {
	stack := []rune{}
	close := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, c := range s {
		if value, found := close[c]; found {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
