package simple_app

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type RunnerFunc = func(logger *zap.SugaredLogger, app Application) error

func Run(runner RunnerFunc, opts ...OptionsFunc) {
	if runner == nil {
		panic("nil runner func")
	}

	logger, err := zap.NewProduction(zap.AddCaller())
	if err != nil {
		panic("Unable to create logger " + err.Error())
	}

	sugar := logger.Sugar()
	defer logger.Sync()

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
