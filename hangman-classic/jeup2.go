package hangman

import (
	"fmt"
)

var CurrentUser string

func Jeux2() {
	nouvelleRune := []rune(Data.Lettre)
	x := CheckLetter(Data.MotChoisi, nouvelleRune[0])
	fmt.Printf("oui")
	Data.Stock = append(Data.Stock, Data.Lettre)
	if len(x) >= 1 { //si >1 c'est que la lettre est présente.
		for i := 0; i < len(x); i++ {
			Data.Mot[x[i]] = rune(Data.MotChoisi[x[i]] - 32)
			Data.Mot2 = string(Data.Mot)
			fmt.Printf("lettre bonne")
		}
		if TestWin(Data.Mot) {
			Data.Victory = true
			Data.Erreur = 0
			Data.Jose = ""
			Data.Stock = []string{}
			if Data.Players[CurrentUser] == 0 {
				fmt.Printf("Score de %v enregistré", CurrentUser)
				Data.Players[CurrentUser] = 1
			} else {
				fmt.Printf("Score de %v amélioré", CurrentUser)
				Data.Players[CurrentUser] += 1
			}
		}
	} else {
		Data.Erreur++
		fmt.Println("nique")
		if Data.Erreur == 1 {
			Data.Jose = "../assets/pendu1.jpg"
		} else if Data.Erreur == 2 {
			Data.Jose = "../assets/pendu2.jpg"
		} else if Data.Erreur == 3 {
			Data.Jose = "../assets/pendu3.jpg"
		} else if Data.Erreur == 4 {
			Data.Jose = "../assets/pendu4.jpg"
		} else if Data.Erreur == 5 {
			Data.Jose = "../assets/pendu5.jpg"
		} else if Data.Erreur == 6 {
			Data.Jose = "../assets/pendu6.jpg"
		} else if Data.Erreur == 7 {
			Data.Jose = "../assets/pendu7.jpg"
		} else if Data.Erreur == 8 {
			Data.Jose = "../assets/pendu8.jpg"
		} else if Data.Erreur == 9 {
			Data.Jose = "../assets/pendu9.jpg"
		} else if Data.Erreur == 10 {
			Data.Jose = "../assets/pendu10.jpg"
		}
	}
	if len(x) == 0 {
		Data.Essai = 10 - Data.Erreur
	}
}
