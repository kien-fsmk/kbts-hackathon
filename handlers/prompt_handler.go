package handlers

import (
	"github.com/kien-fsmk/kbts-hackathon/contracts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type PromptHandler struct {
	logger *logrus.Entry
	config *viper.Viper
}

func (h *PromptHandler) HandlerUserPrompts(rw http.ResponseWriter, r *http.Request) {
	// Return the response
	returnResponse(rw, contracts.BaseResponse{Code: "200", Message: "Success"}, http.StatusOK)
}

func NewPromptHandler(logger *logrus.Entry, config *viper.Viper) PromptHandler {
	return PromptHandler{
		logger: logger,
		config: config,
	}
}
