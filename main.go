package main

import (
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//捕获panic并打印日志
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	Route := gin.Default()      //获取gin引擎实例
	routers.ARouterGroup(Route) //注册a路由
	routers.BRouterGroup(Route) //注册b路由

	s := &http.Server{
		Addr:    ":8099",
		Handler: Route,
	}

	go func() {
		s.ListenAndServe()
	}()

	//退出服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-quit
	if err := s.Shutdown(nil); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
