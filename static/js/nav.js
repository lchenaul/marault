document.addEventListener("DOMContentLoaded", () => {
  const menuToggle = document.querySelector(".nav-toggle");
  const navMenu = document.querySelector("#navMenu");

  const servicesItem = document.querySelector(".nav-item.has-dropdown");
  const caretBtn = document.querySelector(".nav-caret");

  const isMobile = () => window.matchMedia("(max-width: 768px)").matches;

  // 1) Mobile: toggle the whole menu
  if (menuToggle && navMenu) {
    menuToggle.addEventListener("click", () => {
      const open = navMenu.classList.toggle("is-open");
      menuToggle.setAttribute("aria-expanded", open ? "true" : "false");

      // When closing menu, collapse Services dropdown too
      if (!open && servicesItem) servicesItem.classList.remove("open");
    });
  }

  // 2) Mobile: ONLY the caret expands/collapses Services
  // (The Services link is untouched so it navigates normally)
  if (caretBtn && servicesItem) {
    caretBtn.addEventListener("click", (e) => {
      if (!isMobile()) return;
      e.preventDefault();
      e.stopPropagation();
      servicesItem.classList.toggle("open");
    });
  }

  // 3) Close menu / dropdown when tapping outside (mobile only)
  document.addEventListener("click", (e) => {
    if (!isMobile()) return;
    if (!navMenu) return;

    const clickedInsideNav =
      navMenu.contains(e.target) || (menuToggle && menuToggle.contains(e.target));

    if (!clickedInsideNav) {
      navMenu.classList.remove("is-open");
      if (menuToggle) menuToggle.setAttribute("aria-expanded", "false");
      if (servicesItem) servicesItem.classList.remove("open");
    }
  });

  // 4) Close on ESC
  document.addEventListener("keydown", (e) => {
    if (e.key !== "Escape") return;

    if (navMenu) navMenu.classList.remove("is-open");
    if (menuToggle) menuToggle.setAttribute("aria-expanded", "false");
    if (servicesItem) servicesItem.classList.remove("open");
  });
});


