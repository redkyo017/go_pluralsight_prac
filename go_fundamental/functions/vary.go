package main

import "log"

func main() {
	finishResults := []int{2, 5, 8, 9, 1, 4, 3}

	best := findTheBest(finishResults...)
	log.Println(best)
}

func findTheBest(results ...int) int {
	best := results[0]
	for _, res := range results {
		if res < best {
			best = res
		}
	}
	return best
}
