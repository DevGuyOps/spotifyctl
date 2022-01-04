package spotifyctl

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

const redirectURI = "http://localhost:8080/callback"

var (
	server *http.Server
	auth   = spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistReadCollaborative,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserLibraryRead,
		spotify.ScopeUserLibraryModify,
	)
	state = ""
)

func SetupTokenServer() {
	http.HandleFunc("/callback", writeTokenToFile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	server = &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	log.Fatal(server.ListenAndServe())
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

	// Create config dir
	err = os.MkdirAll(getTokenDir(), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(getTokenPath(), file, 0644)
	if err != nil {
		fmt.Printf("Err writing token: %s\n", err)
		return
	}

	fmt.Printf("Token written to file: %s\n", getTokenPath())
	fmt.Println("Enjoy using spotifyctl!")

	// Close the server because we are done with it
	server.Close()
}

func NewClient() (*spotify.Client, error) {
	jsonFile, err := os.Open(getTokenPath())
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var token *oauth2.Token
	json.Unmarshal(byteValue, &token)

	client := auth.NewClient(token)
	return &client, nil
}

func getTokenPath() string {
	return getTokenDir() + "/token.json"
}

func getTokenDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return home + "/.config/spotifyctl"
}
