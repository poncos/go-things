package main

import (
	"fmt"
)

type stack struct {
	buffer []string
}


func (s *stack) push(elem string) {
	s.buffer = append(s.buffer, elem)
}

func (s *stack) pop() string {
	if len(s.buffer) < 1 {
		panic("unable to extract element from empty stack")
	}

	// Get the last element from the buffer
	result := s.buffer[len(s.buffer)-1]

	// Remove the last element from the buffer
	s.buffer = s.buffer[:len(s.buffer)-1]

	return result
}

func (s *stack) size() int {
	return len(s.buffer)
}

func main() {
	var stackInstance stack

	stackInstance.push("element1")
	stackInstance.push("element2")

	stackSize := stackInstance.size()
	fmt.Printf("Stack size: %d\n", stackSize)
	fmt.Printf("Pop result: %s\n", stackInstance.pop())
	fmt.Printf("Stack size: %d\n", stackInstance.size())
	fmt.Printf("Pop result: %s\n", stackInstance.pop())
	fmt.Printf("Stack size: %d\n", stackInstance.size())
}
