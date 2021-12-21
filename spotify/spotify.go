package spotify

import (
	"fmt"
	"main/spotify/services"
	"main/spotify/services/body/token"
	"main/utils"
	"strings"
)

type User struct {
	userID    string
	userToken string
	userName  string
}

func NewUser(body token.Body) (user User) {
	me := services.GetMe(body.AccessToken)
	user.userID = me.ID
	user.userName = me.DisplayName
	user.userToken = body.AccessToken

	return user
}

func (user *User) CreatePlaylist(playlist utils.Playlist) {
	body := services.CreatePlaylist(user.userID, user.userToken, playlist.Name, "Default description")

	var tracks string

	for _, video := range playlist.Videos {

		body := services.FindSong(user.userToken, video)
		if len(body.Tracks.Items) > 0 {
			tracks += body.Tracks.Items[0].URI + ","
		} else {

			//TODO: cambiar esto
			aux := strings.Split(video, "-")
			if len(aux) > 1 {
				body = services.FindSong(user.userToken, aux[1])
				if len(body.Tracks.Items) > 0 {
					tracks += body.Tracks.Items[0].URI + ","
				} else {
					fmt.Println("not found..." + video)
				}
			}

		}
	}

	services.AddSongToPlaylist(body.ID, user.userToken, tracks)
}

func (user *User) GetPlaylist(url string) (playlist utils.Playlist) {

	//url, err := parseUrl(url)
	//utils.HandleError(err)

	body := services.GetPlaylistItem(url, user.userToken)

	playlist.Name = body.Name

	playlist.Videos = make([]string, 0)

	for _, v := range body.Tracks.Items {
		playlist.Videos = append(playlist.Videos, v.Track.Name)

	}

	return playlist
}
