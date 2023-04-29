package simple_app

type Logger interface {
	Warn(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

type dummyLogger struct{}

func (d *dummyLogger) Warn(s string, a ...any)  {}
func (d *dummyLogger) Error(s string, a ...any) {}
func (d *dummyLogger) Fatal(s string, a ...any) {}
