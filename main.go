package main

import (
	"os"
	"os/signal"
	"strings"

	"github.com/kien-fsmk/kbts-hackathon/server"

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

type Payment struct {
	PaymentID   string    `json:"payment_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   string    `json:"created_at"`
	Customer    Customer  `json:"customer"`
	Recipient   Recipient `json:"recipient"`
}

type Customer struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
}

type Recipient struct {
	RecipientID             string      `json:"recipient_id"`
	Name                    string      `json:"name"`
	Email                   string      `json:"email"`
	BusinessRegistrationNum string      `json:"business_registration_number"`
	BankAccount             BankAccount `json:"bank_account"`
}

type BankAccount struct {
	BankName    string `json:"bank_name"`
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
	Country     string `json:"country"`
}

// Starting a http server
func main() {
	httpServer := server.NewServer(logger, config)

	httpServer.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	httpServer.Stop()
}
