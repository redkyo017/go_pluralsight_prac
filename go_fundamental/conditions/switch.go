package main

import "log"

func main() {
	switch "Kubernetes" {
	case "kubernetes":
		log.Println("case 1. kubernetes with lower case \"k\".")
		log.Println("1")
	case "Kubernetes":
		log.Println("case 2. kubernetes with a capital \"K\".")
		log.Println("2")
		fallthrough
	case "K8s":
		log.Println("case 3. kubernetes shortname \"Kates\".")
		log.Println("3")
	case "Istio":
		log.Println("case 4. service mesh is the future.")
		log.Println("4")
	default:
		log.Println("hit the default.")
	}
}
