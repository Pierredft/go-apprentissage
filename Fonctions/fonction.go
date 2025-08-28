package main

import "fmt"

// Fonction simple
func saluer(nom string) {
	fmt.Println("Bonjour", nom)
}

// Fonction avec retour
func additionner(a, b int) int {
	return a + b
}

// Fonction avec plusieurs retours
func diviser(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division par zero")
	}
	return a / b, nil
}

// Fonction avec paramètres nommés
func creerPersonne(nom, prenom string, age int) (string, int) {
	nomComplet := prenom + " " + nom
	return nomComplet, age
}

func main() {
	saluer("Pierre")

	resultat := additionner(5, 3)
	fmt.Println("5 + 3 =", resultat)

	quotient, err := diviser(10, 2)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("10 / 2 =", quotient)
	}

	nom, age := creerPersonne("Dupont", "Jean", 30)
	fmt.Println("Nom complet:", nom, "Âge:", age)
}
