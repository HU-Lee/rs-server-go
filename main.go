package main

import (
	"rsweb/src/baseball"
	_ "rsweb/src/naver"
	"rsweb/src/postgres"
	"rsweb/src/server"
)

func main() {
	postgres.DatabaseInit()
	baseball.GenerateThree()
	server.ServerStart()
}
