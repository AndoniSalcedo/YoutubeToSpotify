package config

import (
	"main/utils"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	utils.HandleError(err)

	return os.Getenv(key)
}

//Env variables
var (
	SpotifyClientID     = getEnvVariable("SpotifyClientID")
	SpotifyClientSecret = getEnvVariable("SpotifyClientSecret")
	YoutubeClientID     = getEnvVariable("YoutubeClientID")
	YoutubeClientSecret = getEnvVariable("YoutubeClientSecret")
)

//Flag Msg
var (
	FlagUrl  = "Url of the playlist you want to transform"
	FlagTypo = "type of tranformation you want to do\n\t 1 : Youtube To Spotify \n\t 2 : Spotify To Youtube"
)

//Typo
var (
	YoutubeToSpotify = 1
	SpotifyToYoutube = 2
)
