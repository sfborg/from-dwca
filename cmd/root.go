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

	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/spf13/cobra"
)

var opts []config.Option

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "from-dwca",
	Short: "Imports data from Darwin Core Archive to a sqlite database.",
	Long: `Takes path to a Darwin Core Archive, extracts data and metadata,
and imports it into a sqlite database. The database schema is created
based on a version of sgma schema.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag(cmd)
		flags := []flagFunc{debugFlag, rootDirFlag, jobsNumFlag, inMemoryFlag}
		for _, v := range flags {
			v(cmd)
		}

		cfg := config.New(opts...)
		fd := fdwca.New(cfg)
		if len(args) != 1 {
			cmd.Help()
			os.Exit(0)
		}

		path := args[0]
		slog.Info("Importing DwCA data", "file", path)
		err := fd.GetDwCA(path)
		if err != nil {
			slog.Error("Cannot get DarwinCore Archive", "error", err)
			os.Exit(1)
		}

		slog.Info("Exporting data to SQLite", "file", path)
		err = fd.ExportData()
		if err != nil {
			slog.Error("Cannot export data", "error", err)
			os.Exit(1)
		}

		slog.Info("Making SFG Archive", "file", path)
		err = fd.DumpData()
		if err != nil {
			slog.Error("Cannot dump data", "error", err)
			os.Exit(1)
		}

		slog.Info("DwCA data has been imported successfully", "file", path)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.from-dwca.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("debug", "d", false, "set debug mode")
	rootCmd.Flags().StringP("root-dir", "r", "", "root directory for temporary files")
	rootCmd.Flags().IntP("jobs-number", "j", 0, "number of concurrent jobs")
	rootCmd.Flags().BoolP("in-memory", "m", false, "set sqlite database in memory")
}
