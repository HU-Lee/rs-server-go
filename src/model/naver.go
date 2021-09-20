package model

type LoginResponse struct {
	RedirectUrl string `json:"redirectUrl"`
	Success     bool   `json:"success"`
}

// 실제 API에서는 refresh_token으로 토큰 재발급과
// service_provider로 토큰 삭제가 있지만 여기서는 사용 X
type TokenRequest struct {
	AccessToken string `json:"accessToken"`
	State       string `json:"state"`
	Code        string `json:"code"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsAuth      bool   `json:"isAuth"`
}

// 네이버 프로필 요청시 받을 format
type FromNaver struct {
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	Response   TokenResponse
}
