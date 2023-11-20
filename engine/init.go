package engine

import "fmt"

func (g *Structure) init() {

	g.list_web_page = []string{"templates/index.html", "templates/pageTitanic.html", "templates/pageAlien.html"}
	g.indice_web = 1
	g.running = true
	g.mot_secret = ""
	g.mot_cachee = ""
	g.currentLetter = ""
	g.letterTest = []string{}
	g.try = 0
	g.reset = false
	g.win = false
	g.loose = false
	g.end = false
	fmt.Println(g.mot_cachee)
}

