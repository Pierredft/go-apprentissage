package main

import "fmt"

func main() {
	// Slice (taille dynamique)

	var scores []int
	scores = append(scores, 10, 20, 30)

	// Ou directement
	noms := []string{"Alice", "Bob", "Charlie"}

	fmt.Println("Scores:", scores)
	fmt.Println("Noms:", noms)

	//  Parcourir avec range
	for i, nom := range noms {
		fmt.Printf("Index %d: %s\n", i, nom)
	}

	// Parcourir seulement les valeurs
	for _, score := range scores {
		fmt.Println("Score:", score)
	}
}
