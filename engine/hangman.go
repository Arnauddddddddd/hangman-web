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
	letterInWord := false //la lettre tapé est fausse//
	for i := 0; i < len(g.mot_secret); i++ {
		if rune(g.mot_secret[i]) == letter {
			list[i*2] = letter
			letterInWord = true //la lettre tapé est vraie//
		}
	}
	if !letterInWord {
		g.try += 1
		return //rééssayer une lettre//
	}
	g.mot_cachee = ""
	for i := 0; i < len(list); i++ {
		g.mot_cachee += string(list[i])
	}
}


func (g *Structure) defWord(movie int) {
	listMovies := []string{g.mot_titanic, g.mot_alien ,g.mot_lotr, g.mot_fc,  g.mot_jurassic, g.mot_spiderman} // nom des films//
	g.mot_secret = listMovies[movie-1] //liste de film//
	
	if len(g.mot_cachee) < len(g.mot_secret)*2 {
		for i := 0; i < len(g.mot_secret); i++ {
			g.mot_cachee += "_ "
		}
	}
}


func (g *Structure) verifWin() {
	list := []rune(g.mot_cachee)
	list2 := []rune(g.mot_secret)
	for i := 0; i < len(g.mot_secret); i++ {
		if list[i*2] != list2[i] {
			return
		}
	}
	g.win = true //tu as gagné//
	g.end = true //tu as perdu//
}

