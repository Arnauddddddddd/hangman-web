package engine

import (
	"html/template"
	"net/http"
)


type DrawWeb struct {
	Motsecret string
	Motcachee string
	Usingletters string
	Win bool
	Try int
	End bool
}


func(g *Structure) web() {
	http.HandleFunc("/", g.index)
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.Handle("/police/", http.StripPrefix("/police/", http.FileServer(http.Dir("police"))))
	http.Handle("/audio/", http.StripPrefix("/audio/", http.FileServer(http.Dir("audio"))))
	http.HandleFunc("/pageTitanic", g.pageTitanic)
	http.HandleFunc("/pageAlien", g.pageAlien)
	http.HandleFunc("/pageLOTR", g.pageLOTR)
	http.HandleFunc("/pageFightclub", g.pageFightclub)
	http.HandleFunc("/pageJurassic", g.pageJurassic)
	http.HandleFunc("/pageSpiderman", g.pageSpiderman)
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

func(g *Structure) pageLOTR(w http.ResponseWriter, r *http.Request) {
	g.defWord("lotr")
	r.ParseForm()
	g.pageHangman(w, r, 3)
}

func(g *Structure) pageFightclub(w http.ResponseWriter, r *http.Request) {
	g.defWord("fightclub")
	r.ParseForm()
	g.pageHangman(w, r, 4)
}

func(g *Structure) pageJurassic(w http.ResponseWriter, r *http.Request) {
	g.defWord("jurassic")
	r.ParseForm()
	g.pageHangman(w, r, 5)
}

func(g *Structure) pageSpiderman(w http.ResponseWriter, r *http.Request) {
	g.defWord("spiderman")
	r.ParseForm()
	g.pageHangman(w, r, 6)
}


func (g *Structure) pageHangman(w http.ResponseWriter, r *http.Request, indice int) {
	tmpl := template.Must(template.ParseFiles(g.list_web_page[indice]))
	A, B, C, D, E, F, G :=  r.Form.Get("A"), r.Form.Get("B"), r.Form.Get("C"), r.Form.Get("D"), r.Form.Get("E"), r.Form.Get("F"), r.Form.Get("G")
	H, I, J, K, L, M, N :=  r.Form.Get("H"), r.Form.Get("I"), r.Form.Get("J"), r.Form.Get("K"), r.Form.Get("L"), r.Form.Get("M"), r.Form.Get("N")
	O, P, Q, R, S, T, U :=  r.Form.Get("O"), r.Form.Get("P"), r.Form.Get("Q"), r.Form.Get("R"), r.Form.Get("S"), r.Form.Get("T"), r.Form.Get("U")
	V, W, X, Y, Z :=        r.Form.Get("V"), r.Form.Get("W"), r.Form.Get("X"), r.Form.Get("Y"), r.Form.Get("Z")
	g.currentLetter = A+B+C+D+E+F+G+H+I+J+K+L+M+N+O+P+Q+R+S+T+U+V+W+X+Y+Z
	if len(g.currentLetter) > 0 {
		if g.verif(rune(g.currentLetter[0])) {
			g.usingLetters += "  " + string(g.currentLetter[0])
			g.letterTest = append(g.letterTest, string(g.currentLetter[0]))
			g.inWord(rune(g.currentLetter[0]))
		}
	}
	g.verifWin()
	if g.try >= 10 {
		g.end = true
	}
	if g.end {
		back := r.Form.Get("back")
		g.reset_button = back
		if len(g.reset_button) > 0 {
			g.init()
			g.index(w, r)
			return
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
	tmpl.Execute(w, web)
}