package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type integer int

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// parametros de tipos
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	PrintList("Reales", "Myles")
	PrintList(123, 456, 789)
	PrintList("Hola", 456, true)

	var num1 integer = 100
	var num2 integer = 300
	// pareametros de tipos
	fmt.Println(Sum[uint](4, 5, 6, 7))
	fmt.Println(Sum[float32](4.5, 5.4, 6.8, 7.2))
	fmt.Println(Sum(num1, num2))

}
