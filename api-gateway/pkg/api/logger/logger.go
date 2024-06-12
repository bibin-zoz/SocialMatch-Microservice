package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	fmt.Println("logger initializing")
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	fmt.Println("logger setup done.......")
	return logger
}
