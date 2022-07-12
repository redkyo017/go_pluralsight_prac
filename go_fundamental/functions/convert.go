package main

import (
	"log"
	"strings"
)

func main() {
	name := "hung han"
	course := "go fund course"

	log.Println(converter(name, course))
}

func converter(s1, s2 string) (string, string) {
	return strings.ToUpper(s1), strings.Title(s2)
}
