package main

import (
	"log"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			log.Println("Boomm!")
			break
		}
		log.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
