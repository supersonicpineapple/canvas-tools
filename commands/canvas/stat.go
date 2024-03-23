package canvas

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/supersonicpineapple/go-jsoncanvas/canvas"

	"github.com/supersonicpineapple/canvas-tools"
)

var StatCmd = &cobra.Command{
	Use:   "stat",
	Short: "show some canvas stats",
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

		fmt.Printf("Canvas: %s\n", filepath.Base(canvasFilePath))

		fmt.Printf("\tNodes: %d\n", len(c.Nodes))
		fmt.Print(nodeDetails(c))

		fmt.Printf("\tEdges: %d\n", len(c.Edges))

		b := canvas_tools.NodesBoundingBox(c.Nodes)

		fmt.Printf("\n\tMeta:\n")
		fmt.Printf("\t\tArea: %d pixel² %v\n", b.Area(), b)
		fmt.Printf("\t\tNodes density: %f\n", float64(len(c.Nodes)*1000000)/float64(b.Area()))

		return nil
	},
}

func nodeDetails(c *canvas.Canvas) string {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\t\tText nodes: %d\n", len(c.FilterNodes().ByType("text"))))
	buf.WriteString(fmt.Sprintf("\t\tFile nodes: %d\n", len(c.FilterNodes().ByType("file"))))
	buf.WriteString(fmt.Sprintf("\t\tLink nodes: %d\n", len(c.FilterNodes().ByType("link"))))
	buf.WriteString(fmt.Sprintf("\t\tGroup nodes: %d\n", len(c.FilterNodes().ByType("group"))))

	return buf.String()
}
