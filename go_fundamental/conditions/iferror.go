package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("test1.txt")

	if err != nil {
		log.Println("ther error is", err)
	}
}
