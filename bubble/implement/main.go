package main

import "fmt"

func bubbleSort(number []int){
	var n int = len(number)
	var i int
	for i = 0; i < n; i++{
		if !sweep(number,i){
			return
		}
	}
}
func sweep(numbers []int, prevArgs int) bool{
	var n int = len(numbers)

	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false
	for secondIndex < n - prevArgs{
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		fmt.Println("comparin the following : ", firstNumber, secondNumber)

		if firstNumber > secondNumber{
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted):", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}
