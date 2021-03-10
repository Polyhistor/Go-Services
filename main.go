package main

import "log"

func main() {
	var whatToSay string = "yo"
	var dude string
	var result string
	var result2 string
	var i int

	result, dude = saySomething(whatToSay)
	log.Print(result + dude)

	result2, dude = saySomething("something else")
	log.Print(result2 + dude)

	i = 7
	log.Println(i)
}

func saySomething(s string) (string, string) {
	return s, "dude"
}
