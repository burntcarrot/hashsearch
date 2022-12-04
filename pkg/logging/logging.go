package logging

import (
	"io"
	"os"

	"github.com/burntcarrot/hashsearch/pkg/config"
	"github.com/sirupsen/logrus"
)

// Logger is a globally-available logger.
var Logger *logrus.Logger

// InitLogger returns a logrus logger with
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Set JSON as the format.
	logger.Formatter = &logrus.JSONFormatter{}

	// Read file.
	file, err := os.OpenFile(config.LOGGING_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		logger.Fatal(err)
	}

	// Set two output streams - stderr and file.
	logger.SetOutput(io.MultiWriter(os.Stderr, file))

	// Set exit handler for closing the file.
	logrus.RegisterExitHandler(func() {
		if file == nil {
			return
		}
		file.Close()
	})

	return logger
}
