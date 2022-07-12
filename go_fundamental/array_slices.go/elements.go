package main

import "log"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(mySlice)

	sliceOfSlice := mySlice[2:5]

	log.Println(sliceOfSlice)
}
