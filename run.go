package simple_app

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type RunnerFunc = func(app Application) error

func Run(runner RunnerFunc, logger Logger, opts ...OptionsFunc) {
	if runner == nil {
		panic("nil runner func")
	}

	if logger == nil {
		logger = &dummyLogger{}
	}

	app := cli.NewApp()
	for _, opt := range opts {
		opt(app)
	}

	app.Action = func(ctx *cli.Context) error {
		app := NewApp(logger)
		if err := runner(&app); err != nil {
			return errors.WithMessage(err, "create app")
		}

		return app.Run()
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal("Unable to run", err)
	}
}
