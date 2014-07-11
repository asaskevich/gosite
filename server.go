package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
			return "Hello World"
		})
	m.Get("/test", func() string {
			return "Test"
		})
	m.Run()
}
