package engine

type Structure struct {
	running bool
	list_web_page []string
	indice_web int
	mot_secret string
	mot_cachee string
	currentLetter string
	usingLetters string
	letterTest []string
	listWordsTitanic []string
	listWordsJurassic []string
	listWordsAlien []string
	listWordsLotr []string
	listWordsFc []string
	mot_titanic string
	mot_alien string
	mot_lotr string
	mot_fc string
	mot_jurassic string
	reset bool
	reset_button string
	win bool
	end bool
	try int
}



func (g *Structure) Run() {
	g.init()
	for g.running {
		g.web()
	}
}

