package engine

import (
	"math/rand"
	"time"
)

func (g *Structure) init() {
	// initialisation de toutes les variables de la structure Engine

	 
	// initialisation des variables en lien avec le web
	g.list_web_page = []string{"pages/index.html", "pages/pageTitanic.html", "pages/pageAlien.html", "pages/pageLOTR.html", "pages/pageFightclub.html", "pages/pageJurassic.html", "pages/pageSpiderman.html"} //page web de chaque film//
	g.indice_web = 0
	g.running = true

	//inititialise les variables permettant la résolution du hangman
	g.mot_secret = "" 			
	g.mot_cachee = "" 			
	g.currentLetter = "" 		
	g.usingLetters = "" 		
	g.letterTest = []string{}

	//initialise les variables servant de changer l'état du jeu
	g.try = 0 					
	g.win = false 				
	g.end = false 				
	
	//initialise chaque liste de film
	g.listWordsTitanic = []string{"PAQUEBOT", "ICEBERG", "NAVIRE", "AMOUR", "NAUFRAGE", "RICHE", "PAUVRE", "SAUVETAGE", "GLACE", "OCEAN", "COULER", "DIAMANT", "PASSAGER", "COQUE", "EPAVE", "CROISIERE", "CABINE", "DETRESSE", "LUXUEUX", "CANOT"}
	g.listWordsAlien = []string{"MORT", "CREATURE", "VAISSEAU", "SURVIVANT", "MEURTRIER", "ORGANISME", "LANCE-FLAMME", "ANDROIDE", "ESPACE", "REINE-ALIEN", "OEUF", "PROIE", "HOTE", "TUER", "MONSTRE", "PLANETE", "EXPLOSION", "DESTRUCTION", "SURVIE", "CONDUITS"}
	g.listWordsLotr = []string{"ELF", "ORC", "DRAGON", "MAGIE", "ARMURE", "GUERRIER", "SORCIER", "CREATURE", "ARGENT", "GROTTE", "ROYAUME", "DAGUE", "CHEVAL", "MAGIE", "PERIPLE", "BATAILLE", "GOBELIN", "ANNEAU", "POUVOIR", "TERRIER", "TOUR", "PRECIEUX"}
	g.listWordsFc = []string{"SAVON", "COMBAT", "CIGARETTE", "VIOLENCE", "NARRATEUR", "ANARCHIE", "BOMBE", "INSOMNIE", "DESTRUCTION", "REGLES", "CHAOS", "HALLUCINATION", "REVOLUTION", "PSYCHANALYSE", "SABOTAGE", "DESILUSION", "AMNESIE", "SCHIZOPHRENIE", "REGLES", "REGLES" }
	g.listWordsJurassic = []string{"CARNIVORE", "HERBIVORE", "ATTAQUE", "PANNE", "PARC", "DEPENSER", "MOUSTIQUE", "PEUR", "ILE", "FORET", "SAVANE", "ADN", "OEUF", "DINOSAURE", "VENIN", "ECHANTILLONS", "CLONAGE", "FOSSILE", "PALEONTOLOGIE", "PREDATEUR" }
	g.listWordsSpiderman = []string{"COSTUME", "IMMEUBLE", "ARAIGNEE", "RESPONSABILITE", "POUVOIR", "TOILE", "HERO", "ONCLE-BEN", "MASQUE", "COMBAT", "VILAIN", "MORSURE", "PHOTOGRAPHE", "COURAGE", "CRIME", "VILLE", "SURHOMME", "SAUVETAGE", "EQUILIBRE", "BOUFFON" }
	
	//initialisation des variables "randoms"
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	//Initialise aléatoirement chaque mot de film en un mot aléatoire en fonction du film
	g.mot_titanic = g.listWordsTitanic[r.Intn(20)]
	g.mot_alien = g.listWordsAlien[r.Intn(20)]
	g.mot_lotr = g.listWordsLotr[r.Intn(20)]
	g.mot_fc = g.listWordsFc[r.Intn(20)]
	g.mot_jurassic = g.listWordsJurassic[r.Intn(20)]
	g.mot_spiderman = g.listWordsSpiderman[r.Intn(20)]
}
