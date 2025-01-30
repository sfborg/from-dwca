package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gnames/gnfmt"
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

func cacheDirFlag(cmd *cobra.Command) {
	cache, _ := cmd.Flags().GetString("cache-dir")
	if cache != "" {
		opts = append(opts, config.OptCacheDir(cache))
	}
}

func zipFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("zip-output")
	if b {
		opts = append(opts, config.OptWithZipOutput(b))
	}
}

func jobsNumFlag(cmd *cobra.Command) {
	jobs, _ := cmd.Flags().GetInt("jobs-number")
	if jobs > 0 {
		opts = append(opts, config.OptJobsNum(jobs))
	}
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

func fieldsNumFlag(cmd *cobra.Command) {
	s, _ := cmd.Flags().GetString("wrong-fields-num")
	switch s {
	case "":
		return
	case "stop":
		opts = append(opts, config.OptWrongFieldsNum(gnfmt.ErrorBadRow))
	case "ignore":
		opts = append(opts, config.OptWrongFieldsNum(gnfmt.SkipBadRow))
	case "process":
		opts = append(opts, config.OptWrongFieldsNum(gnfmt.ProcessBadRow))
	default:
		slog.Warn("Unknown setting for wrong-fields-num, keeping default",
			"setting", s)
		slog.Info("Supported values are: 'stop' (default), 'ignore', 'process'")
	}
}
