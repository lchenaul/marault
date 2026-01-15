// Services dropdown toggle (caret click + accessibility)
document.addEventListener("DOMContentLoaded", () => {
  const dropdown = document.querySelector(".has-dropdown");
  if (!dropdown) return;

  const caret = dropdown.querySelector(".nav-caret");
  if (!caret) return;

  const close = () => {
    dropdown.classList.remove("open");
    caret.setAttribute("aria-expanded", "false");
  };

  caret.addEventListener("click", (e) => {
    e.preventDefault();
    e.stopPropagation();

    const isOpen = dropdown.classList.toggle("open");
    caret.setAttribute("aria-expanded", String(isOpen));
  });

  // Close if clicking outside
  document.addEventListener("click", (e) => {
    if (!dropdown.contains(e.target)) close();
  });

  // Close on Escape
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") close();
  });
});


