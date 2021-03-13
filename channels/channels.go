package main

import (
	"log"

	"github.com/polyhistor/amazingGo/utils"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := utils.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)
}
