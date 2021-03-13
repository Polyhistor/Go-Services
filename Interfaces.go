package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Colour        string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Anny",
		Breed: "German Shepherd",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "Anthony",
		Colour:        "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorilla)
}

func (d Dog) Says() string {
	return "woof"
}

func (g Gorilla) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("this animal says", a.Says(), "and has", a.NumberOfLegs())
}
