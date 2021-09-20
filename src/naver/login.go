package naver

import (
	"golang.org/x/oauth2"
)

func GetRedirectUrl() string {
	return conf.AuthCodeURL(State, oauth2.AccessTypeOffline)
}
