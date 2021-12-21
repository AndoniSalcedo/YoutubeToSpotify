package services

import (
	"bytes"
	"encoding/json"
	"main/config"
	"main/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"main/youtube/services/body/createplaylist"
	"main/youtube/services/body/findplaylist"
	"main/youtube/services/body/song"
	"main/youtube/services/body/token"
)

func GetToken(code string) (body token.Body) {
	//Form-encode data
	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:8080/callbackY")
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.YoutubeClientID)
	data.Set("client_secret", config.YoutubeClientSecret)
	//Create the request
	req, err := http.NewRequest("POST", "https://accounts.google.com/o/oauth2/token", strings.NewReader(data.Encode()))
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
	//Set up struct
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &body)

	return body
}

func GetPlaylistItem(token string, playlistId string, nextPageToken string) (body findplaylist.Body) {
	//Fetch url
	url := "https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=" + playlistId + "&fields=pageInfo%2CnextPageToken%2C%20items%2Fsnippet%2Ftitle"
	//If is not the first request add pageToken
	if nextPageToken == "" {
		url += "&pageToken=" + nextPageToken
	}
	//Make peticion
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err)
	//Add headers
	req.Header.Add("Authorization", "Bearer "+token)
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
	//Set up struct
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &body)
	return body
}

func FindSong(token string, song string) (body song.Body) {
	//Fetch url
	url := "https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=1&q=" + strings.Replace(song, " ", "%20", -1) + "&fields=items%2Fid%2FvideoId&"
	//Create the request
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err)
	//Add headers
	req.Header.Add("Authorization", "Bearer "+token)
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
	//Set up struct
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &body)

	return body
}

type data struct {
	Snippet struct {
		Title string `json:"title"`
	} `json:"snippet"`
}

func CreatePlaylist(token string, name string, description string) (body createplaylist.Body) {
	//Json-encode data
	values := data{}
	values.Snippet.Title = name
	jsonValue, _ := json.Marshal(values)
	//Create the request
	req, err := http.NewRequest("POST", "https://youtube.googleapis.com/youtube/v3/playlists?part=snippet", bytes.NewBuffer(jsonValue))
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
	//Set up struct
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &body)

	return body
}

type data1 struct {
	Snippet struct {
		PlaylistID string `json:"playlistId"`
		ResourceID struct {
			VideoID string `json:"videoId"`
			Kind    string `json:"kind"`
		} `json:"resourceId"`
	} `json:"snippet"`
}

func AddSongToPlaylist(token string, playlistID string, videoID string) {
	//Json-encode data
	values := data1{}
	values.Snippet.PlaylistID = playlistID
	values.Snippet.ResourceID.VideoID = videoID
	values.Snippet.ResourceID.Kind = "youtube#video"
	jsonValue, _ := json.Marshal(values)
	//Create the request
	req, err := http.NewRequest("POST", "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet", bytes.NewBuffer(jsonValue))
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
}
