package main

import "log"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}
	testMap["C"] = 100
	testMap["H"] = 99
	for key, value := range testMap {
		log.Printf("key is :%v and value is :%v", key, value)
	}
}
