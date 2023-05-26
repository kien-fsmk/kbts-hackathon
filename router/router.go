package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kien-fsmk/kbts-hackathon/handlers"
	"github.com/kien-fsmk/kbts-hackathon/payment"
	"github.com/kien-fsmk/kbts-hackathon/pkg/go-openai"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitRouter(logger *logrus.Entry, config *viper.Viper) *mux.Router {
	router := mux.NewRouter()

	openaiClient := openai.NewOpenAIClient(logger, "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", "davinci:ft-personal:kbts-2023-05-26-03-12-14")
	paymentSvc := payment.NewPaymentService(logger, openaiClient)

	promptHandler := handlers.NewPaymentHandler(logger, config, paymentSvc)

	router.HandleFunc("/api/v1/payment", promptHandler.PaymentHandler).Methods(http.MethodPost)

	return router
}
