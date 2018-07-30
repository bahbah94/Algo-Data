package main

import "fmt"

type Queue struct{
	slice []int
}

//adds int to back of the queue
func (q *Queue) Enqueue(i int){
	q.slice = append(q.slice, i)
}

//removes tge first item from queue
//returns the first item in the queue
func (q *Queue) Dequeue() int{
	var ret int = q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret
}
func main(){
	var q *Queue = new(Queue)
	q.Enqueue(123)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
}