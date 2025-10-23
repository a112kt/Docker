package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Serve static files (for images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/z3ter", z3terHandler)
	http.HandleFunc("/z3ter-birthday", z3terBirthdayHandler)

	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  - http://localhost:8080/z3ter")
	fmt.Println("  - http://localhost:8080/z3ter-birthday")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Docker Demo App</h1>")
	fmt.Fprintf(w, "<p><a href='/z3ter'>View Z3ter Image</a></p>")
	fmt.Fprintf(w, "<p><a href='/z3ter-birthday'>View Z3ter Birthday</a></p>")
}

func z3terHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Z3ter Image</h1>")
	fmt.Fprintf(w, "<img src='/static/Z3ter.jpg' width='800'>")
}

func z3terBirthdayHandler(w http.ResponseWriter, r *http.Request) {

	birthday := time.Date(0, time.May, 21, 0, 0, 0, 0, time.UTC)
	today := time.Now()

	fmt.Fprintf(w, "<h1>🎂 Z3ter Birthday 🎉</h1>")
	fmt.Fprintf(w, "<h2>Birthday Date: May 21</h2>")
	fmt.Fprintf(w, "<h3>Today’s Date: %s</h3>", today.Format("Monday, January 2, 2006"))

	if today.Month() == birthday.Month() && today.Day() == birthday.Day() {
		fmt.Fprintf(w, "<p style='color:green; font-size:20px;'>🎉 Happy Birthday, Z3ter! 🎉</p>")
	} else {
		fmt.Fprintf(w, "<p style='color:blue; font-size:18px;'>Not Z3ter’s birthday today 😄</p>")
	}
}
