package utils

import "math/rand"

// type SomeType struct {
// 	TypeName   string
// 	TypeNumber int
// }

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
