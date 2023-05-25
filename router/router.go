package router

import (
	"github.com/gorilla/mux"
	"github.com/kien-fsmk/kbts-hackathon/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func InitRouter(logger *logrus.Entry, config *viper.Viper) *mux.Router {
	router := mux.NewRouter()

	trainingHandler := handlers.NewTrainingHandler(logger, config)
	promptHandler := handlers.NewPromptHandler(logger, config)

	router.HandleFunc("/api/v1/training/initiate", trainingHandler.TriggerTraining).Methods(http.MethodPost)

	router.HandleFunc("/api/v1/prompt", promptHandler.HandlerUserPrompts).Methods(http.MethodPost)

	return router
}
