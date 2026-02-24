const toggle = document.getElementById("mi-chat-toggle");
const closeBtn = document.getElementById("mi-chat-close");
const panel = document.getElementById("mi-chat-panel");
const form = document.getElementById("mi-chat-form");
const input = document.getElementById("mi-chat-input");
const messages = document.getElementById("mi-chat-messages");

function openChat() {
  panel.classList.remove("mi-hidden");
  toggle.setAttribute("aria-expanded", "true");
  input.focus();
}
function closeChat() {
  panel.classList.add("mi-hidden");
  toggle.setAttribute("aria-expanded", "false");
}
toggle?.addEventListener("click", () => {
  if (panel.classList.contains("mi-hidden")) openChat();
  else closeChat();
});
closeBtn?.addEventListener("click", closeChat);

function addMsg(text, who) {
  const div = document.createElement("div");
  div.className = `mi-msg mi-${who}`;
  div.textContent = text;
  messages.appendChild(div);
  messages.scrollTop = messages.scrollHeight;
  return div;
}

form?.addEventListener("submit", async (e) => {
  e.preventDefault();
  const text = input.value.trim();
  if (!text) return;

  addMsg(text, "user");
  input.value = "";

  const placeholder = addMsg("…", "bot");

  try {
    const res = await fetch("/api/chat", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message: text })
    });

    if (!res.ok) throw new Error("bad response");
    const data = await res.json();
    placeholder.textContent = data.reply || "Sorry — I couldn’t generate a response.";
  } catch {
    placeholder.textContent = "Sorry — something went wrong. Please try again, or use /inquire.";
  }
});