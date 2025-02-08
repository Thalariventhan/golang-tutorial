package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}


func main() {
	var inputStr string
	fmt.Println("Enter an operation: push, pop, peek or exit")
	stack := Stack{}

	for {
		fmt.Print("> ")
		fmt.Scanln(&inputStr)

		switch inputStr {
		case "push":
			var item int
			fmt.Print("Enter a number to push: ")
			fmt.Scanln(&item)
			stack.Push(item)
			fmt.Println("Stack after push:", stack.items)
		case "pop":
			item := stack.Pop()
			if item != -1 {
				fmt.Println("Popped item:", item)
				fmt.Println("Stack after pop:", stack.items)
			} else {
				fmt.Println("Stack is empty. Cannot pop.")
			}
		case "peek":
			item := stack.Peek()
			if item != -1 {
				fmt.Println("Top item:", item)
				fmt.Println("Stack after peek:", stack.items)
			} else {
				fmt.Println("Stack is empty.")
			}
		case "exit":
			return
		default:
			fmt.Println("Invalid operation. Try again.")
		}
	}
}
