// This example demonstrates how to authenticate with Spotify using the authorization code flow.
// In order to run this example yourself, you'll need to:
//
//  1. Register an application at: https://developer.spotify.com/my-applications/
//       - Use "http://localhost:8080/callback" as the redirect URI
//  2. Set the SPOTIFY_ID environment variable to the client ID you got in step 1.
//  3. Set the SPOTIFY_SECRET environment variable to the client secret from step 1.
package main

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopeUserModifyPlaybackState)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	// getToken()
	client := newClient()

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	// Get Playlists
	playlists, _ := client.GetPlaylistsForUser(user.ID)
	// for _, playlist := range playlists.Playlists {
	// 	fmt.Printf("%s -> %s\n", playlist.Name, playlist.URI)
	// }

	// Play playlist
	playlistName := "Hardstyle"
	for _, playlist := range playlists.Playlists {
		if playlist.Name == playlistName {
			fmt.Println(playlist.Name)
			err := client.PlayOpt(&spotify.PlayOptions{
				PlaybackContext: &playlist.URI,
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
