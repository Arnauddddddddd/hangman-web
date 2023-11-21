package engine

type Structure struct {
	running bool
	list_web_page []string
	indice_web int
	mot_secret string
	mot_cachee string
	currentLetter string
	letterTest []string
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

