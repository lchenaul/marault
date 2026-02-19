package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
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
	// Prevent "/" from catching all routes
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

func approachHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/approach" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/approach.html",
		"The Marault Approach | Marault Intelligence",
	)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/services" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/services.html",
		"Services | Marault Intelligence",
	)
}

// Keep YOUR team page (you said don't drop your pages)
func teamHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/team" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/team.html",
		"Executive Team | Marault Intelligence",
	)
}

// Also keep her executive-team route if you want that nav label/path too
func executiveTeamHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/executive-team" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/executive.html",
		"Executive Team | Marault Intelligence",
	)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" {
		http.NotFound(w, r)
		return
	}

	render(
		w,
		"./internal/templates/base.html",
		"./internal/templates/contact.html",
		"Contact | Marault Intelligence",
	)
}

/* =========================
   INQUIRE (GET + POST)
========================= */

func inquireHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inquire" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		render(
			w,
			"./internal/templates/base.html",
			"./internal/templates/inquire.html",
			"Inquire | Marault Intelligence",
		)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		company := r.FormValue("company")
		message := r.FormValue("message")
		services := r.Form["services"]

		selectedServices := strings.Join(services, ", ")

		if err := sendInquiryEmail(name, email, company, selectedServices, message); err != nil {
			log.Println(err)
			http.Error(w, "Unable to send message", http.StatusInternalServerError)
			return
		}

		// Thank you page (you'll add it if missing)
		render(
			w,
			"./internal/templates/base.html",
			"./internal/templates/thankyou.html",
			"Thank You | Marault Intelligence",
		)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

/* =========================
   EMAIL (ENV VARS ONLY)
========================= */

func sendInquiryEmail(name, email, company, services, message string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	toList := os.Getenv("SMTP_TO") // comma-separated, e.g. "a@x.com,b@y.com"

	if from == "" || password == "" || toList == "" {
		return fmt.Errorf("missing SMTP env vars: SMTP_FROM, SMTP_PASSWORD, SMTP_TO")
	}

	to := strings.Split(toList, ",")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body := fmt.Sprintf(
		"New Inquiry\n\nName: %s\nEmail: %s\nCompany: %s\nServices: %s\n\nMessage:\n%s",
		name, email, company, services, message,
	)

	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: New Website Inquiry\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
}

/* =========================
   MAIN
========================= */

func main() {
	mux := http.NewServeMux()

	/* =========================
	   Static files
	========================= */
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	/* =========================
	   Core pages (keep yours + her route)
	========================= */
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/approach", approachHandler)
	mux.HandleFunc("/services", servicesHandler)

	// Keep your team page AND keep her executive-team route (no loss of content)
	mux.HandleFunc("/team", teamHandler)
	mux.HandleFunc("/executive-team", executiveTeamHandler)

	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/inquire", inquireHandler)

	/* =========================
	   Services detail pages (keep all yours)
	========================= */
	mux.HandleFunc(
		"/services/data-visibility-audit",
		servicePageHandler("data_visibility_audit.html", "Data Visibility Audit | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/executive-dashboards-reporting",
		servicePageHandler("executive_dashboards_reporting.html", "Executive Dashboards & Reporting | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/revenue-customer-analytics",
		servicePageHandler("revenue_customer_analytics.html", "Revenue & Customer Analytics | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/forecasting-decision-modeling",
		servicePageHandler("forecasting_decision_modeling.html", "Forecasting & Decision Modeling | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/private-client-analytics",
		servicePageHandler("private_client_analytics.html", "Private Client Analytics | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/custom-website-build",
		servicePageHandler("custom-website-build.html", "Custom Website Build | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/template-based-build",
		servicePageHandler("template-based-build.html", "Template-Based Build | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/website-redesign",
		servicePageHandler("website-redesign.html", "Website Redesign | Marault Intelligence"),
	)
	mux.HandleFunc(
		"/services/ux-ui-design",
		servicePageHandler("ux-ui-design.html", "UX/UI Design | Marault Intelligence"),
	)

	/* =========================
	   Custom 404 wrapper
	========================= */
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


