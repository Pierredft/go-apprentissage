package main

import "fmt"

func main() {
	// Boucle for classique
	for i := 1; i < 6; i++ {
		fmt.Println("Compteur:", i)
	}

	// Boucle while (avec for)
	j := 0
	for j < 3 {
		fmt.Println("J vaut:", j)
		j++
	}

	// Boucle infinie (break pour sortir)
	compteur := 1
	for {
		if compteur >= 4 {
			break
		}
		fmt.Println("Boucle infinie:", compteur)
		compteur++
	}
}
