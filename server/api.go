package server

import (
	"encoding/json"
	"github.com/robfig/config"
	"io/ioutil"
	"os"
)

// Return list of posts in JSON
func GetPosts() string {
	result := map[string]map[string]string{}
	wd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(wd + "/md/posts/")
	for _, f := range files {
		if f.IsDir() {
			id := f.Name()
			cfg := wd + "/md/posts/" + id + "/config.cfg"
			cf, _ := config.ReadDefault(cfg)
			name, _ := cf.String("TOPIC", "name")
			date, _ := cf.String("TOPIC", "date")
			author, _ := cf.String("TOPIC", "author")
			original, _ := cf.String("TOPIC", "original")
			result[id] = map[string]string{"id": id, "name": name, "date": date,
				"author": author, "original": original}
		}
	}
	jsonStr, _ := json.Marshal(result)
	return string(jsonStr)
}
