package logger

type StructuralLogger interface {
	Info(input string, fields ...any)
	Error(input string, fields ...any)
	Debug(input string, fields ...any)
	Warn(input string, fields ...any)
	Fatal(input string, fields ...any)
	With(args ...any) StructuralLogger
}
