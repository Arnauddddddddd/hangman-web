package engine

import (
	"fmt"
	"html/template"
	"net/http"
)

// initialisation de la structure DrawWeb
type DrawWeb struct {
	Motsecret string
	Motcachee string
	Usingletters string
	Win bool
	Try int
	End bool
}


func(g *Structure) web() {
	// chargement de tous les répertoires présents dans "Hangman-Web"
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.Handle("/hangman/", http.StripPrefix("/hangman/", http.FileServer(http.Dir("hangman"))))
	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("javascript"))))
	http.Handle("/police/", http.StripPrefix("/police/", http.FileServer(http.Dir("police"))))
	http.Handle("/audio/", http.StripPrefix("/audio/", http.FileServer(http.Dir("audio"))))
	// chargement de toutes les pages"
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index", g.index)
	http.HandleFunc("/pageTitanic", g.pageTitanic)
	http.HandleFunc("/pageAlien", g.pageAlien)
	http.HandleFunc("/pageLOTR", g.pageLOTR)
	http.HandleFunc("/pageFightclub", g.pageFightclub)
	http.HandleFunc("/pageJurassic", g.pageJurassic)
	http.HandleFunc("/pageSpiderman", g.pageSpiderman)
	// chargement du port utilisé
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
	
}
// fonctions pour chaque page

func(g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[0]))
	tmpl.Execute(w, nil)
}

func(g *Structure) pageTitanic(w http.ResponseWriter, r *http.Request) {
	g.defWord(1)
	g.pageHangman(w, r, 1)
	
}

func(g *Structure) pageAlien(w http.ResponseWriter, r *http.Request) {
	g.defWord(2)
	g.pageHangman(w, r, 2)
}

func(g *Structure) pageLOTR(w http.ResponseWriter, r *http.Request) {
	g.defWord(3)
	g.pageHangman(w, r, 3)
}

func(g *Structure) pageFightclub(w http.ResponseWriter, r *http.Request) {
	g.defWord(4)
	g.pageHangman(w, r, 4)
}

func(g *Structure) pageJurassic(w http.ResponseWriter, r *http.Request) {
	g.defWord(5)
	g.pageHangman(w, r, 5)
}

func(g *Structure) pageSpiderman(w http.ResponseWriter, r *http.Request) {
	g.defWord(6)
	g.pageHangman(w, r, 6)
}


func (g *Structure) pageHangman(w http.ResponseWriter, r *http.Request, indice int) {
	r.ParseForm()  // permet de récupérer les caractères envoyés par les pages html
	tmpl := template.Must(template.ParseFiles(g.list_web_page[indice]))
	letter := r.Form.Get("letter")         // initialisation de la lettre actuelle
	g.currentLetter = letter	   

	if len(g.currentLetter) > 0 {          // si l'information reçue n'est pas vide(donc est une lettre) 
		if g.currentLetter == "giveup" {   // si l'information reçue est égale à "giveup", la partie est terminée car le joueur a abandonné
			g.end = true
		} else if g.verif(rune(g.currentLetter[0])) { // si la lettre n'a pas déja été utilisée : on l'ajoute aux lettres utilisées et on vérifie si elle est dans le mot
			g.usingLetters += "  " + string(g.currentLetter[0])
			g.letterTest = append(g.letterTest, string(g.currentLetter[0]))
			g.inWord(rune(g.currentLetter[0]))
		}
	}
	g.verifWin()       // on vérifie si le joueur a gagné la partie
	if g.try >= 10 {   // si le joueur a eu + de 10 essaies, la partie est terminée
		g.end = true
	}
	if g.end {    // si la partie est terminée et que l'utilisateur appuie sur le bouton "Menu" : On réinitialise tout et on relance la première page
		if g.currentLetter == "back" {
			g.init()
			g.index(w, r)
		}
	}

	web := DrawWeb {
		Motsecret: g.mot_secret,
		Motcachee: g.mot_cachee,
		Usingletters: g.usingLetters,
		Win: g.win,
		Try: g.try,
		End: g.end,
	}
	tmpl.Execute(w, web)  // permet de récupérer les valeurs de la structure DrawWeb sur le html
}