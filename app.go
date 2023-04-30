package simple_app

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/pkg/errors"
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
	logger  Logger
	servers []Server
	closers []Closer
}

func NewApp(logger Logger) App {
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
				serverErrCh <- errors.WithMessagef(err, "serve %q", server.Description())
			}
		}(s)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer close(signalCh)
	defer a.Stop()

	select {
	case <-signalCh:
		a.logger.Info("Shutdown application by signal")
		return nil

	case err := <-serverErrCh:
		return err
	}
}

func (a *App) Stop() {
	var sb strings.Builder
	for i := len(a.closers) - 1; i >= 0; i-- {
		c := a.closers[i]
		if err := c.Close(); err != nil {
			err := errors.WithMessagef(err, "close %q", c.Description())

			if sb.Len() > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(err.Error())
		}
	}

	if sb.Len() > 0 {
		a.logger.Error("Closer errors", sb.String())
	}
}
