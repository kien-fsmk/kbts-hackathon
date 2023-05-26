package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kien-fsmk/kbts-hackathon/pkg/go-openai"
	"github.com/sirupsen/logrus"
)

type PaymentService struct {
	logger       *logrus.Entry
	openAIClient *openai.OpenAIClient
}

func NewPaymentService(logger *logrus.Entry, openai *openai.OpenAIClient) *PaymentService {
	return &PaymentService{
		logger:       logger,
		openAIClient: openai,
	}
}

func (p *PaymentService) CreatePayment(payment *Payment) error {
	return nil
}

func (p *PaymentService) CategorizePayment(ctx context.Context, payment Payment) (*Payment, error) {
	prompt := payment.Description
	p.openAIClient.Completion(ctx, prompt)
}

func (p *PaymentService) loadPaymentFromFile(paymentID string) ([]Payment, error) {
	jsonFile, err := os.Open("sample/sample_payments.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var payments []Payment
	err = json.Unmarshal(byteValue, &payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
