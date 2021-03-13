package main

import "log"

func main() {
	var isTrue bool = true

	if isTrue {
		log.Println("isTrue is ", isTrue)
	}

	switch isTrue {
	case true:
		log.Println("yo")
	}

}
