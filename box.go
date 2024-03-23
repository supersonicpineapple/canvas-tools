package canvas_tools

import (
	"github.com/supersonicpineapple/go-jsoncanvas/canvas"
)

type Pixel struct {
	X int
	Y int
}

type Box struct {
	TopLeft  Pixel
	BotRight Pixel
}

func NodesBoundingBox(nodes []*canvas.Node) Box {
	var minX, maxY int
	var maxX, minY int

	for i, node := range nodes {
		if i == 0 {
			minX, maxX = node.X, node.X
			minY, maxY = node.Y, node.Y
			continue
		}

		if node.X < minX {
			minX = node.X
		}
		if node.Y < minY {
			minY = node.Y
		}
		if node.X+node.Width > maxX {
			maxX = node.X + node.Width
		}
		if node.Y+node.Height > maxY {
			maxY = node.Y + node.Height
		}
	}

	topLeft := Pixel{
		X: minX,
		Y: minY,
	}
	botRight := Pixel{
		X: maxX,
		Y: maxY,
	}

	return Box{
		topLeft,
		botRight,
	}
}

func (b *Box) Width() int {
	return b.BotRight.X - b.TopLeft.X
}

func (b *Box) Height() int {
	return b.BotRight.Y - b.TopLeft.Y
}
