package engine

type Structure struct {
	//variables en lien avec le web
	running            bool
	list_web_page      []string
	indice_web         int

	//variables permettant la résolution du hangman
	mot_secret         string
	mot_cachee         string
	currentLetter      string
	usingLetters       string
	letterTest         []string   

    //liste des mots de chaque film
	listWordsTitanic   []string
	listWordsJurassic  []string
	listWordsAlien     []string
	listWordsLotr      []string
	listWordsFc        []string
	listWordsSpiderman []string

	//mot de chaque film
	mot_titanic        string
	mot_alien          string
	mot_lotr           string
	mot_fc             string
	mot_jurassic       string
	mot_spiderman      string

	//variables servant de changer l'état du jeu
	win                bool
	end                bool
	try                int
}

func (g *Structure) Run() {
	//lance le programme
	g.init()		
	for g.running {	 //lance en continue la fonction web
		g.web()
	}
}
