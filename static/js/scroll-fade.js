document.addEventListener("DOMContentLoaded", () => {
  let targets;

  try {
    targets = document.querySelectorAll(
      `
      .services-white-band,
      .services-details-inner > section,
      .services-audit-card,
      .services-revenue-card,
      .services-forecasting-card,
      .services-private-card,
      .services-web-inner,
      .services-template-inner,
      .services-redesign-inner,
      .services-uxui-card,
      .svc-features,
      .svc-exec-panel,
      .svc-checklist,
      .services-grid,
      .svc-pill,
      .svc-check-row
      `
    );
  } catch (e) {
    console.error("scroll-fade selector error:", e);
    return;
  }

  if (!targets.length) return;

  if (!("IntersectionObserver" in window)) {
    targets.forEach(el => el.classList.add("is-visible"));
    return;
  }

  const observer = new IntersectionObserver(
    (entries, obs) => {
      entries.forEach(entry => {
        if (!entry.isIntersecting) return;
        entry.target.classList.add("is-visible");
        obs.unobserve(entry.target);
      });
    },
    {
      root: null,
      threshold: 0.18,
      rootMargin: "0px 0px -10% 0px"
    }
  );

  targets.forEach(el => observer.observe(el));
});

