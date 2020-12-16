package main

import (
	"context"
	"engineering/pkg/http_server"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	//添加rpc service...

	//添加http service
	httpService := http_server.NewHttpServer()

	g.Go(func() error {
		log.Println("http server start")

		initRouter(httpService)

		go func() {
			<- ctx.Done()
			log.Println("http server done")
			httpService.Shutdown(context.TODO())
		}()
		return httpService.Start()
	})

	g.Go(func() error {
		listenSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(listenSignals))
		signal.Notify(sig, listenSignals...)
		for {
			log.Printf("received signal")
			select {
			case <-ctx.Done():
				log.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				return clearWork()
			}
		}
	})

	g.Wait()
	log.Println("service exit")

}

func clearWork() error {
	fmt.Println("start to clear")
	return nil
}