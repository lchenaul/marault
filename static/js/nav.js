document.addEventListener("DOMContentLoaded", () => {
  const nav = document.querySelector(".nav");
  const menuToggle = document.querySelector(".nav-toggle");
  const navMenu = document.querySelector("#navMenu");
  const dropdownItems = document.querySelectorAll(".nav-item.has-dropdown");

  const isMobile = () => window.matchMedia("(max-width: 768px)").matches;

  const closeDropdowns = () => {
    dropdownItems.forEach((item) => {
      item.classList.remove("open");
      const caret = item.querySelector(".nav-caret");
      if (caret) caret.setAttribute("aria-expanded", "false");
    });
  };

  const closeMenu = () => {
    if (navMenu) navMenu.classList.remove("is-open");
    if (menuToggle) menuToggle.setAttribute("aria-expanded", "false");
    closeDropdowns();
  };

  if (menuToggle && navMenu) {
    menuToggle.addEventListener("click", (e) => {
      if (!isMobile()) return;
      e.stopPropagation();
      const open = navMenu.classList.toggle("is-open");
      menuToggle.setAttribute("aria-expanded", open ? "true" : "false");
    });
  }

  dropdownItems.forEach((item) => {
    const caret = item.querySelector(".nav-caret");
    if (!caret) return;

    caret.addEventListener("click", (e) => {
      if (!isMobile()) return;

      e.preventDefault();
      e.stopPropagation();

      const willOpen = !item.classList.contains("open");

      dropdownItems.forEach((other) => {
        if (other !== item) {
          other.classList.remove("open");
          const otherCaret = other.querySelector(".nav-caret");
          if (otherCaret) otherCaret.setAttribute("aria-expanded", "false");
        }
      });

      item.classList.toggle("open", willOpen);
      caret.setAttribute("aria-expanded", willOpen ? "true" : "false");
    });
  });

  document.addEventListener("click", (e) => {
    if (!isMobile()) return;
    if (nav && !nav.contains(e.target)) closeMenu();
  });

  window.addEventListener("resize", () => {
    if (!isMobile()) closeMenu();
  });
});
