package spotify

import (
	"errors"
	"net/http"
	"os/exec"
	"runtime"

	"main/config"
	"main/spotify/services"
	"main/utils"

	"main/spotify/services/body/token"
)

var (
	scope = "playlist-modify-public playlist-modify-private"
)

type Authorizator struct {
	ch chan *token.Body
}

func (auth *Authorizator) callback(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	token := services.GetToken(code)

	auth.ch <- &token
}

func New() (auth Authorizator) {
	http.HandleFunc("/callbackS", auth.callback)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	auth.ch = make(chan *token.Body)

	return auth
}

func (auth *Authorizator) Auth() (user User) {
	//Fetch url
	url := "https://accounts.spotify.com/authorize?client_id=" + config.SpotifyClientID + "&redirect_uri=http://localhost:8080/callbackS&response_type=code&scope=" + scope
	//Open auth window in the default browser
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errors.New("unsoported platform")
	}
	utils.HandleError(err)

	//Wait complete auth
	token := <-auth.ch
	//Get user information
	user = NewUser(*token)

	return user
}
