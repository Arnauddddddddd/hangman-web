package engine

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Structure) init() {

	g.list_web_page = []string{"pages/index.html", "pages/pageTitanic.html", "pages/pageAlien.html", "pages/pageLOTR.html", "pages/pageFightclub.html"}
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
	g.usingLetters = " "
	g.reset_button = ""
	g.listWordsTitanic = []string{"PAQUEBOT", "ICEBERG", "NAVIRE", "AMOUR", "NAUFRAGE", "RICHE", "PAUVRE", "SAUVETAGE", "GLACE", "OCEAN", "COULER", "DIAMANT", "PASSAGER", "COQUE", "EPAVE", "CROISIERE", "CABINE", "DETRESSE", "LUXUEUX", "CANOT"}
	g.listWordsAlien = []string{"MORT", "CREATURE", "VAISSEAU", "SURVIVANT", "MEURTRIER", "ORGANISME", "LANCE_FLAMME", "ANDROIDE", "ESPACE", "REINE", "OEUF", "PROIE", "HOTE", "TUER", "MONSTRE", "PLANETE", "EXPLOSION", "DESTRUCTION", "SURVIE", "CONDUITS"}
	g.listWordsLotr = []string{"ELFE", "ORC", "DRAGON", "MAGIE", "ARMURE", "GUERRIER", "SORCIER", "CREATURE", "ARGENT", "GROTTE", "ROYAUME", "DAGUE", "CHEVAL", "MAGIE", "PERIPLE", "BATAILLE", "GOBELIN", "ANNEAU", "POUVOIR", "TERRIER", "TOUR", "BATON"}
	g.listWordsFc = []string{"SAVON", "COMBAT", "CIGARETTE", "VIOLENCE", "NARRATEUR", "ANARCHIE", "BOMBE", "INSOMNIE", "DESTRUCTION", "REGLES", "CHAOS", "HALLUCINATION", "REVOLUTION", "PSYCHANALYSE", "REGLES", "REGLES", "REGLES", "REGLES", "REGLES", "REGLES" }

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println(len(g.listWordsAlien), len(g.listWordsTitanic))
	g.mot_titanic = g.listWordsTitanic[r.Intn(len(g.listWordsTitanic))]
	g.mot_alien = g.listWordsAlien[r.Intn(len(g.listWordsTitanic))]
	g.mot_lotr = g.listWordsLotr[r.Intn(len(g.listWordsTitanic))]
	g.mot_fc = g.listWordsFc[r.Intn(len(g.listWordsFc))]
}
