package main

import (
	"github.com/kien-fsmk/kbts-hackathon/server"
	"os"
	"os/signal"
	"strings"

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
}
