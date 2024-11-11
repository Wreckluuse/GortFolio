package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Advice struct {
	Id     string
	Advice string
	Source string
}

type StoicQuote struct {
	Data struct {
		Author string
		Quote  string
	}
}

type UselessFact struct {
	Id        string
	Text      string
	Source    string
	SourceUrl string
	Language  string
	Permalink string
}

type Affirmation struct {
	Affirmation string
}

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// Serve the main page
	// http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/getAdvice", handleAdviceRequest)
	http.HandleFunc("/getAffirmation", handleAffirmationRequest)
	http.HandleFunc("/getStoic", handleStoicRequest)
	http.HandleFunc("/getUselessFact", handleUselessFactRequest)
	http.HandleFunc("/getQuote", handleQuoteRequest)
	// Start server
	http.ListenAndServe(":8080", nil)
}

// func serveTemplate(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))
// 	tmpl.Execute(w, nil)
// }

func handleQuoteRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	w.Write([]byte("Test"))
}

func handleStoicRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getQuote("st")))
}

func handleUselessFactRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getQuote("uf")))
}

func handleAdviceRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getQuote("ad")))
}

func handleAffirmationRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getQuote("af")))
}

func getQuote(flag string) string {
	var rv string
	switch flag {
	case "st":
		resp, err := http.Get("https://stoic.tekloon.net/stoic-quote")

		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var stoicQuote StoicQuote
		json.Unmarshal([]byte(body), &stoicQuote)
		rv = stoicQuote.Data.Quote
	case "uf":
		resp, err := http.Get("https://uselessfacts.jsph.pl/api/v2/facts/random")

		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var uselessfact UselessFact
		json.Unmarshal([]byte(body), &uselessfact)

		rv = uselessfact.Text
	case "ad":
		resp, err := http.Get("https://kk-advice.koyeb.app/api/advice")

		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var advice Advice
		json.Unmarshal([]byte(body), &advice)

		rv = advice.Advice
	case "af":
		resp, err := http.Get("https://www.affirmations.dev/")

		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var affirmation Affirmation
		json.Unmarshal([]byte(body), &affirmation)

		rv = affirmation.Affirmation
	}
	return rv
}
