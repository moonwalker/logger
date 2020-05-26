package logger

type Logger interface {
	UseTextFormatter()
	SetLevel(level int)
	SetContext(context map[string]interface{})

	Println(args ...interface{})
	Printf(format string, args ...interface{})

	Trace(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warning(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}
