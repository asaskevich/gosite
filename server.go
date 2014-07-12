package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

// Entry point for server
func main() {
	m := martini.Classic()
	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("auth"))
	m.Use(sessions.Sessions("go-webapp-part2", store))
	// Redirect from "/" to "/home" url
	m.Get("/", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})
	m.Run()
}
