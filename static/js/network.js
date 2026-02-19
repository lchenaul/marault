const canvas = document.getElementById("network-canvas");
if (!canvas) return;

const ctx = canvas.getContext("2d");

let width, height, dpr;
const NODE_COUNT = 70;
const MAX_DISTANCE = 130;
let nodes = [];

const prefersReducedMotion = window.matchMedia?.("(prefers-reduced-motion: reduce)")?.matches;

function resizeCanvas() {
  dpr = window.devicePixelRatio || 1;
  width = window.innerWidth;
  height = window.innerHeight;

  canvas.width = Math.floor(width * dpr);
  canvas.height = Math.floor(height * dpr);
  canvas.style.width = `${width}px`;
  canvas.style.height = `${height}px`;

  ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
}

window.addEventListener("resize", resizeCanvas, { passive: true });
resizeCanvas();

function createNodes() {
  nodes = [];
  for (let i = 0; i < NODE_COUNT; i++) {
    nodes.push({
      x: Math.random() * width,
      y: Math.random() * height,
      vx: (Math.random() - 0.5) * 0.2,
      vy: (Math.random() - 0.5) * 0.2
    });
  }
}

createNodes();

let rafId = null;

function animate() {
  ctx.clearRect(0, 0, width, height);

  for (let i = 0; i < nodes.length; i++) {
    const a = nodes[i];

    a.x += a.vx;
    a.y += a.vy;

    if (a.x <= 0 || a.x >= width) a.vx *= -1;
    if (a.y <= 0 || a.y >= height) a.vy *= -1;

    for (let j = i + 1; j < nodes.length; j++) {
      const b = nodes[j];
      const dx = a.x - b.x;
      const dy = a.y - b.y;
      const dist = Math.sqrt(dx * dx + dy * dy);

      if (dist < MAX_DISTANCE) {
        ctx.strokeStyle = `rgba(201, 162, 77, ${1 - dist / MAX_DISTANCE})`;
        ctx.lineWidth = 1;
        ctx.beginPath();
        ctx.moveTo(a.x, a.y);
        ctx.lineTo(b.x, b.y);
        ctx.stroke();
      }
    }

    ctx.fillStyle = "#c9a24d";
    ctx.beginPath();
    ctx.arc(a.x, a.y, 2, 0, Math.PI * 2);
    ctx.fill();
  }

  rafId = requestAnimationFrame(animate);
}

function start() {
  if (prefersReducedMotion) return;
  if (rafId) cancelAnimationFrame(rafId);
  rafId = requestAnimationFrame(animate);
}

function stop() {
  if (rafId) cancelAnimationFrame(rafId);
  rafId = null;
}

document.addEventListener("visibilitychange", () => {
  if (document.hidden) stop();
  else start();
});

start();


