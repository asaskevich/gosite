package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
)

// Entry point for server
func main() {
	m := martini.Classic()
	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("auth"))
	m.Use(sessions.Sessions("asaskevich", store))
	// Redirect from "/" to "/home" url
	m.Get("/", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
		http.Redirect(w, r, "/home", http.StatusFound)
		return "OK"
	})
	// Register feedback posting
	m.Post("/feedback", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
		return r.FormValue("name") + " " + r.FormValue("email")
	})
	m.Run()
}
