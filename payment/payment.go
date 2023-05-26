package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/kien-fsmk/kbts-hackathon/pkg/go-openai"
	"github.com/sirupsen/logrus"
)

type PaymentService struct {
	logger       *logrus.Entry
	openAIClient *openai.OpenAIClient

	RawPayments         []Payment
	PaymentWithCategory []Payment
}

func NewPaymentService(logger *logrus.Entry, openai *openai.OpenAIClient) *PaymentService {
	payments, _ := loadPaymentFromFile("sample_txn")
	return &PaymentService{
		logger:              logger,
		openAIClient:        openai,
		RawPayments:         payments,
		PaymentWithCategory: make([]Payment, 0),
	}
}

func (p *PaymentService) CreatePayment(payment *Payment) error {
	return nil
}

func (p *PaymentService) CategorizePayment(ctx context.Context, payment Payment) (*Payment, error) {
	prompt := fmt.Sprintf("Classify the category of this description: %s. \n", payment.Description)
	resp, err := p.openAIClient.Completion(ctx, prompt)
	if err != nil {
		p.logger.Errorf("error creating completion: %v", err)
		return nil, err
	}
	cleanedCategory := strings.ReplaceAll(resp, "\\n\\n###\\n\\n ", "")
	category, _ := strconv.Atoi(cleanedCategory)
	payment.Category = CategoryStr[PaymentCategory(category)]
	return &payment, nil
}

func loadPaymentFromFile(fileName string) ([]Payment, error) {
	jsonFile, err := os.Open("sample/" + fileName + ".json")
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

func (p *PaymentService) CategorizePayments(ctx context.Context, payments []Payment) ([]Payment, error) {
	var (
		wg                  sync.WaitGroup
		inputCh             = make(chan Payment, len(payments))
		outputCh            = make(chan error)
		errs                []error
		m                   sync.Mutex
		paymentWithCategory = make([]Payment, 0)
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for payment := range inputCh {

				payment, err := p.CategorizePayment(ctx, payment)
				if err != nil {
					outputCh <- err
					return
				}
				m.Lock()
				paymentWithCategory = append(paymentWithCategory, *payment)
				m.Unlock()
			}
		}()
	}

	// Send the work to the workers
	go func() {
		for _, payment := range payments {
			inputCh <- payment
		}
		close(inputCh)
	}()

	// Wait for the workers to finish and collect any errors
	go func() {
		wg.Wait()
		close(outputCh)
	}()
	for err := range outputCh {
		if err != nil {
			errs = append(errs, err)
		}
	}
	return paymentWithCategory, nil
}
