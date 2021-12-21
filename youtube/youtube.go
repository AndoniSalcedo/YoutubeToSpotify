package youtube

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"main/utils"
	"main/youtube/services"
	"main/youtube/services/body/token"
)

type User struct {
	userToken string
}

func parseTitle(video string) string {
	//TODO: Perfeccionar esta funcion

	video = strings.ToLower(video)

	r1 := regexp.MustCompile(`\((.*)\)|\[(.*)\]`)
	r2 := regexp.MustCompile(`[^a-zA-Z0-9\s-]+`)

	video = r1.ReplaceAllString(video, "")
	video = r2.ReplaceAllString(video, "")

	video = strings.Replace(video, "official video", "", -1)
	video = strings.Replace(video, " X ", " ", -1)
	video = strings.Replace(video, "ft", "", -1)
	video = strings.Replace(video, "out now", "", -1)
	video = strings.Replace(video, "feat.", "", -1)
	video = strings.Replace(video, " x ", " ", -1)
	video = strings.Replace(video, "dj", "", -1)
	video = strings.Replace(video, "DJ", "", -1)
	video = strings.Replace(video, "official music video", "", -1)
	video = strings.Replace(video, "videoclip", "", -1)
	video = strings.Replace(video, "vs", "", -1)
	video = strings.Replace(video, "prod", "", -1)
	video = strings.Replace(video, "by", "", -1)
	return video
}
func parseUrl(url string) (string, error) {

	header := `https://www.youtube.com/playlist\?list=`

	r := regexp.MustCompile(header + `[[:alnum:]]+`)

	if !r.MatchString(url) {
		return "", errors.New("URL no valida")
	}

	return strings.Split(url, "=")[1], nil
}

func NewUser(body token.Body) (user User) {
	user.userToken = body.AccessToken

	return user
}

func (user *User) CreatePlaylist(playlist utils.Playlist) {

	plBody := services.CreatePlaylist(user.userToken, playlist.Name, "Description")
	fmt.Println("PlayList created: " + plBody.ID)
	for _, video := range playlist.Videos {
		vdBody := services.FindSong(user.userToken, video)

		if len(vdBody.Items) > 0 {
			fmt.Println("song " + video + " found with ID: " + vdBody.Items[0].ID.VideoID)
			services.AddSongToPlaylist(user.userToken, plBody.ID, vdBody.Items[0].ID.VideoID)
			fmt.Println("Added")
		} else {
			fmt.Println("no encontrada " + video)
		}
	}

}

func (user *User) GetPlaylist(url string) (playlist utils.Playlist) {

	playlistID, err := parseUrl(url)
	utils.HandleError(err)

	body := services.GetPlaylistItem(user.userToken, playlistID, "")

	playlist.Name = "Prueba"

	playlist.Videos = make([]string, 0)

	for len(playlist.Videos) < body.PageInfo.TotalResults {
		for _, video := range body.Items {
			playlist.Videos = append(playlist.Videos, parseTitle(video.Snippet.Title))
		}
		body = services.GetPlaylistItem(user.userToken, playlistID, body.NextPageToken)
	}
	return playlist
}
