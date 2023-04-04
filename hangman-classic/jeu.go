package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

type Page struct {
	Mot        []rune
	MotChoisi  []rune
	Erreur     int
	Lettre     string
	Difficulty string
	Victory    bool
	Mot2       string
	Jose       string
	Stock      []string
	Essai      int
	Theme      string
	Taunt      string
	MotComplet string
	Players    map[string]int
}

var Data = Page{[]rune{}, []rune{}, 0, "", "", false, "", "", []string{}, 0, "", "", "", make(map[string]int)}

func Jeux() {
	dictionnaire := []string{""}
	var diff string
	rand.Seed(time.Now().UnixNano()) //cette fonction permet de tirer une valeur au hasard selon l'heure du pc
	//gérer selon la variable difficulty
	if Data.Difficulty == "Facile" {
		diff = "./assets/words.txt"
	} else if Data.Difficulty == "Moyen" {
		diff = "./assets/words2.txt"
	} else if Data.Difficulty == "Difficile" {
		diff = "./assets/words3.txt"
	}
	content, err := os.Open(diff)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content) //on lit ligne par ligne le .txt
	for scanner.Scan() {
		dictionnaire = append(dictionnaire, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rword := dictionnaire[rand.Intn(len(dictionnaire))]
	for i := 0; len(rword) <= 0; i++ {
		rword = dictionnaire[rand.Intn(len(dictionnaire))]
	}
	Data.MotChoisi = []rune(rword) //On transforme le mot aléatoire en []rune pour faciliter la suite
	Data.Mot = Randomletters(Data.MotChoisi)
	Data.Mot2 = string(Data.Mot)
	Data.MotComplet = rword

	//reset à faire si le joueur quitte pendant une partie
	Data.Erreur = 0
	Data.Jose = ""
	Data.Stock = []string{}
	Data.Victory = false
	Data.Essai = 10
}
