package main

import "log"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	log.Println(mySlice)

	for _, i := range mySlice {
		log.Println(i)
	}

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)

	log.Printf("my new slice  %d have length of %d and capacity of %d", mySlice, len(mySlice), cap(mySlice))
}
