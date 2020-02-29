package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func getToken() {
	// first start an HTTP server
	http.HandleFunc("/callback", writeTokenToFile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	http.ListenAndServe(":8080", nil)
}

func writeTokenToFile(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// Write token to file
	file, _ := json.MarshalIndent(tok, "", " ")
	_ = ioutil.WriteFile("token.json", file, 0644)

	fmt.Println("Token written to file: token.json")
}

func newClient() spotify.Client {
	jsonFile, err := os.Open("token.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var token *oauth2.Token
	json.Unmarshal(byteValue, &token)

	return auth.NewClient(token)
}
