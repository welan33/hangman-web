package main

import (
	"encoding/json"
	"fmt"
	hc "hangmanweb/hangman-classic"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handleRoot(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/index.html")
}

func handleOther(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/regle.html")
}

// cette "page" gère juste la fonctionnement du jeu, elle n'affiche aucun html, l'user ne le voit pas
func gameManage(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		fmt.Printf("Redirection\n")
		http.ServeFile(res, req, "./static/index.html")
	case "POST":
		fmt.Printf("Lettre reçue\n")
		req.ParseForm()
		variable := req.Form.Get("input")
		if variable == "" {
			fmt.Printf("Rentrer une lettre")
			http.Redirect(res, req, "/play", http.StatusFound)
		} else {
			fmt.Println(variable)
			hc.Data.Lettre = variable
			hc.Jeux2()
			if hc.Data.Victory {
				hc.CurrentUser = ""
				http.Redirect(res, req, "/win", http.StatusFound)
			} else if hc.Data.Erreur == 10 {
				hc.CurrentUser = ""
				http.Redirect(res, req, "/lose", http.StatusFound)
			}
		}
	}
	http.Redirect(res, req, "/play", http.StatusFound)
}

//page du jeu
func handlePlay(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Println("lol")
		jeuHtml := []string{"./static/jeu.html"}
		tmpl, _ := template.ParseFiles(jeuHtml...)
		err := tmpl.Execute(res, hc.Data)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
		}
	case "POST":
		fmt.Printf("Début de partie selon la difficulté\n")
		req.ParseForm()
		hc.Data.Difficulty = req.Form.Get("difficulty")
		println(hc.Data.Difficulty)
		hc.CurrentUser = req.Form.Get("input")
		if hc.CurrentUser == "" {
			hc.CurrentUser = "Unknown"
		}
		println(hc.CurrentUser)
		//on éxécute la première partie du code qui sélectionne le mot random et le "vide"
		hc.Jeux()
		jeuHtml := []string{"./static/jeu.html"}
		tmpl, _ := template.ParseFiles(jeuHtml...)
		err := tmpl.Execute(res, hc.Data)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
		}
	}
}

func handleWin(res http.ResponseWriter, req *http.Request) {
	if hc.Data.Difficulty != "Difficile" {
		hc.Data.Taunt = "Rejouer, pourquoi pas monter la difficulté ?"
	} else if hc.Data.Difficulty == "Difficile" {
		hc.Data.Taunt = "Rejouer, tu vas quand même pas baisser la difficulté ?"
	}
	winHtml := []string{"./static/win.html"}
	tmpl, _ := template.ParseFiles(winHtml...)
	err := tmpl.Execute(res, hc.Data)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

func handleLose(res http.ResponseWriter, req *http.Request) {
	if hc.Data.Difficulty != "Facile" {
		hc.Data.Taunt = "Un peu naze, pourquoi ne pas baisser la difficulté"
	} else if hc.Data.Difficulty == "Facile" {
		hc.Data.Taunt = "Un peu naze, malheureusement tu es déjà avec la difficulté la plus basse "
	}
	loseHtml := []string{"./static/lose.html"}
	tmpl, _ := template.ParseFiles(loseHtml...)
	err := tmpl.Execute(res, hc.Data)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

func leaderBoard(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("./static/leaderboard.html")
	tmpl.Execute(res, hc.Data)
}

func handleEnd(res http.ResponseWriter, req *http.Request) {
	jon, err := json.Marshal(hc.Data.Players)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jon))
	os.WriteFile("save.txt", jon, 0644)
	http.Redirect(res, req, "/", http.StatusFound)
	os.Exit(0)
}

func main() {

	//Gestion des routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hangman", gameManage)
	http.HandleFunc("/regle", handleOther)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/win", handleWin)
	http.HandleFunc("/lose", handleLose)
	http.HandleFunc("/end", handleEnd)

	//test
	http.HandleFunc("/leaderboard", leaderBoard)
	//test
	//retirer le /static de l'url
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fsAssets := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))

	//charger les users
	way, errt := ioutil.ReadFile("./save.txt")
	if errt != nil {
		fmt.Println(errt)
	}
	err := json.Unmarshal(way, &hc.Data.Players)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hc.Data)

	//Use the default DefaultServeMux.
	fmt.Println("Listening at http://localhost:8080")
	erry := http.ListenAndServe("localhost:8080", nil)
	if erry != nil {
		log.Fatal(erry)
	}
}
