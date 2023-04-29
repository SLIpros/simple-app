package simple_app

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Runner = func(logger *zap.SugaredLogger, app Application) error

func Run(runner Runner, opts ...OptionsFunc) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Unable to create logger " + err.Error())
	}

	sugar := logger.Sugar()
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			panic("Unable to sync logger " + err.Error())
		}
	}(logger)

	app := cli.NewApp()
	for _, opt := range opts {
		opt(app)
	}

	app.Action = func(ctx *cli.Context) error {
		app := NewApp(sugar)
		if err := runner(sugar, &app); err != nil {
			return errors.WithMessage(err, "create app")
		}

		return app.Run()
	}

	if err := app.Run(os.Args); err != nil {
		sugar.Fatal("Unable to run", zap.Error(err))
	}
}
