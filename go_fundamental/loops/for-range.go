package main

import "log"

func main() {
	course := []string{
		"course A",
		"course B",
		"course C",
		"course D",
	}
	for _, i := range course {
		log.Println("con co be be", i)
	}
}
