package main

import "log"

func main() {

	if a, b := 15, 5; a > b {
		log.Println("a larger than b")
	} else if a == b {
		log.Println("a equal to b")
	} else {
		log.Println("a smaller than b")
	}
}
