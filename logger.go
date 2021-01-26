package logger

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Fields is a map used to log custom fields.
type Fields map[string]interface{}

//Init initializes the logger to be used in any package.
func init() {
	logger = logrus.New()

	dsn, exists := os.LookupEnv("SENTRY_DSN")

	if exists {
		hook, err := logrus_sentry.NewSentryHook(dsn, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		if err != nil {
			logger.Errorf("It was not possible to enable Sentry handler: %w", err)
			return
		}

		hook.StacktraceConfiguration.Enable = true

		logger.Hooks.Add(hook)
	}
}

// WithFields log custom fields.
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logger.WithFields(fields)
}

// Info log the args as informative data.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Error log the args as an error.
func Error(args ...interface{}) {
	logger.Error(args...)
}
