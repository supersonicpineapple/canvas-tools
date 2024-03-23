# canvas-tools

A work in progress collection of tools working with [jsoncanvas](https://jsoncanvas.org/) files.

## Usage

### Canvas stats

```
$ go run ./cmd/simple canvas stat sample.canvas
Canvas: sample.canvas
        Nodes: 4
                Text nodes: 1
                File nodes: 3
                Link nodes: 0
                Group nodes: 0
        Edges: 1
```

With human readable trace-level logging:
```
$ go run ./cmd/simple --human-readable --logging --log-level -1 canvas ...
```