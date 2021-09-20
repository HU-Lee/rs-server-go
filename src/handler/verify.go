package handler

import (
	"net/http"
)

// Path와 Method 확인
func verifyReq(res http.ResponseWriter, req *http.Request, path string, method string) {
	if req.URL.Path != path || req.Method != method {
		panic("403 Forbidden")
	}
}
