package main

import "log"

func main() {
	a := 15
	b := 5

	if a > b {
		log.Println("a larger than b")
	} else if a == b {
		log.Println("a equal to b")
	} else {
		log.Println("a smaller than b")
	}
}
