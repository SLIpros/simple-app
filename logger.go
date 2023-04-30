package simple_app

type Logger interface {
	Info(...any)
	Error(...any)
	Fatal(...any)
}

type dummyLogger struct{}

func (d *dummyLogger) Info(a ...any)  {}
func (d *dummyLogger) Error(a ...any) {}
func (d *dummyLogger) Fatal(a ...any) {}
