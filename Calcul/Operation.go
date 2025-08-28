package main

import "fmt"

func main() {
	a := 10
	b := 5

	fmt.Println("addition:", a+b)
	fmt.Println("Soustraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulo:", a%b)

	// Opérateurs de comparaison
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)
}
