package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest-api-server/internal/user"
	"rest-api-server/pkg/logger"
	"time"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)
	start(router, logger)

}

func start(router *httprouter.Router, logger logger.Logger) {

	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port :1234")
	logger.Fatalln(server.Serve(listener))

}
