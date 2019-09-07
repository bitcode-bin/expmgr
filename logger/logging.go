package logger

type Logger interface {
	Debug(string)
	Info(string)
	Error(string)

	WithFields(fields map[string]interface{}) Logger
}

type Level int

const (
	Debug Level = 1
	Info  Level = 2
	Error Level = 3
)
