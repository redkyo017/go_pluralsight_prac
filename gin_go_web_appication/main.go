package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("con co be be")
	log.Println("con meo bay bay")
	// http.ListenAndServe(":3000", controller)
	r := registerRoutes()

	r.Run(":8080")
}
