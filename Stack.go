package main

import "fmt"

func main() {
	s := new(Stack)

	var ref = *s
	ref.push("a")
	ref.push("b")
	ref.push("c")
	ref.push("d")
	fmt.Println(ref.size())
	fmt.Println(ref.pop())
	fmt.Println(ref.pop())
	fmt.Println(ref.size())
}

type Stack struct {
	stack []interface{}
}

func (s *Stack) push(item interface{}) {
	s.stack = append(s.stack, item)
}

func (s *Stack) pop() interface{} {
	if (len(s.stack) == 0) {
		return nil
	}
	item := s.stack[len(s.stack) - 1]
	s.stack = s.stack[0:len(s.stack) - 1]
	return item
}

func (s *Stack) size() int {
	return len(s.stack)
}