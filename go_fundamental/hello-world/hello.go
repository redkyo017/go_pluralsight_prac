package main

import (
	"fmt"
	"log"
	"strconv"
)

var (
	name, course string
	module       = "4"
	clip         = 2
)

func main() {
	fmt.Println("hello Pluralsight!")
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		log.Println("total of module and clip", total)
	}
}
