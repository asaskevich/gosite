package server

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// Parse given template into string
func ParseTemplate(name string, values map[string]string) string {
	wd, _ := os.Getwd()
	input, err := ioutil.ReadFile(wd + "/templates/" + name + ".temp")
	if err != nil {
		fmt.Printf("Some errors with ReadFile: %v\n", err)
		return ""
	}
	code := string(input)
	for key, value := range values {
		code = strings.Replace(code, "{{"+key+"}}", value, -1)
	}
	return code
}
