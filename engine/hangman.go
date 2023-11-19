package engine

func (g *Structure) verif(letter rune) bool {
	for _, i := range(g.letterTest) {
		if rune(i[0]) == letter {
			return false
		}
	}
	return true
}


func  (g *Structure) inWord(letter rune) {
	list := []rune(g.mot_cachee)
	letterInWord := false
	for i := 0; i < len(g.mot_secret); i++ {
		if rune(g.mot_secret[i]) == letter {
			list[i*2] = letter
			letterInWord = true
		}
	}
	if !letterInWord {
		g.try += 1
		return
	}
	g.mot_cachee = ""
	for i := 0; i < len(list); i++ {
		g.mot_cachee += string(list[i])
	}
}