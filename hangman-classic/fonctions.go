package hangman

import (
	"math/rand"
	"time"
)

//cette fonction permet de fournir le mot vide qui apparaît au joueur avec des "_"(95) et quelques lettres au hasard selon la longueur du mot.
func Randomletters(s []rune) []rune {
	rand.Seed(time.Now().UnixNano())
	tab := []rune(s)
	motVide := make([]rune, len(tab))
	for i := 0; i < len(motVide); i++ {
		motVide[i] = 95
	}
	for nombreLettre := len(tab)/2 - 1; nombreLettre > 0; nombreLettre-- {
		x := rand.Intn(len(tab))
		if motVide[x] == 95 {
			motVide[x] = rune(tab[x] - 32)
		} else {
			nombreLettre++
		}
	}
	return motVide
}

//on vérifie la présence de la lettre dans le mot random. Si elle n'est pas égale on regarde au cas où si ce n'est pas une voyelle correspond à un accent présent.
//on renvoie un tableau d'int qui se remplie seulement si il y a une égalité
func CheckLetter(s []rune, letter rune) []int {
	tab := []int{}
	for i := 0; i < len(s); i++ {
		if rune(s[i]) != letter {
			if letter == 101 || letter == 69 || (letter >= 232 && letter <= 235) || (letter >= 232-32 && letter <= 235-32) { //gère les e
				if s[i] >= 232 && s[i] <= 235 || s[i] == 101 {
					tab = append(tab, i)
				}
			} else if letter == 97 || letter == 65 || (letter >= 224 && letter <= 230) || (letter >= 224-32 && letter <= 230-32) { //gère les a
				if s[i] >= 224 && s[i] <= 230 || s[i] == 97 {
					tab = append(tab, i)
				}
			} else if letter == 105 || letter == 73 || (letter >= 236 && letter <= 239) || (letter >= 236-32 && letter <= 239-32) { //gère les i
				if s[i] >= 236 && s[i] <= 239 || s[i] == 105 {
					tab = append(tab, i)
				}
			} else if letter == 111 || letter == 79 || (letter >= 242 && letter <= 246) || (letter >= 242-32 && letter <= 246-32) { //gère les o
				if s[i] >= 242 && s[i] <= 246 || s[i] == 111 {
					tab = append(tab, i)
				}
			} else if letter == 117 || letter == 85 || (letter >= 249 && letter <= 252) || (letter >= 249-32 && letter <= 252-32) { //gère les u
				if s[i] >= 249 && s[i] <= 252 || s[i] == 117 {
					tab = append(tab, i)
				}
			}
		} else if rune(s[i]) == letter {
			tab = append(tab, i)
		}
	}
	return tab
}

func TestWin(tab []rune) bool { //on vérifie si il reste des "_" dans le pendu
	for i := 0; i < len(tab); i++ {
		if tab[i] == 95 {
			return false
		}
	}
	return true
}
