package spotifyctl

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

func Authorise() {
	SetupTokenServer()
}

func Play() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	err = client.Play()
	if err != nil {
		return err
	}

	return nil
}

func Pause() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	err = client.Pause()
	if err != nil {
		return err
	}

	return nil
}

func Next() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	err = client.Next()
	if err != nil {
		return err
	}

	return nil
}

func Previous() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	err = client.Previous()
	if err != nil {
		return err
	}

	return nil
}

func UpdateLikeStatusForCurrentTrack() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	playing, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	if playing.Item == nil {
		return fmt.Errorf("No track is playing")
	}

	trackDetails := fmt.Sprintf(
		"%s [%s]",
		playing.Item.Name,
		formatArtistNames(playing.Item.Artists),
	)

	// Determine if the track is already in the library
	hasTracks, err := client.UserHasTracks(playing.Item.ID)
	if err != nil {
		return err
	}

	if !hasTracks[0] {
		// Save current track
		err = client.AddTracksToLibrary(playing.Item.ID)
		if err != nil {
			return err
		}

		fmt.Printf("Liked -> %s\n", trackDetails)
	} else {
		// Remove current track
		err = client.RemoveTracksFromLibrary(playing.Item.ID)
		if err != nil {
			return err
		}

		fmt.Printf("Unliked -> %s\n", trackDetails)
	}

	return nil
}

func PlaylistsList() error {
	limit := 50
	limitRef := &limit

	client, err := NewClient()
	if err != nil {
		return err
	}

	user, err := client.CurrentUser()
	if err != nil {
		return err
	}

	playlists, _ := client.GetPlaylistsForUserOpt(user.ID, &spotify.Options{
		Limit: limitRef,
	})
	for _, playlist := range playlists.Playlists {
		fmt.Printf("[%s] %s\n", playlist.ID, playlist.Name)
	}

	return nil
}

func PlaylistPlay(playlistID string) error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	user, err := client.CurrentUser()
	if err != nil {
		return err
	}

	playlists, _ := client.GetPlaylistsForUser(user.ID)
	for _, playlist := range playlists.Playlists {
		if playlist.ID.String() == playlistID {
			err := client.PlayOpt(&spotify.PlayOptions{
				PlaybackContext: &playlist.URI,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Now playing: %s\n", playlist.Name)
			break
		}
	}

	return fmt.Errorf("Playlist does not exist: %s", playlistID)
}

func DeviceList() ([]string, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	devices, err := client.PlayerDevices()
	if err != nil {
		log.Fatal(err)
	}

	deviceList := []string{}
	for _, device := range devices {
		deviceList = append(deviceList, device.Name)
	}

	return deviceList, nil
}

func DeviceSetContext() {

}

func DeviceCurrentContext() {

}
