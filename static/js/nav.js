document.addEventListener("DOMContentLoaded", () => {
  const menuToggle = document.querySelector(".nav-toggle");
  const navMenu = document.querySelector("#navMenu");

  const dropdownItems = document.querySelectorAll(".nav-item.has-dropdown");

  const isMobile = () => window.matchMedia("(max-width: 768px)").matches;

  const closeAllDesktopDropdowns = () => {
    document.querySelectorAll(".dropdown.open").forEach(d => d.classList.remove("open"));
    document.querySelectorAll(".nav-caret[aria-expanded='true']").forEach(c =>
      c.setAttribute("aria-expanded", "false")
    );
  };

  const closeMobileMenu = () => {
    if (navMenu) navMenu.classList.remove("is-open");
    if (menuToggle) menuToggle.setAttribute("aria-expanded", "false");
    dropdownItems.forEach(item => item.classList.remove("open")); // accordion state
  };

  // 1) Mobile: toggle the whole menu
  if (menuToggle && navMenu) {
    menuToggle.addEventListener("click", (e) => {
      e.stopPropagation();
      const open = navMenu.classList.toggle("is-open");
      menuToggle.setAttribute("aria-expanded", open ? "true" : "false");

      // when closing menu, collapse any open service accordions
      if (!open) dropdownItems.forEach(item => item.classList.remove("open"));
    });
  }

  // 2) Caret behavior: desktop toggles .dropdown.open, mobile toggles .nav-item.open
  dropdownItems.forEach(item => {
    const caret = item.querySelector(".nav-caret");
    const dropdown = item.querySelector(".dropdown");
    if (!caret || !dropdown) return;

    caret.addEventListener("click", (e) => {
      e.preventDefault();
      e.stopPropagation();

      if (isMobile()) {
        // Mobile accordion behavior (your CSS expects .nav-item.open)
        item.classList.toggle("open");
        return;
      }

      // Desktop behavior (her CSS/JS expects .dropdown.open)
      const isOpen = dropdown.classList.contains("open");

      closeAllDesktopDropdowns();
      if (!isOpen) {
        dropdown.classList.add("open");
        caret.setAttribute("aria-expanded", "true");
      }
    });
  });

  // 3) Click outside: close desktop dropdowns; on mobile close menu entirely
  document.addEventListener("click", (e) => {
    if (isMobile()) {
      if (!navMenu) return;
      const clickedInsideNav =
        navMenu.contains(e.target) || (menuToggle && menuToggle.contains(e.target));
      if (!clickedInsideNav) closeMobileMenu();
      return;
    }

    // Desktop: close dropdowns when clicking outside
    closeAllDesktopDropdowns();
  });

  // 4) ESC closes everything (both modes)
  document.addEventListener("keydown", (e) => {
    if (e.key !== "Escape") return;
    closeAllDesktopDropdowns();
    closeMobileMenu();
  });

  // 5) On resize across breakpoint: clean up states so you don't get stuck open
  window.addEventListener("resize", () => {
    closeAllDesktopDropdowns();
    dropdownItems.forEach(item => item.classList.remove("open"));
    if (!isMobile() && navMenu) {
      navMenu.classList.remove("is-open");
      if (menuToggle) menuToggle.setAttribute("aria-expanded", "false");
    }
  }, { passive: true });
});



