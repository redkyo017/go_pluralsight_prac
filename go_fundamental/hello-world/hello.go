package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	name := os.Getenv("USER")
	course := "go fund"
	module := "4"
	clip := 2

	fmt.Println("hello Pluralsight!")
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Clip is of type", reflect.TypeOf(clip))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	fmt.Println("memory address of *course* variable is", &course)
	var ptr *string = &course
	fmt.Println("pointing course variable at address", ptr, "which notes this value", *ptr)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		log.Println("total of module and clip", total)
	}
	log.Println("course before update", course)
	updateCourse(&course)
	log.Println("course after update", course)
}

func updateCourse(course *string) string {
	*course = "new course"
	log.Println("update course to", *course)
	return *course
}
