document.addEventListener("DOMContentLoaded", () => {
  const toggle = document.querySelector(".mobile-nav-toggle");
  const menu = document.getElementById("mobileNavMenu");
  const dropdownItems = document.querySelectorAll(".mobile-nav-item.has-mobile-dropdown");

  if (!toggle || !menu) return;

  const closeAllDropdowns = () => {
    dropdownItems.forEach((item) => {
      item.classList.remove("open");
      const caret = item.querySelector(".mobile-nav-caret");
      if (caret) caret.setAttribute("aria-expanded", "false");
    });
  };

  const closeMenu = () => {
    menu.classList.remove("is-open");
    toggle.classList.remove("is-open");
    toggle.setAttribute("aria-expanded", "false");
    document.body.classList.remove("mobile-menu-open");
    closeAllDropdowns();
  };

  toggle.addEventListener("click", (e) => {
    e.preventDefault();
    e.stopPropagation();

    const isOpen = menu.classList.toggle("is-open");
    toggle.classList.toggle("is-open", isOpen);
    toggle.setAttribute("aria-expanded", isOpen ? "true" : "false");
    document.body.classList.toggle("mobile-menu-open", isOpen);

    if (!isOpen) closeAllDropdowns();
  });

  dropdownItems.forEach((item) => {
    const caret = item.querySelector(".mobile-nav-caret");
    if (!caret) return;

    caret.addEventListener("click", (e) => {
      e.preventDefault();
      e.stopPropagation();

      const willOpen = !item.classList.contains("open");

      dropdownItems.forEach((other) => {
        if (other !== item) {
          other.classList.remove("open");
          const otherCaret = other.querySelector(".mobile-nav-caret");
          if (otherCaret) otherCaret.setAttribute("aria-expanded", "false");
        }
      });

      item.classList.toggle("open", willOpen);
      caret.setAttribute("aria-expanded", willOpen ? "true" : "false");
    });
  });

  document.addEventListener("click", (e) => {
    if (!menu.classList.contains("is-open")) return;

    const nav = document.querySelector(".mobile-nav");
    if (nav && !nav.contains(e.target)) {
      closeMenu();
    }
  });

  window.addEventListener("resize", () => {
    if (window.innerWidth > 768) {
      closeMenu();
    }
  });
});