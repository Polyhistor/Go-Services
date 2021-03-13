package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("the result of division is ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	result = x / y

	if y == 0 {
		return result, errors.New("Cannot devide by zero")
	}

	return result, nil
}
