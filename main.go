package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Reales", "Myles")
	PrintList(123, 456, 789)
	PrintList("Hola", 456, true)
}
