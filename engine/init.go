package engine

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func (g *Structure) init() {

	g.list_web_page = []string{"pages/index.html", "pages/pageTitanic.html", "pages/pageAlien.html"}
	g.indice_web = 0
	g.running = true
	g.mot_secret = ""
	g.mot_cachee = ""
	g.currentLetter = ""
	g.letterTest = []string{}
	g.try = 0
	g.reset = false
	g.win = false
	g.end = false
	g.reset_button = ""
	g.listWordsTitanic = []string{}

	readFile, err := os.Open("/titanic.txt")
	if err != nil {
		fmt.Println(err)
}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		g.listWordsTitanic = append(g.listWordsTitanic, fileScanner.Text())
	}
	readFile.Close()
	fmt.Println(g.listWordsTitanic)

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(g.listWordsTitanic))
	g.mot_titanic = g.listWordsTitanic[random]
	g.mot_alien = ""
}
