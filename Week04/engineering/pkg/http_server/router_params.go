package http_server

import "github.com/gin-gonic/gin"

type RouterParams struct {
	Method string
	RelativePath string
	f gin.HandlerFunc
}

func (r RouterParams) AddRouter(server *HttpServer) {
	server.Handle(r.Method,r.RelativePath, r.f)
}

func NewRouterParams(method string, relativePath string, f gin.HandlerFunc) *RouterParams {
	return &RouterParams{Method: method, RelativePath: relativePath, f: f}
}


