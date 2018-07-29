package main

import "fmt"

func bubbleSort(number []string){
	var n int = len(number)
	var i int
	for i = 0; i < n; i++{
		sweep(number)

	}
}
func sweep(animals []string) {
	var n int = len(animals)

	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < n {
		var firstNumber string = animals[firstIndex]
		var secondNumber string = animals[secondIndex]
		fmt.Println("comparin the following : ", firstNumber, secondNumber)

		if firstNumber < secondNumber{
			animals[firstIndex] = secondNumber
			animals[secondIndex] = firstNumber

		}
		firstIndex++
		secondIndex++
	}

}

func main() {
	var animals []string = []string{
		"dog",
		"cat",
		"alligator",
		"cheetah",
		"rat",
		"moose",
		"cow",
		"mink",
		"porcupine",
		"dung beetle",
		"camel",
		"steer",
		"bat",
		"hamster",
		"horse",
		"colt",
		"bald eagle",
		"frog",
		"rooster",
	}
	fmt.Println("Unsorted):", animals)

	bubbleSort(animals)
	fmt.Println("Sorted:", animals)
}

