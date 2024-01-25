package engine

type Structure struct {
	//variables en lien avec le web
	running            bool
	pagesWeblist       []string
	iteratorWebPage    int
	closeWindow		   bool

	//variables permettant la résolution du hangman
	secretWord         string
	hiddenWord         string
	currentLetter      string
	usingLetters       string
	letterTested       []string   

    //liste des mots de chaque film
	listWordsTitanic   []string
	listWordsJurassic  []string
	listWordsAlien     []string
	listWordsLotr      []string
	listWordsFc        []string
	listWordsSpiderman []string

	//mot de chaque film
	wordTitanic        string
	wordAlien          string
	wordLOTR           string
	wordFightclub      string
	wordJurassic       string
	wordSpiderman      string

	//variables servant de changer l'état du jeu
	win                bool
	end                bool
	try                int
}

func (g *Structure) Run() {
	//lance le programme
	g.init()		
	if g.running {	 //lance en continue la fonction web
		g.web()
	}
}
