package logger

import "github.com/sirupsen/logrus"

func (logger *Logger) WithField(key string, value interface{}) *Logger {
	logger.logrus = logger.logrus.WithField(key, value)
	return logger
}

func (logger *Logger) WithFields(fields Fields) *Logger {
	logger.logrus = logger.logrus.WithFields(logrus.Fields(fields))
	return logger
}

func (logger *Logger) WithRequestID(requestID string) *Logger {
	return logger.WithField("request_id", requestID)
}

func (logger *Logger) WithURI(url string) *Logger {
	return logger.WithField("url", url)
}
func (logger *Logger) WithUserID(userID string) *Logger {
	return logger.WithField("user_id", userID)
}

func (logger *Logger) WithError(err error) *Logger {
	return logger.WithField("error", err)
}

func (logger *Logger) WithAuthType(authType string) *Logger {
	return logger.WithField("auth_type", authType)
}

func (logger *Logger) WithAuthValue(authValue string) *Logger {
	return logger.WithField("auth_value", authValue)
}

func (logger *Logger) WithConfirmationCode(confirmationCode string) *Logger {
	return logger.WithField("confirmation_code", confirmationCode)
}

func (logger *Logger) WithStackTrace(stackTrace string) *Logger {
	return logger.WithField("stack_trace", stackTrace)
}
