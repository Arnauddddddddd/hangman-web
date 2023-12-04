package engine

import "fmt"

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


func (g *Structure) defWord(movie string) {
	if movie == "alien" {
		g.mot_secret = g.mot_alien
	} else if movie == "titanic" {
		g.mot_secret = g.mot_titanic
	} else if movie == "lotr" {
		g.mot_secret = g.mot_lotr
	} else if movie == "fightclub" {
		g.mot_secret = g.mot_fc
	} else if movie == "jurassic" {
		g.mot_secret = g.mot_fc
	}
	if len(g.mot_cachee) < len(g.mot_secret)*2 {
		for i := 0; i < len(g.mot_secret); i++ {
			g.mot_cachee += "_ "
		}
	}
}


func (g *Structure) verifWin() {
	list := []rune(g.mot_cachee)
	list2 := []rune(g.mot_secret)
	fmt.Println(list, list2)
	for i := 0; i < len(g.mot_secret); i++ {
		if list[i*2] != list2[i] {
			return
		}
	}
	g.win = true
	g.end = true
}
