package main

import "log"

func main() {
	slices := make([]int, 5, 10)

	log.Printf("this slice have length of: %d and capacity of %d", len(slices), cap(slices))
}
