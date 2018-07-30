package main

import "fmt"

type Stack struct{
	slice []int
}

//push adds int to top of stack,where 0 index is bottom of stack
func (s *Stack) Push(i int){
	s.slice = append(s.slice, i)
}

//returns top item but does not remove it from stack
func (s *Stack) Peek() int{
	return s.slice[len(s.slice)-1]
}

//removes and returns top items from stack
func (s *Stack) Pop() int{
	var ret int = s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return ret
}

func main(){
	var s *Stack = new(Stack)
	s.Push(400)
	s.Push(50)
	s.Push(90)
	s.Push(98)
	fmt.Println(s)
}