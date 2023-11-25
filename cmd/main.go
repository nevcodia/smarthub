package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/api/route"
	"github.com/nevcodia/smarthub/bootstrap"
	"github.com/nevcodia/smarthub/middleware"
	"io"
	"os"
)

func setLogOutput() {
	f, _ := os.Create("go-gin-poc.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	//setLogOutput()
	app := bootstrap.App()

	env := app.Env

	s3Client := app.S3

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	route.Setup(env, s3Client, server)

	var serverAddress string
	if env.Host != "" {
		serverAddress = env.Host
	}
	if env.Port == "" {
		serverAddress += ":80"
	} else {
		serverAddress += ":" + env.Port
	}
	server.Run(serverAddress)
}
