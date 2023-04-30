package simple_app

type Logger interface {
	Info(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

type dummyLogger struct{}

func (d *dummyLogger) Info(s string, a ...any)  {}
func (d *dummyLogger) Error(s string, a ...any) {}
func (d *dummyLogger) Fatal(s string, a ...any) {}
