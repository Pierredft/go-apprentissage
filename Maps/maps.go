package main

import "fmt"

func main() {
	// Créer une map
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	// Ou directement
	capitales := map[string]string{
		"France":  "Paris",
		"Italie":  "Rome",
		"Espagne": "Madrid",
	}

	fmt.Println("Ages:", ages)
	fmt.Println("Capitales:", capitales)

	//  Accéder à une valeur
	fmt.Println("Âge d'Alice:", ages["Alice"])

	// Vérifier si une clé existe
	age, existe := ages["David"]
	if existe {
		fmt.Println("Âges de David:", age)
	} else {
		fmt.Println("David n'existe pas dans la map")
	}

	// Parcourir une map
	for pays, capitale := range capitales {
		fmt.Printf("%s -> %s\n", pays, capitale)
	}

	// Supprimer un élément
	delete(ages, "Bob")
	fmt.Println("Après suppression:", ages)
}
