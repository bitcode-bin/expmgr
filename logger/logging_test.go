package logger

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	log := NewDefaultLogger()
	log.SetLevel(Debug)

	log.Debug("d")
	log.Info("i")
	log.Error("e")

	log.WithFields(map[string]interface{}{"K": "V"}).Info("ii")
}

func TestLogWriter(t *testing.T) {
	var buf bytes.Buffer

	log := NewLogWriter(&buf)
	log.Debug("d")
	log.Info("i")
	log.Error("e")

	fmt.Println("***")
	fmt.Println(buf.String())
}
