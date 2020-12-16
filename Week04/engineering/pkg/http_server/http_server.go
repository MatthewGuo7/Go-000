package http_server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
}

func NewHttpServer() *HttpServer {
	return &HttpServer{Engine:gin.New()}
}

func (r *HttpServer) Start() error  {
	return r.Run()
}

func (r *HttpServer)Shutdown(ctx context.Context)  {
	fmt.Println("shutdown...")
}

func (r *HttpServer)AddRouter(routers ...IRouter)  {
	for _, router := range routers{
		router.AddRouter(r)
	}
}
