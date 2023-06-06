package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type integer int

// parametros de tipos
func Sum[T ~int | ~float64](nums ...T) T {
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
	fmt.Println(Sum(4, 5, 6, 7))
	fmt.Println(Sum(4.5, 5.4, 6.8, 7.2))
	fmt.Println(Sum(num1, num2))

}
