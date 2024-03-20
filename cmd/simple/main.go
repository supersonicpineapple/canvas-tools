package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/supersonicpineapple/canvas-tools/commands/canvas"
)

var (
	logging       bool
	humanReadable bool
	noColor       bool
	logLevel      int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "simple",
		Short: "Simple cobra cli program to wire up commands.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !logging {
				zerolog.SetGlobalLevel(zerolog.Disabled)
			} else if humanReadable {
				log.Logger = log.Output(zerolog.ConsoleWriter{
					Out:        os.Stderr,
					NoColor:    noColor,
					TimeFormat: time.RFC3339,
				})
			} else {
				zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
				log.Logger = log.Output(os.Stderr)
			}

			zerolog.SetGlobalLevel(zerolog.Level(logLevel))
		},
	}

	fs := rootCmd.PersistentFlags()
	fs.BoolVar(&logging, "logging", false, "Activate logging")
	fs.BoolVar(&humanReadable, "human-readable", false, "Activate human readable logging")
	fs.BoolVar(&noColor, "no-color", false, "No color output for logging (only effective with human-readable==true)")
	fs.IntVar(&logLevel, "log-level", 1, "Set log level: panic=5;fatal=4;error=3;warn=2;info=1(default);debug=0;trace=-1")

	rootCmd.AddCommand(canvas.Cmd)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
