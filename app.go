package simple_app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Server interface {
	Description() string
	Serve() error
}

type Closer interface {
	Description() string
	Close() error
}

type Application interface {
	AddServer(servers ...Server)
	AddCloser(closers ...Closer)
}

type App struct {
	logger  *zap.SugaredLogger
	servers []Server
	closers []Closer
}

func NewApp(logger *zap.SugaredLogger) App {
	return App{
		logger: logger,
	}
}

func (a *App) AddServer(servers ...Server) {
	a.servers = append(a.servers, servers...)

	for _, s := range a.servers {
		c, ok := s.(Closer)
		if !ok {
			continue
		}

		a.closers = append(a.closers, c)
	}
}

func (a *App) AddCloser(closers ...Closer) {
	a.closers = append(a.closers, closers...)
}

func (a *App) Run() error {
	serverErrCh := make(chan error)

	for _, s := range a.servers {
		go func(server Server) {
			if err := server.Serve(); err != nil {
				serverErrCh <- errors.WithMessagef(err, "serving %q",
					server.Description())
			}
		}(s)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer close(signalCh)
	defer a.Stop()

	select {
	case <-signalCh:
		a.logger.Warn("Shutdown application by signal")
		return nil

	case err := <-serverErrCh:
		return err
	}
}

func (a *App) Stop() {
	closerErrors := make([]error, 0, len(a.closers))
	for i := len(a.closers) - 1; i >= 0; i-- {
		c := a.closers[i]
		if err := c.Close(); err != nil {
			closerErrors = append(closerErrors,
				errors.WithMessagef(err, "closing %q", c.Description()))
		}
	}

	if len(closerErrors) > 0 {
		a.logger.Error("Closer errors",
			zap.Errors("errors", closerErrors))
	}
}
