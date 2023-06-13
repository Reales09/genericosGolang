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

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {

		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	// PrintList("Reales", "Myles")
	// PrintList(123, 456, 789)
	// PrintList("Hola", 456, true)

	// var num1 integer = 100
	// var num2 integer = 300
	// // pareametros de tipos
	// fmt.Println(Sum[uint](4, 5, 6, 7))
	// fmt.Println(Sum[float32](4.5, 5.4, 6.8, 7.2))
	// fmt.Println(Sum(num1, num2))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}

	// fmt.Println(Includes(strings, "a"))
	// fmt.Println(Includes(strings, "f"))
	// fmt.Println(Includes(numbers, 4))
	// fmt.Println(Includes(numbers, 5))

	fmt.Println(filter(numbers, func(value int) bool { return value > 2 }))
	fmt.Println(filter(strings, func(value string) bool { return value > "b" }))

}
