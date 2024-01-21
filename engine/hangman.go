package engine

func (g *Structure) verif(letter rune) bool {
	// renvoie vrai si la lettre passée en paramètre a déja été testé, sinon faux
	for _, i := range(g.letterTested) {
		if rune(i[0]) == letter {
			return false
		}
	}
	return true
}


func  (g *Structure) inWord(letter rune) {
	// vérifie si la lettre passée en paramètre se trouve dans le mot secret
	list := []rune(g.hiddenWord)
	letterInWord := false 		  // initialisation d'un boolean indiquant si la lettre est dans le mot secret
	for i := 0; i < len(g.secretWord); i++ {
		if rune(g.secretWord[i]) == letter {
			list[i*2] = letter
			letterInWord = true   // la lettre tapée est dans le mot secret
		}
	}
	if !letterInWord {  // si la lettre n'est pas dans le mot, on ajoute un essai et on arrète la fonction ici.
		g.try += 1
		return
	}					// si la lettre est dans le mot, on actualise l'affichage du mot sur la page web
	g.hiddenWord = ""
	for i := 0; i < len(list); i++ {
		g.hiddenWord += string(list[i])
	}
}


func (g *Structure) defWord(movie int) {
	// définit quelle doit être le mot du film en fonction de l'entier reçu en paramètre
	listMovies := []string{g.wordTitanic, g.wordAlien ,g.wordLOTR, g.wordFightclub,  g.wordJurassic, g.wordSpiderman} // initialisation d'une liste qui contient tous les mots de chaque film
	g.secretWord = listMovies[movie-1] // initialise le mot secret en fonction de l'entier reçu en paramètre
	// affiche des "_" en fonction du nombre de lettres du mot secret
	if len(g.hiddenWord) < len(g.secretWord)*2 {
		for i := 0; i < len(g.secretWord); i++ {
			if byte(g.secretWord[i]) == '-' {
				g.hiddenWord += "- "
			} else {
				g.hiddenWord += "_ "
			}
		}
	}
}


func (g *Structure) verifWin() {
	// vérifie si le joueur a gagné la partie
	list := []rune(g.hiddenWord)
	list2 := []rune(g.secretWord)
	for i := 0; i < len(g.secretWord); i++ { // si le mot qui se complète n'est pas égale au mot secret, la fonction s'arrête
		if list[i*2] != list2[i] {
			return
		}
	}
	g.win = true // si la fonction ne s'arrête pas, le joueur a gagné
	g.end = true // la partie est terminée
}

