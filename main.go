package main

import "github.com/Mulderink1/go_1.18beta2/generics"

func main() {
	var g generics.GenericT[int, int32]
	g.Data = 44
	g.Digit = 42
	g.PrintG()

	var s generics.GenericT[string, int]
	s.Data = "Mike123"
	s.Digit = 54
	s.PrintG()

	generics.Generic(true)
}
