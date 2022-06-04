package commands

import (
	"github.com/spf13/cobra"

	"github.com/BinacsLee/escheduler/version"
)

var (
	// VersionCmd the version command
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version Command",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Version: %s.%s.%s, CommitHash: %s\n", version.Maj, version.Min, version.Fix, version.GitCommit)
		},
	}
)
