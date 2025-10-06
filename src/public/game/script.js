let coin = document.getElementById("mouse__coin");
let w = document.getElementsByClassName("win");

if (w.length >= 1) {
  for (let i = 0; i <= 200; i++) {
    createConfetti();
  }
}

function createConfetti() {
  let p = document.getElementsByClassName("popup");
  let c = document.createElement("div");
  let colors = ["red", "#00ff0d", "#ffd900", "#3700ff", "#ff00f2"];

  c.className = "confetti";
  c.style.top = "-50vh";
  c.style.left = `${Math.floor(Math.random() * 100)}vw`;
  c.style.backgroundColor = colors[Math.floor(Math.random() * colors.length)];
  c.style.animationDuration = `${Math.random() * 1}s`;
  c.style.animationDelay = `${Math.random() * 2}s`;
  p[0].append(c);
}

// Make the coin slide
document.addEventListener("mousemove", (m) => {
  coin.style.transform = `translate(${calcCoinPos(m.x)}px, 0)`;
  //coin.style.marginLeft = calcCoinPos(m.x) + "px";
});

function calcCoinPos(cursorPosX) {
  let cursorPosMin = window.innerWidth / 2 - coin.parentElement.clientWidth / 2;
  let cursorPosMax = window.innerWidth / 2 + coin.parentElement.clientWidth / 2;
  let cursorPos = cursorPosX;

  if (cursorPosX < cursorPosMin) {
    return 0 - coin.clientWidth / 2;
  }

  if (cursorPosX >= cursorPosMin && cursorPosX <= cursorPosMax) {
    return cursorPosX - cursorPosMin - coin.clientWidth / 1.5;
  }

  if (cursorPosX > cursorPosMax) {
    return cursorPosMax - cursorPosMin - coin.clientWidth / 1.5;
  }

  return cursorPos;
}
