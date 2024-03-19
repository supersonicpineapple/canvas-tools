package canvas

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/supersonicpineapple/canvas-tools"
)

var CatCmd = &cobra.Command{
	Use:   "cat",
	Short: "simple cat program for canvas (wip)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			canvasFilePath = args[0]
		}

		r, err := canvas_tools.NewRepo(canvasFilePath)
		if err != nil {
			return fmt.Errorf("can't open repo: %w", err)
		}

		c, err := r.One()
		if err != nil {
			return fmt.Errorf("can't get one canvas: %w", err)
		}

		// TODO: format flag: spew-dump, json, etc.
		spew.Dump(c)
		return nil
	},
}
