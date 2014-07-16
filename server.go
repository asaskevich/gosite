package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/asaskevich/gosite/server"
	"github.com/martini-contrib/render"
	"strconv"
)

// Entry point for server
func main() {
	m := martini.Classic()
	// set up the session, and map it into martini
	store := sessions.NewCookieStore([]byte("auth"))
	m.Use(sessions.Sessions("asaskevich", store))
	// render html templates from templates directory
	m.Use(render.Renderer())
	// Redirect from "/" to "/home" url
	m.Get("/", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			http.Redirect(w, r, "/home", http.StatusFound)
			return "OK"
		})
	// Register feedback posting
	m.Post("/feedback", func(w http.ResponseWriter, r *http.Request, session sessions.Session) string {
			result := server.WriteFeedback(r.FormValue("email"), r.FormValue("name"), r.FormValue("text"))
			if result == server.OK {
				http.Redirect(w, r, "/about", http.StatusFound)
				return "OK"
			} else {
				// TODO
				// Render template with error
				http.Redirect(w, r, "/error/"+strconv.Itoa(result), http.StatusFound)
				return "SOME ERRORS"
			}
			return "OK"
		})
	// Rendering error page
	m.Get("/error/**", func(params martini.Params, r render.Render) {
			result := params["_1"]
			r.HTML(200, "error", result)
		})
	m.Run()
}
