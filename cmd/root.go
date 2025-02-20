/*
Copyright © 2024 Dmitry Mozzherin <dmozzherin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log/slog"
	"os"
	"path/filepath"

	dwca "github.com/gnames/dwca/pkg"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/sfgaio"
	"github.com/spf13/cobra"
)

var opts []config.Option

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "from-dwca <dwca_file_or_url> <output_file>",
	Short: "Converts Darwin Core Archive to Species File Group Archive.",
	Long: `Takes path to a Darwin Core Archive, extracts data and metadata,
and converts it to Species File Group Archive. The database schema is created
based on a version of sgma schema.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		versionFlag(cmd)
		flags := []flagFunc{
			debugFlag, cacheDirFlag, jobsNumFlag, zipFlag, fieldsNumFlag,
		}
		for _, v := range flags {
			v(cmd)
		}

		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}

		slog.Info("Converting DwCA to SFGA")
		dwcaPath := args[0]
		outputPath := args[1]

		ext := filepath.Ext(outputPath)
		if ext == ".sqlite" {
			opts = append(opts, config.OptWithBinOutput(true))
		}

		cfg := config.New(opts...)

		err = sysio.New(cfg).Init()
		if err != nil {
			slog.Error("Cannot initialize file system", "error", err)
			os.Exit(1)
		}

		sfga := sfgaio.New()
		err = sfga.Create(cfg.CacheSfgaDir)
		if err != nil {
			slog.Error("Cannot create SFGA database", "error", err)
			os.Exit(1)
		}
		_, err = sfga.Connect()
		if err != nil {
			slog.Error("Cannot initialize SFGA database", "error", err)
			os.Exit(1)
		}

		fd := fdwca.New(cfg, sfga)

		var arc dwca.Archive

		slog.Info("Importing DwCA data", "file", dwcaPath)
		arc, err = fd.GetDwCA(dwcaPath)
		if err != nil {
			slog.Error("Cannot get DarwinCore Archive", "error", err)
			os.Exit(1)
		}

		slog.Info("Exporting data to SQLite")
		err = fd.ImportDwCA(arc)
		if err != nil {
			slog.Error("Cannot export data", "error", err)
			os.Exit(1)
		}

		slog.Info("Making SFGArchive")
		err = fd.ExportSFGA(outputPath)
		if err != nil {
			slog.Error("Cannot dump data", "error", err)
			os.Exit(1)
		}

		slog.Info("DwCA data has been imported successfully")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("debug", "d", false, "set debug mode")
	rootCmd.Flags().StringP("cache-dir", "c", "", "cache directory for temporary files")
	rootCmd.Flags().StringP("wrong-fields-num", "w", "",
		`how to process rows with wrong fields number
     choices: 'stop', 'skip', 'process'
     default: 'stop'`)
	rootCmd.Flags().IntP("jobs-number", "j", 0, "number of concurrent jobs")
	rootCmd.Flags().BoolP("zip-output", "z", false, "compress output with zip")
	rootCmd.Flags().BoolP("version", "V", false, "shows app's version")
}
