package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/supersonicpineapple/canvas-tools/commands/canvas"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "simple",
		Short: "Simple cobra cli program to wire up commands.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(canvas.Cmd)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
	}
}
