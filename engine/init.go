package engine

import "fmt"

func (g *Structure) init() {

	g.list_web_page = []string{"templates/index.html", "templates/page2.html"}
	g.indice_web = 1
	g.running = true
	g.mot_secret = "TRACTEUR"
	g.mot_cachee = ""
	g.currentLetter = ""
	g.letterTest = []string{}
	g.try = 0
	for i := 0; i < len(g.mot_secret); i++ {
		g.mot_cachee += "_ "
	}
	fmt.Println(g.mot_cachee)
}

