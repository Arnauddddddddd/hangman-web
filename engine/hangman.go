package engine

func (g *Structure) verif(letter rune) {
	for _, i := range(g.letterTest) {
		if rune(i[0]) == letter {
			return
		}
	}
	g.letterTest = append(g.letterTest, string(letter))
}
