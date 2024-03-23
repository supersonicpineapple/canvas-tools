package example

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(SampleCanvasCmd)
}

var Cmd = &cobra.Command{
	Use:   "example",
	Short: "Some examples",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
