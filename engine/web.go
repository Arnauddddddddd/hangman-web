package engine

import (
	"fmt"
	"html/template"
	"net/http"
)


type DrawWeb struct {
	Motsecret string
	Motcachee string
	Condition bool
	Testultime int
}


func(g *Structure) web() {
	http.HandleFunc("/", g.index)
	http.HandleFunc("/page2", g.page2)
	http.ListenAndServe(":8080", nil)
	
}

func(g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[0]))
	tmpl.Execute(w, nil)
}

func(g *Structure) page2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[1]))
	r.ParseForm()
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
		g.verif(rune(g.currentLetter[0]))
	}
	nombres := DrawWeb {
		Motsecret: g.mot_secret,
		Motcachee: g.mot_cachee,
		Condition: false,
		Testultime: len(g.letterTest),
	}
	fmt.Println(g.currentLetter, g.letterTest)
	tmpl.Execute(w, nombres)
}
