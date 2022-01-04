package spotifyctl

import "github.com/zmb3/spotify"

// Format the artist names nicely
func formatArtistNames(artists []spotify.SimpleArtist) string {
	artistList := ""
	for i := 0; i < len(artists); i++ {
		artistList += artists[i].Name

		if i < len(artists)-1 {
			artistList += ", "
		}
	}

	return artistList
}
