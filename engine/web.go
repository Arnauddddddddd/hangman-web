package engine

import (
	"math/rand"
	"html/template"
	"net/http"
)


type NombrePair struct {
	Nombre string
	Parite bool
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

	nbbool := false
	nbAlea := rand.Intn(6)+1
	if nbAlea % 2 == 0 {
		nbbool = true
	}

	nombres := NombrePair {
		Nombre: r.FormValue("page2"),
		Parite: nbbool,
	}
	
	tmpl.Execute(w, nombres)
}