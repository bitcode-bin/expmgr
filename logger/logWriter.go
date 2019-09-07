package logger

import "io"

// logWriter wraps defaultLogger and allows the use of a custom io.Writer.
type logWriter struct {
	*defaultLogger
}

func NewLogWriter(w io.Writer) *logWriter {
	lw := &logWriter{}
	lw.defaultLogger = NewDefaultLogger()
	lw.log.SetOutput(w)
	return lw
}
