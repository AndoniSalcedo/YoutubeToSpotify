package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"main/config"
	"main/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"main/spotify/services/body/createplaylist"
	"main/spotify/services/body/findplaylist"
	"main/spotify/services/body/me"
	"main/spotify/services/body/song"
	"main/spotify/services/body/token"
)

func GetToken(code string) (body token.Body) {
	//Form-encode data
	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:8080/callbackS")
	data.Set("grant_type", "authorization_code")
	//Create the request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	utils.HandleError(err)
	//Add Headers
	secret := base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientID + ":" + config.SpotifyClientSecret))

	req.Header.Add("Authorization", "Basic "+secret)
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

func GetMe(token string) (body me.Body) {
	//Create the request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	utils.HandleError(err)
	//Add Headers
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

func CreatePlaylist(userID string, token string, name string, description string) (body createplaylist.Body) {
	//Json-encode data
	values := map[string]string{"name": name, "description": description, "public": "false"}
	jsonValue, _ := json.Marshal(values)
	//Create the request
	req, err := http.NewRequest("POST", "https://api.spotify.com/v1/users/"+userID+"/playlists", bytes.NewBuffer(jsonValue))
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
	url := "https://api.spotify.com/v1/search?q=" + strings.Replace(song, " ", "%20", -1) + "&type=track&limit=1"
	//Create the request
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err)
	//Add Headers
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

func AddSongToPlaylist(playlistID string, token string, tracks string) {
	//Create the request
	req, err := http.NewRequest("POST", "https://api.spotify.com/v1/playlists/"+playlistID+"/tracks?uris="+tracks, nil)
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Authorization", "Bearer "+token)
	//Make peticion
	client := &http.Client{}
	res, err := client.Do(req)
	utils.HandleError(err)
	defer res.Body.Close()
	utils.HandleResponse(res)
}

func GetPlaylistItem(playlistID string, token string) (body findplaylist.Body) {
	//Fetch url
	url := "https://api.spotify.com/v1/playlists/" + playlistID + "?market=ES"
	//Create the request
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err)
	//Add Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
