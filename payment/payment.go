package payment

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/kien-fsmk/kbts-hackathon/pkg/go-openai"
	"github.com/sirupsen/logrus"
)

type PaymentService struct {
	logger          *logrus.Entry
	openAIClient    *openai.OpenAIClient
	PaymentDatabase map[string]Payment
}

func NewPaymentService(logger *logrus.Entry, openai *openai.OpenAIClient) *PaymentService {
	return &PaymentService{
		logger:          logger,
		openAIClient:    openai,
		PaymentDatabase: make(map[string]Payment),
	}
}

func (p *PaymentService) CreatePayment(ctx context.Context, payment Payment) (*Payment, error) {
	id := uuid.New().String()
	categorizedPayment, err := p.CategorizePayment(context.Background(), payment)
	if err != nil {
		return nil, err
	}
	p.PaymentDatabase[id] = *categorizedPayment
	return categorizedPayment, nil
}

func (p *PaymentService) CategorizePayment(ctx context.Context, payment Payment) (*Payment, error) {
	prompt := fmt.Sprintf("Classify the category of this description: `%s`", payment.Description)
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
