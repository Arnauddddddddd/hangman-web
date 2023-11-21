package engine

import "fmt"

func (g *Structure) init() {

	g.list_web_page = []string{"pages/index.html", "pages/pageTitanic.html", "pages/pageAlien.html"}
	g.indice_web = 0
	g.running = true
	g.mot_secret = ""
	g.mot_cachee = ""
	g.currentLetter = ""
	g.letterTest = []string{}
	g.try = 0
	g.reset = false
	g.win = false
	g.end = false
	g.reset_button = ""
	fmt.Println(g.mot_cachee)
}
