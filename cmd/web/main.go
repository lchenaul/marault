package main

import (
	"html/template"
	"log"
	"net/http"
)

/* =========================
   GENERIC SERVICE PAGE HANDLER
========================= */
func servicePageHandler(templateFile string, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"./internal/templates/base.html",
			"./internal/templates/"+templateFile,
		)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		data := struct {
			Title string
		}{
			Title: title,
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, JS, images, videos)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	/* =========================
	   CORE PAGES
	========================= */
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/services", servicesHandler)

	/* =========================
	   BUSINESS INTELLIGENCE
	========================= */
	mux.HandleFunc(
		"/services/data-visibility-audit",
		servicePageHandler(
			"data_visibility_audit.html",
			"Data Visibility Audit — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/executive-dashboards-reporting",
		servicePageHandler(
			"executive_dashboards_reporting.html",
			"Executive Dashboards & Reporting — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/revenue-customer-analytics",
		servicePageHandler(
			"revenue_customer_analytics.html",
			"Revenue & Customer Analytics — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/forecasting-decision-modeling",
		servicePageHandler(
			"forecasting_decision_modeling.html",
			"Forecasting & Decision Modeling — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/private-client-analytics",
		servicePageHandler(
			"private_client_analytics.html",
			"Private Client Analytics — Marault Intelligence",
		),
	)

	/* =========================
	   WEB SERVICES
	========================= */
	mux.HandleFunc(
		"/services/custom-website-build",
		servicePageHandler(
			"custom_website_build.html",
			"Custom Website Build — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/template-based-build",
		servicePageHandler(
			"template_based_build.html",
			"Template-Based Build — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/website-redesign",
		servicePageHandler(
			"website_redesign.html",
			"Website Redesign — Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/ux-ui-design",
		servicePageHandler(
			"ux_ui_design.html",
			"UX / UI Design — Marault Intelligence",
		),
	)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/* =========================
   HOME
========================= */
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./internal/templates/base.html",
		"./internal/templates/home.html",
	)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Marault Intelligence",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

/* =========================
   SERVICES OVERVIEW
========================= */
func servicesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./internal/templates/base.html",
		"./internal/templates/services.html",
	)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Services — Marault Intelligence",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}


