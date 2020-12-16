package main

import (
	"engineering/apis/web"
	"engineering/pkg/http_server"
)

func RegisterAPI(server *http_server.HttpServer, router http_server.IRouter)  {
	server.AddRouter(router)
}

func initRouter(server *http_server.HttpServer)  {
	RegisterAPI(server, web.NewUserAPI(InitUserAPI()))
}


