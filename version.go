package simple_app

import "fmt"

type Version struct {
	GitCommit string
	GitBranch string
	Version   string
	BuildDate string
	Name      string
}

func (d *Version) String() string {
	return fmt.Sprintf("Name: %s, Commit: %s, branch: %s, version: %s, build date: %s",
		d.Name,
		d.GitCommit,
		d.GitBranch,
		d.Version,
		d.BuildDate,
	)
}
