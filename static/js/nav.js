// static/js/nav.js
document.addEventListener("DOMContentLoaded", () => {
  const dropdownItem = document.querySelector(".nav-item.has-dropdown");
  const caret = document.querySelector(".nav-caret");

  if (!dropdownItem || !caret) return;

  caret.addEventListener("click", (e) => {
    e.preventDefault();
    e.stopPropagation();

    const isOpen = dropdownItem.classList.toggle("is-open");
    caret.setAttribute("aria-expanded", String(isOpen));
  });

  // close when clicking outside
  document.addEventListener("click", (e) => {
    if (!dropdownItem.contains(e.target)) {
      dropdownItem.classList.remove("is-open");
      caret.setAttribute("aria-expanded", "false");
    }
  });

  // close on ESC
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") {
      dropdownItem.classList.remove("is-open");
      caret.setAttribute("aria-expanded", "false");
    }
  });
});

