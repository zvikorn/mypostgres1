package main

import (
	"mypostgres1/cmd/apiserver"
)

func main() {
	server := apiserver.NewServer()
	server.Init()
}
