package handlers

import (
	"github.com/kien-fsmk/kbts-hackathon/contracts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type TrainingHandler struct {
	logger *logrus.Entry
	config *viper.Viper
}

func (h *TrainingHandler) TriggerTraining(rw http.ResponseWriter, r *http.Request) {
	// TODO: Trigger the training process

	// Return the response
	returnResponse(rw, contracts.BaseResponse{Code: "200", Message: "Success"}, http.StatusOK)
}

func NewTrainingHandler(logger *logrus.Entry, config *viper.Viper) TrainingHandler {
	return TrainingHandler{
		logger: logger,
		config: config,
	}
}
