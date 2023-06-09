package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kien-fsmk/kbts-hackathon/contracts"
	"github.com/kien-fsmk/kbts-hackathon/payment"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PaymentHandler struct {
	logger     *logrus.Entry
	config     *viper.Viper
	paymentSvc *payment.PaymentService
}

func (h *PaymentHandler) PaymentHandler(rw http.ResponseWriter, r *http.Request) {
	var request contracts.CreatePaymentRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	categorizedPayment, err := h.paymentSvc.CreatePayment(context.Background(), payment.Payment(request))
	if err != nil {
		h.logger.Errorf("Error in creating payment: %v", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnResponse(rw, &categorizedPayment, http.StatusOK)
}

func NewPaymentHandler(logger *logrus.Entry, config *viper.Viper, paymentSvc *payment.PaymentService) PaymentHandler {
	return PaymentHandler{
		logger:     logger,
		config:     config,
		paymentSvc: paymentSvc,
	}
}
