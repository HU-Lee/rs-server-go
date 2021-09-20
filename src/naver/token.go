package naver

import (
	"context"
	"net/http"
	"rsweb/src/model"
	"rsweb/src/util"
	"time"

	"golang.org/x/oauth2"
)

// 최초 토큰 발급
func GetFirstToken(code string) (string, error) {

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 5 * time.Second}
	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return "", err
	} else {
		return tok.AccessToken, nil
	}
}

// 토큰으로 프로필 받아오기
func GetProfileByToken(access string) model.TokenResponse {

	client := &http.Client{}

	// 프로필 요청 Open Api
	prof_req, _ := http.NewRequest("GET", "https://openapi.naver.com/v1/nid/me", nil)
	prof_req.Header.Set("Host", "openapi.naver.com")
	prof_req.Header.Set("Pragma", "no-cache")
	prof_req.Header.Set("Authorization", "Bearer "+access)
	prof_req.Header.Set("Accept", "*/*")

	prof_res, err := client.Do(prof_req)
	if err != nil {
		return model.TokenResponse{
			IsAuth: false,
		}
	}

	rawData := model.FromNaver{}
	err = util.GetResJson(prof_res, &rawData)
	if err != nil {
		return model.TokenResponse{
			IsAuth: false,
		}
	}

	data := rawData.Response
	data.AccessToken = access
	if data.Id != "" {
		data.IsAuth = true
	}

	return data
}
