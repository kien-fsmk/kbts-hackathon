package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kien-fsmk/kbts-hackathon/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	logger     *logrus.Entry
	config     *viper.Viper
	router     *mux.Router
	httpServer *http.Server
}

func NewServer(logger *logrus.Entry, config *viper.Viper) *Server {
	return &Server{
		logger: logger,
		config: config,
		router: router.InitRouter(logger, config),
	}
}

func (s *Server) Start() {
	serverAddress := fmt.Sprintf("%s:%s", s.config.GetString("server.host"), s.config.GetString("server.port"))

	s.logger.Infof("Starting server in address: %v", serverAddress)

	s.httpServer = &http.Server{Addr: serverAddress, Handler: s.router}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				s.logger.Info("Server closed")
			} else {
				s.logger.Fatalf("Error in server: %v", err.Error())
			}
		}
	}()
}

func (s *Server) Stop() {
	s.logger.Warn("Stopping the server")

	tCtx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	err := s.httpServer.Shutdown(tCtx)
	if err != nil {
		s.logger.Errorf("Error in closing the server: %v", err.Error())
	}
}
