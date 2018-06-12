package logger

type Logger interface {
	DebugLevel(debug bool)
	SetContext(context map[string]interface{})

	Println(args ...interface{})
	Printf(format string, args ...interface{})
	Info(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})

	Error(msg string, err error)
	Fatal(msg string, err error)
}
