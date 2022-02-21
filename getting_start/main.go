package main

import (
	"fmt"
	"gs/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("con co be be")
	log.Println("con meo bay bay")
	controller := controllers.NewUserController()
	http.ListenAndServe(":3000", controller)
}
