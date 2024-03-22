package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const port = ":5555"

var templates = template.Must(template.ParseFiles("templates/code.html"))

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Status 404: Page Not Found")
		return
	}

	templates.ExecuteTemplate(w, "code.html", "")
}

func dynamique(w http.ResponseWriter, r *http.Request) {
	var fichier *os.File

	var s []string
	var str string
	recept := r.FormValue("fichier")

	receptxt := r.FormValue("formulaire")
	for i := range receptxt {
		if (receptxt[i] < 32 || receptxt[i] > 127) && (receptxt[i] != '\n' && receptxt[i] != '\r') {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "veuillez entrer un caractere appartenant a la table ascii")
			return
			// faire une exeption pour le /n
		}
	}
	content := strings.ReplaceAll(receptxt, "\r\n", "\\n")

	contente := strings.Split(content, "\\n")
	if strings.Contains(receptxt, "ù") {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return

	}

	if recept == "shadow" {
		fichier, _ = os.Open("shadow.txt")
	} else if recept == "standard" {
		fichier, _ = os.Open("standard.txt")
	} else if recept == "thinkertoy" {
		fichier, _ = os.Open("thinkertoy.txt")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Status 400: Bad Request")
		return
	}

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	for _, element := range contente {
		if len(element) > 0 {
			line := []rune(element)

			for a := 0; a < 8; a++ {

				for i := 0; i < len(line); i++ {
					group := (int(line[i]) - 32) * 9
					adress := group + a + 1

					str += (s[adress])
				}
				str += (string(rune('\n')))
			}
		} else {
			str += (string(rune('\n')))
		}
	}

	templates.ExecuteTemplate(w, "code.html", "")
	fmt.Fprint(w, str)
	strfile := []byte(str)

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 1024)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(strfile) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}

func export(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename= result-ascii")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	http.ServeFile(w, r, "./output.txt")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii", dynamique)
	// http.HandleFunc("/export", export)
	http.HandleFunc("/Export", export)

	fmt.Println("http://localhost:5555")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css/"))))
	http.ListenAndServe(port, nil)
}
