package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		log.Println("we got the following even number", tmpNum)
	case 1, 3, 5, 7, 9:
		log.Println("we got the following odd number", tmpNum)

	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
