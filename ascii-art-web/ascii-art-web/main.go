package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	file2, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0o600)
	defer file2.Close() // on ferme automatiquement à la fin de notre programme
	if r.URL.Path != "/" {
		http.Error(w, "Error 404: file not found.", http.StatusNotFound)
		return
	}
	_, err := template.ParseFiles("main.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}

	// erreur 400 403
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "main.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("biblio")
		text := r.FormValue("textarea")
		text = strings.ReplaceAll(text, "\r", "\n")
		text2 := strings.Split(text, "\n")
		for count := 0; count < len(text2); count++ {
			for _, i := range text2[count] {
				if i < 32 || i > 126 {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Error 400: Bad Request"))
					return
				}
			}
		}
		// fmt.Println(text)
		ascii(text, name)
		data2, err := ioutil.ReadFile("test.txt") // lire le fichier
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Print(string(data2))
		fmt.Fprintf(w, "%s\n", string(data2))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	fmt.Println("Server launch at : http://localhost:8070/ ")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8070", nil)
}
