package handler

import (
	"encoding/json"
	"net/http"
	"rsweb/src/model"
	"rsweb/src/naver"
	"rsweb/src/postgres"
	"rsweb/src/util"
)

func NaverLoginHandler(res http.ResponseWriter, req *http.Request) {
	verifyReq(res, req, "/naver/login", "GET")

	url := naver.GetRedirectUrl()
	json.NewEncoder(res).Encode(model.LoginResponse{
		RedirectUrl: url,
		Success:     true,
	})
}

func NaverTokenHandler(res http.ResponseWriter, req *http.Request) {
	verifyReq(res, req, "/naver/auth", "POST")

	tok_req := model.TokenRequest{}
	err := util.GetReqJson(req, &tok_req)
	if err != nil {
		http.Error(res, "Request Body is Invalid", http.StatusBadRequest)
		return
	}

	var access string
	if tok_req.State == naver.State {
		access, err = naver.GetFirstToken(tok_req.Code)
	} else {
		access = tok_req.AccessToken
	}

	tok_res := naver.GetProfileByToken(access)

	if tok_req.State == naver.State && tok_res.IsAuth && postgres.FindUserByNaverId(tok_res.Id) < 0 {
		postgres.InsertUser(tok_res)
	}
	json.NewEncoder(res).Encode(tok_res)
}
