package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}*/

type Option struct {
	Value, Id, Text string
	Selected        string
}

const HTML = `
<!DOCTYPE html>
<html lang="en">
     <head>
        <meta charset="utf-8">
        <title>selected attribute</title>
    </head>
    <body>
        <form method="GET">
            <div>
                <label>Places:</label>
                <select id="places" name="places">
                   <option value=""></option>
                    {{range .}}
                    <option value="{{.Value}}" id="{{.Id}}" {{if eq .Selected .Value}}selected="selected"{{end}}>{{.Text}}</option>
                    {{end}}
                </select>
            </div>
            <input type="submit" value="submit">
        </form>
    </body>
</html>
`

var placesPageTmpl *template.Template = template.Must(template.New("PlacesPage").Parse(HTML))

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}

func name(w http.ResponseWriter, r *http.Request) {
	var selected string

	if r.FormValue("places") != "" {
		selected = r.FormValue("places")
	}

	options := []Option{
		{"Value1", "Id1", "Text1", selected},
		{"Value2", "Id2", "Text2", selected},
		{"Value3", "Id3", "Text3", selected},
	}

	if err := placesPageTmpl.Execute(w, options); err != nil {
		fmt.Println("Failed to build page", err)
	}
}
