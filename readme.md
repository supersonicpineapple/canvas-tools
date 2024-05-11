# canvas-tools

A work in progress collection of tools working with [jsoncanvas](https://jsoncanvas.org/) files.

## Usage

### Example: commands/example/simple-canvas.go

The [simple-canvas](./commands/example/sample-canvas.go) command makes use of the [go-jsoncanvas](https://github.com/supersonicpineapple/go-jsoncanvas) library.

```
$ go run ./cmd/simple example sample-canvas
```

This example creates a simple canvas sample file and stores it at `./data/sample.canvas`.

### Canvas stats

```
$ go run ./cmd/simple canvas stat data/sample.canvas
Canvas: sample.canvas
	Nodes: 4
		Text nodes: 3
		File nodes: 0
		Link nodes: 0
		Group nodes: 1
	Edges: 2
```

With human readable trace-level logging:

```
$ go run ./cmd/simple --human-readable --logging --log-level -1 canvas ...
```
