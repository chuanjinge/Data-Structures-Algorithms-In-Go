package main

import "fmt"

func main() {
	que := new(QueueUsingStack)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println(que.Remove())
	que.Add(4)
	que.Add(5)
	que.Add(6)
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
}

type QueueUsingStack struct {
	stk1 Stack
	stk2 Stack
}

func (que *QueueUsingStack) Add(value int) {
	que.stk1.Push(value)
}

func (que *QueueUsingStack) Remove() int {
	var value int
	if que.stk2.IsEmpty() == false {
		value = que.stk2.Pop().(int)
		return value
	}

	for que.stk1.IsEmpty() == false {
		value = que.stk1.Pop().(int)
		que.stk2.Push(value)
	}

	value = que.stk2.Pop().(int)
	return value
}

type Stack struct {
	s []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() interface{} {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}