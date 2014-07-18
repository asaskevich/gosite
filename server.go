package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/asaskevich/gosite/server"
	"github.com/martini-contrib/render"
	"strconv"
	"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"fmt"
	"os"
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
			result := server.WriteFeedback(r.FormValue("name"), r.FormValue("email"), r.FormValue("text"))
			if result == server.OK {
				http.Redirect(w, r, "/about?ok", http.StatusFound)
				return "OK"
			} else {
				// Render template with error
				http.Redirect(w, r, "/error/"+strconv.Itoa(result), http.StatusFound)
				return "SOME ERRORS"
			}
			return "OK"
		})
	// Rendering error page
	m.Get("/error/**", func(params martini.Params, r render.Render) {
			result := params["_1"]
			resstr, err := strconv.Atoi(result)
			error := map[string]interface{}{}
			error["id"] = resstr
			error["_err"] = err
			error["message"] = server.GetMessage(resstr)
			error["params"] = params
			error["header"] = r.Header()
			res, _ := json.MarshalIndent(error, "", "    ")
			r.HTML(200, "error", string(res))
		})
	// Rendering markdown page inside "/arch/" page
	// with custom ans simple template parser
	m.Get("/arch", func() string {
			wd, _ := os.Getwd()
			input, err := ioutil.ReadFile(wd + "/md/arch.md")
			if err != nil {
				fmt.Printf("Some errors with ReadFile: %v\n", err)
			}
			output := blackfriday.MarkdownCommon(input)
			return server.ParseTemplate("arch", map[string]string{"md": string(output)})
		})
	m.Run()
}
