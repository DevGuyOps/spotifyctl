package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GuySWatson/spotifyctl/pkg/spotifyctl"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "auth",
				Usage: "Authorise the client",
				Action: func(c *cli.Context) error {
					spotifyctl.Authorise()
					return nil
				},
			},
			{
				Name:  "play",
				Usage: "Play track",
				Action: func(c *cli.Context) error {
					return spotifyctl.Play()
				},
			},
			{
				Name:  "pause",
				Usage: "Pause track",
				Action: func(c *cli.Context) error {
					return spotifyctl.Pause()
				},
			},

			{
				Name:  "next",
				Usage: "Next track",
				Action: func(c *cli.Context) error {
					return spotifyctl.Next()
				},
			},
			{
				Name:  "previous",
				Usage: "Previous track",
				Action: func(c *cli.Context) error {
					return spotifyctl.Previous()
				},
			},
			{
				Name:  "playlist",
				Usage: "Playlist",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List user playlists",
						Action: func(c *cli.Context) error {
							return spotifyctl.PlaylistsList()
						},
					},
					{
						Name:  "play",
						Usage: "Play a playlist",
						Action: func(c *cli.Context) error {
							return spotifyctl.PlaylistPlay(c.Args().First())
						},
					},
				},
			},
			{
				Name:  "device",
				Usage: "Playback device",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List playback device",
						Action: func(c *cli.Context) error {
							devices, err := spotifyctl.DeviceList()
							if err != nil {
								return err
							}

							for _, device := range devices {
								fmt.Printf("%s \n", device)
							}

							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
