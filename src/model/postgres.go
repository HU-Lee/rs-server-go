package model

type FamilyResponse struct {
	FamilyList []Family `json:"familyList"`
	Success    bool     `json:"success"`
}

type Family struct {
	Id   string `json:"id"`
	Role int    `json:"role"`
}
