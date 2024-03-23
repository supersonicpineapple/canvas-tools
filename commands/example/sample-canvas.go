package example

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/supersonicpineapple/go-jsoncanvas/canvas"

	"github.com/supersonicpineapple/canvas-tools"
)

var SampleCanvasCmd = &cobra.Command{
	Use:   "sample-canvas",
	Short: "Create a sample canvas.",
	Run: func(cmd *cobra.Command, args []string) {
		dataDir := "data"
		if err := os.MkdirAll(dataDir, 0775); err != nil {
			log.Fatal().Err(err).Str("dir", dataDir).Msg("can't ensure dir exists")
		}

		c := canvas.NewCanvas()

		fooNode := canvas.NewNode().SetText("foo").
			TranslateY(canvas.DefaultHeight + canvas.DefaultGap)
		barNode := canvas.NewNode().SetText("bar").
			TranslateY(canvas.DefaultHeight + canvas.DefaultGap).
			TranslateX(canvas.DefaultWidth + canvas.DefaultGap)

		c.AddNodes(
			canvas.NewNode().SetText("hello world."),
			fooNode,
			barNode,
		)

		// add a group that encapsulates all nodes.
		b := canvas_tools.NodesBoundingBox(c.Nodes)
		groupNode := canvas.NewNode().SetGroup("Test Group").
			SetWidth(b.Width() + canvas.DefaultGap).SetHeight(b.Height() + canvas.DefaultGap).
			TranslateX(-canvas.DefaultGap / 2).TranslateY(-canvas.DefaultGap / 2)
		c.AddNodes(groupNode)

		c.AddEdges(
			canvas.NewEdge(fooNode, barNode, "right", "left"),
			canvas.NewEdge(barNode, c.Nodes[0], "top", "right"),
		)

		r, err := canvas_tools.NewRepoFromCanvas(c, "data/sample.canvas", true)
		if err != nil {
			log.Fatal().Err(err).Msg("can't open repo with canvas")
		}

		if err := r.Commit(); err != nil {
			log.Fatal().Err(err).Msg("can't commit repo")
		}
	},
}
