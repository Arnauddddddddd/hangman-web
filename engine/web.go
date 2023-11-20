package engine

import (
	"fmt"
	"html/template"
	"net/http"
)


type DrawWeb struct {
	Motsecret string
	Motcachee string
	Win bool
	Loose bool
	Try int
	End bool
}


func(g *Structure) web() {
	http.HandleFunc("/", g.index)
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.HandleFunc("/pageTitanic", g.pageTitanic)
	http.HandleFunc("/pageAlien", g.pageAlien)
	http.ListenAndServe(":8080", nil)
	
}

func(g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[0]))
	tmpl.Execute(w, nil)
}

func(g *Structure) pageTitanic(w http.ResponseWriter, r *http.Request) {
	g.defWord("titanic")
	r.ParseForm()
	g.pageHangman(w, r, 1)
}

func(g *Structure) pageAlien(w http.ResponseWriter, r *http.Request) {
	g.defWord("alien")
	r.ParseForm()
	g.pageHangman(w, r, 2)
}


func (g *Structure) pageHangman(w http.ResponseWriter, r *http.Request, indice int) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[indice]))
	A := r.Form.Get("A")
	B := r.Form.Get("B")
	C := r.Form.Get("C")
	D := r.Form.Get("D")
	E := r.Form.Get("E")
	F := r.Form.Get("F")
	G := r.Form.Get("G")
	H := r.Form.Get("H")
	I := r.Form.Get("I")
	J := r.Form.Get("J")
	K := r.Form.Get("K")
	L := r.Form.Get("L")
	M := r.Form.Get("M")
	N := r.Form.Get("N")
	O := r.Form.Get("O")
	P := r.Form.Get("P")
	Q := r.Form.Get("Q")
	R := r.Form.Get("R")
	S := r.Form.Get("S")
	T := r.Form.Get("T")
	U := r.Form.Get("U")
	V := r.Form.Get("V")
	W := r.Form.Get("W")
	X := r.Form.Get("X")
	Y := r.Form.Get("Y")
	Z := r.Form.Get("Z")
	g.currentLetter = A+B+C+D+E+F+G+H+I+J+K+L+M+N+O+P+Q+R+S+T+U+V+W+X+Y+Z
	if len(g.currentLetter) > 0 {
		if g.verif(rune(g.currentLetter[0])) {
			g.letterTest = append(g.letterTest, string(g.currentLetter[0]))
			g.inWord(rune(g.currentLetter[0]))
		}
	}
	g.verifWin()
	if g.try >= 10 {
		g.loose = true
		g.end = true
	}

	nombres := DrawWeb {
		Motsecret: g.mot_secret,
		Motcachee: g.mot_cachee,
		Win: g.win,
		Loose: g.loose,
		Try: g.try,
		End: g.end,
	}
	fmt.Println(g.currentLetter, g.letterTest, g.reset)
	tmpl.Execute(w, nombres)
}