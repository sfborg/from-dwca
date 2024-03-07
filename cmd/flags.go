package cmd

import (
	"fmt"
	"log/slog"
	"os"

	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/spf13/cobra"
)

type flagFunc func(cmd *cobra.Command)

func debugFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("debug")
	if b {
		lopts := &slog.HandlerOptions{Level: slog.LevelDebug}
		handle := slog.NewJSONHandler(os.Stderr, lopts)
		logger := slog.New(handle)
		slog.SetDefault(logger)
	}
}

func rootDirFlag(cmd *cobra.Command) {
	root, _ := cmd.Flags().GetString("root-dir")
	if root != "" {
		opts = append(opts, config.OptRootPath(root))
	}
}

func jobsNumFlag(cmd *cobra.Command) {
	jobs, _ := cmd.Flags().GetInt("jobs-number")
	if jobs > 0 {
		opts = append(opts, config.OptJobsNum(jobs))
	}
}

func inMemoryFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("in-memory")
	opts = append(opts, config.OptInMemory(b))
}

func versionFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("version")
	if b {
		version := fdwca.GetVersion()
		fmt.Printf(
			"\nVersion: %s\nBuild:   %s\n\n",
			version.Version,
			version.Build,
		)
		os.Exit(0)
	}
}
