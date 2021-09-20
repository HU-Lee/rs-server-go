package server

import (
	"net/http"
	"os"
	"rsweb/src/handler"

	"github.com/rs/cors"
)

func ServerStart() {
	port := os.Getenv("PORT")

	http.HandleFunc("/naver/login", handler.NaverLoginHandler)
	http.HandleFunc("/naver/auth", handler.NaverTokenHandler)

	http.HandleFunc("/get/family", handler.GetFamilyHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://192.168.0.2:3000",
			"https://hu-lee.github.io",
		},
	})
	handler := cors.Default().Handler(http.DefaultServeMux)

	http.ListenAndServe(":"+port, c.Handler(handler))
}
