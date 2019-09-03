package main

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
}

func NewLogger() Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	return log
}
