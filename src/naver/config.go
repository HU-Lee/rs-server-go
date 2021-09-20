package naver

import (
	"os"

	"golang.org/x/oauth2"
)

const (
	clientURL = "https://hu-lee.github.io/rs-client"
	// clientURL   = "http://192.168.0.2:3000/rs-client"
	State       = "ROLLINGstone"
	CallbackURL = clientURL + "/"
	AuthURL     = "https://nid.naver.com/oauth2.0/authorize"
	TokenURL    = "https://nid.naver.com/oauth2.0/token"
)

var conf *oauth2.Config

func init() {
	// activate only on local
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Config Load Failed...")
	// }

	conf = &oauth2.Config{
		ClientID:     os.Getenv("NAVER_CLIENT_ID"),
		ClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
		RedirectURL: CallbackURL,
	}
}
