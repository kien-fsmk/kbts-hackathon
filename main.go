package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/kien-fsmk/kbts-hackathon/payment"
	"github.com/kien-fsmk/kbts-hackathon/pkg/go-openai"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	serviceName = "kbts-backend"
	version     = "0.0.1"
	config      *viper.Viper
	logger      *logrus.Entry
)

func newConfig() (*viper.Viper, error) {
	config := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	return config, nil
}

func init() {
	var err error
	logger = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
		"service": serviceName,
		"version": version,
	})

	config, err = newConfig()
	if err != nil {
		logger.WithError(err).Fatalf("error loading configuration \n")
	}
}

// Starting a http server
func main() {
	// httpServer := server.NewServer(logger, config)

	// httpServer.Start()

	// c := make(chan os.Signal)
	// signal.Notify(c, os.Interrupt, os.Kill)

	// <-c

	// httpServer.Stop()

	openaiClient := openai.NewOpenAIClient(logger, "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", "davinci:ft-personal:kbts-2023-05-26-03-12-14")
	paymentSvc := payment.NewPaymentService(logger, openaiClient)

	objStr := `
	{
		"payment_id": "a50798bf-4520-4e16-92c4-29e68c962d2b",
		"amount": 284769,
		"description": "concert ticket purchase",
		"status": "success",
		"created_at": "2023-03-14",
		"customer": {
			"customer_id": "d516c34c-0c02-4758-952e-082ae8af0804",
			"email": "mklossek3w@army.mil"
		},
		"recipient": {
			"recipient_id": "8e9ac67a-2543-41b9-a35a-9be7970aab20",
			"name": "Aimee Lanchberry",
			"email": "mklossek3w@thetimes.co.uk",
			"business_registration_number": "39-3459453",
			"bank_account": {
				"bank_name": "Citibank of Singapore",
				"account_no": "4017952912145",
				"account_name": "Matthaeus Klossek",
				"country": "SG"
			}
		}
	}
	`

	var payment payment.Payment
	json.Unmarshal([]byte(objStr), &payment)

	paymentSvc.CategorizePayment(context.Background(), payment)
}
