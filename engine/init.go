package engine

import (
	"fmt"
	"math/rand"
	"time"
)

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
	g.listWordsTitanic = []string{"PAQUEBOT", "ICEBERG", "NAVIE", "AMOUR", "NAUFRAGE", "RICHE", "PAUVRE", "SAUVETAGE", "GLACE", "OCEAN", "COULER", "DIAMANT", "PASSAGER", "COQUE", "EPAVE", "CROISIERE", "CABINE", "DETRESSE", "LUXUEUX", "CANOT"}
	g.listWordsAlien = []string{"MORT", "CREATURE", "VAISSEAU", "SURVIVANT", "MEURTRIER", "ORGANISME", "LANCE_FLAMME", "ANDROIDE", "ESPACE", "REINE", "OEUF", "PROIE", "HOTE", "TUER", "MONSTRE", "PLANETE", "EXPLOSION", "DESTRUCTION", "SURVIE", "CONDUITS"}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println(len(g.listWordsAlien), len(g.listWordsTitanic))
	g.mot_titanic = g.listWordsTitanic[r.Intn(len(g.listWordsTitanic))]
	g.mot_alien = g.listWordsAlien[r.Intn(len(g.listWordsTitanic))]
}
