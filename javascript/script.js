var audios = [[new Audio("../audio/alien.mp3"), false], [new Audio("../audio/titanic.mp3"), false] , [new Audio("../audio/lotr.mp3"), false], [new Audio("../audio/fightclub.mp3"), false], [new Audio("../audio/jurassic.mp3"), false], [new Audio("../audio/spiderman.mp3"), false]]
var pages = ["pageAlien", "pageTitanic", "pageLOTR", "pageFightclub", "pageJurassic", "pageSpiderman"]


function openPage(i) {
  for(let i = 0; i < audios.length; i++) {
    if (audios[i][1]) {
      audios[i][0].currentTime = 0;
      audios[i][0].pause();
    }
  }
  audios[i][0].play();
  audios[i][1] = true;
  window.open(pages[i])
  return;
}



