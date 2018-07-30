package main

import "fmt"

func main(){
	var lookingFor int = 6
	var sortedList []int = []int{1,2,4,6,7,9,10}
	fmt.Println("Looking for ", lookingFor, "in sorted list:", sortedList)

	index := binarySearch(sortedList, lookingFor)
	if index >= 0{
		fmt.Println("Found the number", lookingFor, "at:", index)
	}else {
		fmt.Println("Didnt find the no.", lookingFor, ":(")
	}
}

func binarySearch(numList []int, key int) int{
	low := 0
	high := len(numList) - 1

	if low <= high {

		mid := ((high + low) / 2)

		if numList[mid] > key {

			return binarySearch(numList[:mid], key)

		} else if numList[mid] < key {

			return binarySearch(numList[mid+1:], key)

		} else  if numList[mid] == key {

			return mid
		}
	}

	return -1

}

