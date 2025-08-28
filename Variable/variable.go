package main

import "fmt"

func main() {

	// Différentes façons de déclarer des variables
	var nom string = "Pierre"
	var age int = 25
	var prix float64 = 19.99
	var actif bool = true

	// Déclaration courte (Go devine le type)
	ville := "Paris"
	temperature := 20.5

	fmt.Println("Nom:", nom)
	fmt.Println("Âge:", age)
	fmt.Println("Prix:", prix)
	fmt.Println("Actif:", actif)
	fmt.Println("Ville:", ville)
	fmt.Println("Température:", temperature)
}
