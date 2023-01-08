package main

import (
	"github.com/atadzan/url-shortener/app/model"
	"github.com/atadzan/url-shortener/app/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
