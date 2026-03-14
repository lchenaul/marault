package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func isMobile(r *http.Request) bool {
	ua := strings.ToLower(r.UserAgent())

	mobileSignals := []string{
		"iphone",
		"android",
		"mobile",
		"ipad",
		"ipod",
	}

	for _, s := range mobileSignals {
		if strings.Contains(ua, s) {
			return true
		}
	}

	return false
}

func getBaseTemplate(r *http.Request) string {
	if isMobile(r) {
		return "./internal/templates/base-mobile.html"
	}
	return "./internal/templates/base.html"
}

/* =========================
   GENERIC SERVICE PAGE HANDLER
========================= */
func servicePageHandler(templateFile string, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			getBaseTemplate(r),
			"./internal/templates/"+templateFile,
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Page  string
		}{
			Title: title,
			Page:  "services",
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	mux := http.NewServeMux()

	/* =========================
	   Static files
	========================= */
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	/* =========================
	   Core pages
	========================= */
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/approach", approachHandler)
	mux.HandleFunc("/executive-team", executiveTeamHandler)
	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/inquire", inquireHandler)
	mux.HandleFunc("/philosophy", philosophyHandler)

	/* =========================
	   Services
	========================= */
	mux.HandleFunc("/services", servicesHandler)

	mux.HandleFunc(
		"/services/data-visibility-audit",
		servicePageHandler(
			"data-visibility-audit.html",
			"Data Visibility Audit | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/revenue-customer-analytics",
		servicePageHandler(
			"revenue.html",
			"Revenue & Customer Analytics | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/custom-website-build",
		servicePageHandler(
			"custom-website-build.html",
			"Custom Website Build | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/executive-dashboards-reporting",
		servicePageHandler(
			"executive-dashboards-reporting.html",
			"Executive Dashboards & Reporting | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/forecasting-decision-modeling",
		servicePageHandler(
			"forecasting-decision-modeling.html",
			"Forecasting & Decision Modeling | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/private-client-analytics",
		servicePageHandler(
			"private-client-analytics.html",
			"Private Client Analytics | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/template-based-build",
		servicePageHandler(
			"template-based-build.html",
			"Template-Based Build | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/website-redesign",
		servicePageHandler(
			"website-redesign.html",
			"Website Redesign | Marault Intelligence",
		),
	)

	mux.HandleFunc(
		"/services/ux-ui-design",
		servicePageHandler(
			"ux-ui-design.html",
			"UX/UI Design | Marault Intelligence",
		),
	)

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

/* =========================
   HOME PAGE
========================= */
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(
		getBaseTemplate(r),
		"./internal/templates/home.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "Marault Intelligence",
		Page:  "home",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

/* =========================
   APPROACH PAGE
========================= */
func approachHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		getBaseTemplate(r),
		"./internal/templates/approach.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "The Marault Approach | Marault Intelligence",
		Page:  "approach",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

/* =========================
   EXECUTIVE TEAM
========================= */
func executiveTeamHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		getBaseTemplate(r),
		"./internal/templates/executive.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "Executive Team | Marault Intelligence",
		Page:  "executive",
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
		getBaseTemplate(r),
		"./internal/templates/services.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "Services | Marault Intelligence",
		Page:  "services",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

/* =========================
   CONTACT
========================= */
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		getBaseTemplate(r),
		"./internal/templates/contact.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "Contact | Marault Intelligence",
		Page:  "contact",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

/* =========================
   INQUIRE
========================= */
func inquireHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles(
			getBaseTemplate(r),
			"./internal/templates/inquire.html",
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Page  string
		}{
			Title: "Inquire | Marault Intelligence",
			Page:  "inquire",
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		company := r.FormValue("company")
		message := r.FormValue("message")
		services := r.Form["services"]

		selectedServices := strings.Join(services, ", ")

		err := sendInquiryEmail(name, email, company, selectedServices, message)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unable to send message", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles(
			getBaseTemplate(r),
			"./internal/templates/thankyou.html",
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Page  string
		}{
			Title: "Thank You | Marault Intelligence",
			Page:  "thankyou",
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	}
}

/* =========================
   EMAIL SENDER
========================= */
func sendInquiryEmail(name, email, company, services, message string) error {
	from := "caroline@maraultintelligence.com"
	password := "fxhiauwzwnrqhrhk"

	to := []string{
		"caroline@maraultintelligence.com",
		"lindsey@maraultintelligence.com",
	}

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
func philosophyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		getBaseTemplate(r),
		"./internal/templates/philosophy.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Page  string
	}{
		Title: "Data Philosophy | Marault Intelligence",
		Page:  "philosophy",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}









