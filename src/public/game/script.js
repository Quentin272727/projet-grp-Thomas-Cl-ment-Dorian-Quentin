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
