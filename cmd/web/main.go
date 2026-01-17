package main

import (
	"html/template"
	"log"
	"net/http"
)

/* =========================
   TEMPLATE RENDER
========================= */

func render(w http.ResponseWriter, base string, page string, title string) {
	tmpl, err := template.ParseFiles(base, page)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("template parse error:", err)
		return
	}

	data := struct {
		Title string
	}{
		Title: title,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("template execute error:", err)
	}
}

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
			log.Println("template parse error:", err)
			return
		}

		data := struct {
			Title string
		}{
			Title: title,
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println("template execute error:", err)
		}
	}
}

/* =========================
   PAGE HANDLERS
========================= */

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Optional: ensure only "/" renders home
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/home.html",
		"Marault Intelligence",
	)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	// Optional: ensure only "/services" renders services overview
	if r.URL.Path != "/services" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/services.html",
		"Services — Marault Intelligence",
	)
}

func approachHandler(w http.ResponseWriter, r *http.Request) {
	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/approach.html",
		"The Marault Approach — Marault Intelligence",
	)
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/team.html",
		"Executive Team — Marault Intelligence",
	)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/contact.html",
		"Contact — Marault Intelligence",
	)
}

func inquireHandler(w http.ResponseWriter, r *http.Request) {
	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/inquire.html",
		"Inquire — Marault Intelligence",
	)
}

/* =========================
   MAIN
========================= */

func main() {
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Core pages
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/services", servicesHandler)

	// Nav pages (make sure these templates exist)
	mux.HandleFunc("/approach", approachHandler)
	mux.HandleFunc("/team", teamHandler)
	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/inquire", inquireHandler)

	// Service detail pages
	mux.HandleFunc(
		"/services/data-visibility-audit",
		servicePageHandler("data_visibility_audit.html", "Data Visibility Audit — Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/executive-dashboards-reporting",
		servicePageHandler("executive_dashboards_reporting.html", "Executive Dashboards & Reporting — Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/revenue-customer-analytics",
		servicePageHandler("revenue_customer_analytics.html", "Revenue & Customer Analytics — Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/forecasting-decision-modeling",
		servicePageHandler("forecasting_decision_modeling.html", "Forecasting & Decision Modeling — Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/private-client-analytics",
		servicePageHandler("private_client_analytics.html", "Private Client Analytics — Marault Intelligence"),
	)

	// Custom 404 wrapper
	server := http.Server{
		Addr: ":4000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, pattern := mux.Handler(r); pattern != "" {
				mux.ServeHTTP(w, r)
				return
			}
			http.NotFound(w, r)
		}),
	}

	log.Println("Starting server on :4000")
	log.Fatal(server.ListenAndServe())
}


