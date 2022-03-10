package generics

import (
	"fmt"
)

type INumber interface {
	int | int32
}

type IGeneric interface {
	PrintG()
}

type GenericT[T any, T1 INumber] struct {
	Data  T
	Digit T1
}

func (g *GenericT[T1, T2]) PrintG() {
	fmt.Printf("%T\nvalue: %+v\n\n", g, g)
}

func Generic[T any](t T) {
	fmt.Println(t)
}
