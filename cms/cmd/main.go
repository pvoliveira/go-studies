package main

import (
	"os"
	"pvoliveira/learning-path-master-google-go/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
