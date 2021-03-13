package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "yo")
	mySlice = append(mySlice, "yo")
	mySlice = append(mySlice, "yo")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sort.Ints(numbers)
	log.Println(numbers)
}
