package lint

import "flag"

// Flags represents a set of lint flag parameters.
type Flags struct {
	*flag.FlagSet

	ConfigDirPath string
	FixMode       bool
}

// NewFlags creates a new Flags.
func NewFlags(
	args []string,
) Flags {
	f := Flags{
		FlagSet: flag.NewFlagSet("lint", flag.ExitOnError),
	}

	f.StringVar(
		&f.ConfigDirPath,
		"config_dir_path",
		"",
		"path/to/protolint.yaml",
	)
	f.BoolVar(
		&f.FixMode,
		"fix",
		false,
		"mode that the command line can automatically fix some of the problems",
	)
	_ = f.Parse(args)
	return f
}
