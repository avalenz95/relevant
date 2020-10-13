package logging

import "github.com/sirupsen/logrus"

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
