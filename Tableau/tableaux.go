package main

import "fmt"

func main() {
	// Tableau de taille fixe
	var nombres [5]int
	nombres[0] = 10
	nombres[1] = 20
	nombres[2] = 30
	nombres[3] = 40
	nombres[4] = 50

	// Déclaration avec valeurs
	fruits := [3]string{"pomme", "banane", "orange"}

	fmt.Println("Nombres:", nombres)
	fmt.Println("Fruits:", fruits)
	fmt.Println("Premier fruit:", fruits[0])
	fmt.Println("Longueur:", len(fruits))
}
