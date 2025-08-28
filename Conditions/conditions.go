package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Println("Vous êtes majeur")
	} else {
		fmt.Println("Vous êtes mineur")
	}

	// Condition avec plusieur cas
	note := 8
	if note >= 16 {
		fmt.Println("Très bien")
	} else if note >= 14 {
		fmt.Println("Bien")
	} else if note >= 12 {
		fmt.Println("Assez bien")
	} else {
		fmt.Println("Insuffisant")
	}

	// if avec initialisation
	if x := 1 * 2; x > 8 {
		fmt.Println("x est grand:", x)
	} else {
		fmt.Println("x est petit:", x)
	}
}
