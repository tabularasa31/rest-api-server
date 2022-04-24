package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api-server/internal/config"
	"rest-api-server/internal/user"
	"rest-api-server/pkg/logger"
	"time"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	config := config.GetConfig(logger)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)
	start(router, logger, config)
}

func start(router *httprouter.Router, logger *logger.Logger, config *config.Config) {

	logger.Info("start application")

	var listener net.Listener
	var listenError error

	if config.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatalln(err)
		}

		logger.Info("create socket")
		socketPAth := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPAth)

		logger.Info("listen unix socket")
		listener, listenError = net.Listen("unix", socketPAth)
		logger.Infof("server is listening unix socket %s", socketPAth)

	} else {
		logger.Info("listen tcp")
		listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%s", config.Listen.BindIp, config.Listen.Port))
		logger.Infof("server is listening port %s:%s", config.Listen.BindIp, config.Listen.Port)

	}

	if listenError != nil {
		logger.Fatal(listenError)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))

}
