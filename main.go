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

// Starting a http server
func main() {
	httpServer := server.NewServer(logger, config)

	httpServer.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	httpServer.Stop()

	// categorizedPayment, err := paymentSvc.CategorizePayment(context.Background(), paymentSvc.RawPayments[29])
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("\n")
	// fmt.Println("Categorized Payment")
	// fmt.Printf("Description: %s\nCategory: %s\n", categorizedPayment.Description, categorizedPayment.Category)

	// categorizedPayments, err := paymentSvc.CategorizePayments(context.Background(), paymentSvc.RawPayments[:50])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, p := range categorizedPayments {
	// 	fmt.Printf("Description: %s\nCategory: %s\n", p.Description, p.Category)
	// }
}
