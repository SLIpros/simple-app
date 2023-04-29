package simple_app

import "github.com/urfave/cli/v2"

type OptionsFunc = func(cli *cli.App)

func SetCLIFlags(cliFlags []cli.Flag) OptionsFunc {
	return func(cli *cli.App) {
		cli.Flags = cliFlags
	}
}

func SetVersion(version Version) OptionsFunc {
	return func(cli *cli.App) {
		cli.Name = version.Name
		cli.Version = version.String()
	}
}
