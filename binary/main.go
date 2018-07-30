package main

import "fmt"

func main(){
	var lookingFor int = 20
	var sortedList []int = []int{1,2,4,6,7,9,10}
	fmt.Println("Looking for ", lookingFor, "in sorted list:", sortedList)

	index := binarySearch(sortedList, lookingFor)
	if index >= 0{
		fmt.Println("Found the number", lookingFor, "at:", index)
	}else {
		fmt.Println("Didnt find the no.", lookingFor, ":(")
	}
}

func binarySearch(sortedList []int, lookingFor int) int{
	var lo int = 0
	var hi int = len(sortedList)-1

	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue int = sortedList[mid]
		fmt.Println("middle value:", midValue)

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
