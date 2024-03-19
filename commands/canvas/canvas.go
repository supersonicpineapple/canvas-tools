package canvas

import (
	"github.com/spf13/cobra"
)

var (
	canvasFilePath string
)

func init() {
	Cmd.AddCommand(CatCmd)
	Cmd.AddCommand(StatCmd)
}

var Cmd = &cobra.Command{
	Use:   "canvas",
	Short: "canvas utils",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
