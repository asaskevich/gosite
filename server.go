package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"os"
	"labix.org/v2/mgo"
	"fmt"
)

type Feedback struct {
	Name  string
	Email string
	Text  string
}

// Establish connection to MongoDB and write feedback to collection
func WriteToDB(name string, email string, text string) {
	url := os.Getenv("DATABASE")
	fmt.Printf("DB_URL is %v\n", url)
	sess, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		return
	}
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)

	c := sess.DB("asaskevich").C("feedback")
	err = c.Insert(&Feedback{name, email, text})
	if err != nil {
		fmt.Printf("Can't write to mongo, go error %v\n", err)
		return
	}

}

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
			WriteToDB(r.FormValue("email"), r.FormValue("name"), r.FormValue("text"))
			http.Redirect(w, r, "/about", http.StatusFound)
			return "OK"
		})
	m.Run()
}
