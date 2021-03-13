package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// mySlice := []string{"dog", "fish", "banana"}
	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	var myUserSlice []User
	myUser1 := User{FirstName: "Trevor", LastName: "Noah"}
	myUser2 := User{FirstName: "Jack", LastName: "Meng"}
	myUser3 := User{FirstName: "Will", LastName: "Smith"}

	myUserSlice = append(myUserSlice, myUser1)
	myUserSlice = append(myUserSlice, myUser2)
	myUserSlice = append(myUserSlice, myUser3)

	for i, x := range myUserSlice {
		log.Println(i, x)
	}
}
