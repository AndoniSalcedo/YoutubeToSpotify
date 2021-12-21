package youtube

import (
	"errors"
	"net/http"
	"os/exec"
	"runtime"

	"main/config"
	"main/utils"
	"main/youtube/services"

	"main/youtube/services/body/token"
)

var (
	scope = "https://www.googleapis.com/auth/youtube"
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
	http.HandleFunc("/callbackY", auth.callback)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	auth.ch = make(chan *token.Body)

	return auth
}

func (auth *Authorizator) Auth() (user User) {
	//Fetch url
	url := "https://accounts.google.com/o/oauth2/v2/auth?scope=" + scope + "&response_type=code&redirect_uri=http://localhost:8080/callbackY&client_id=" + config.YoutubeClientID
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
