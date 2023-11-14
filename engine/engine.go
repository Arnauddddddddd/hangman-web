package engine

type Structure struct {
	running bool
	list_web_page []string
	indice_web int
}



func (g *Structure) Run() {
	g.init()
	for g.running {
		g.web()
	}
}