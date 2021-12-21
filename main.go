package main

import (
	"errors"
	"flag"

	"main/config"
	"main/spotify"
	"main/utils"
	"main/youtube"
)

func main() {

	var url string
	var typo int

	flag.StringVar(&url, "u", "", config.FlagUrl)
	flag.IntVar(&typo, "t", 1, config.FlagTypo)

	flag.Parse()

	if url == "" {
		utils.HandleError(errors.New("-u string \n\tNo url provided"))
	}

	switch typo {
	case config.YoutubeToSpotify:
		authYoutube := youtube.New()
		userYoutube := authYoutube.Auth()

		playlist := userYoutube.GetPlaylist(url)

		authSpotify := spotify.New()
		userSpotify := authSpotify.Auth()

		userSpotify.CreatePlaylist(playlist)

	case config.SpotifyToYoutube:
		authSpotify := spotify.New()
		userSpotify := authSpotify.Auth()

		playlist := userSpotify.GetPlaylist(url)

		authYoutube := youtube.New()
		userYoutube := authYoutube.Auth()

		userYoutube.CreatePlaylist(playlist)

	default:
		utils.HandleError(errors.New("should pass never"))
	}

}
