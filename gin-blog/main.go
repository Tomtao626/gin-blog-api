package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server start failed, err: %v\n", err)
		return
	}
	/*
		go get -u github.com/fvbock/endless
		router := routers.InitRouter()
		endless.DefaultReadTimeOut = setting.ReadTimeout
		endless.DefaultWriteTimeOut = setting.WriteTimeout
		endless.DefaultMaxHeaderBytes = 1 << 20
		endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
		server := endless.NewServer(endPoint, router)
		server.BeforeBegin = func(add string) {
			log.Printf("Actual pid is %d", syscall.Getpid())
		}
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server start failed, err: %v\n", err)
		}
	*/
	/*
		router := routers.InitRouter()
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        router,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20}
		go func() {
			if err := s.ListenAndServe(); err != nil {
				log.Printf("Server start failed, err: %v\n", err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("shutdown server...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown failed, err: %v\n", err)
		}
		log.Println("Server shutdown success")
	*/
}
