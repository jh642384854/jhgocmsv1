package utils

import "github.com/sirupsen/logrus"

/**
	Logger的初始化示例
 */
func NewLogrusLogger() *logrus.Logger  {
	logrusInstance := logrus.New()
	return logrusInstance
}