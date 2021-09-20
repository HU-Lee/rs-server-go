package handler

import (
	"encoding/json"
	"net/http"
	"rsweb/src/model"
	"rsweb/src/postgres"
)

func GetFamilyHandler(res http.ResponseWriter, req *http.Request) {
	verifyReq(res, req, "/get/family", "POST")

	families, err := postgres.SelectFamily()

	if len(families) <= 0 || err != nil {
		json.NewEncoder(res).Encode(model.FamilyResponse{
			Success:    false,
			FamilyList: nil,
		})
	} else {
		json.NewEncoder(res).Encode(model.FamilyResponse{
			Success:    true,
			FamilyList: families,
		})
	}
}
