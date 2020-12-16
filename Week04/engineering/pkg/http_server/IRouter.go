package http_server

type IRouter interface {
	AddRouter(server *HttpServer)
}
