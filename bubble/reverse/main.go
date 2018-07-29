
package main

import "fmt"

func bubbleSort(number []int){
	var n int = len(number)
	var i int
	for i = 0; i < n; i++{
		sweep(number)

	}
}
func sweep(numbers []int) {
	var n int = len(numbers)

	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < n {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		fmt.Println("comparin the following : ", firstNumber, secondNumber)

		if firstNumber < secondNumber{
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber

		}
		firstIndex++
		secondIndex++
	}

}

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted):", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}
