package canvas_tools

import (
	"fmt"
	"os"
	"sync"

	"github.com/supersonicpineapple/go-jsoncanvas"
	"github.com/supersonicpineapple/go-jsoncanvas/canvas"
)

type Repo struct {
	mutex    *sync.RWMutex
	canvases map[string]*canvas.Canvas
}

func NewRepo(path string) (*Repo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("can't stat file: %w", err)
	}

	r := Repo{
		mutex:    new(sync.RWMutex),
		canvases: make(map[string]*canvas.Canvas),
	}

	if info.IsDir() {
		// TODO: search for canvas files in dir and read them all in
		return nil, fmt.Errorf("got dir %s, please specify a file instead", path)
	} else {
		c, err := jsoncanvas.DecodeFile(path)
		if err != nil {
			return nil, fmt.Errorf("can't parse canvas file %s: %w", path, err)
		}
		r.canvases[path] = c
	}

	return &r, nil
}

// Reset reads all canvases from the filesystem again. This method can be used to reset in-memory changes to a canvas
// or to update the state of in-memory state, in case of external modifications to the canvas file.
func (r *Repo) Reset() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for path := range r.canvases {
		// TODO: pool
		c, err := jsoncanvas.DecodeFile(path)
		if err != nil {
			return fmt.Errorf("can't parse canvas file %s: %w", path, err)
		}

		r.canvases[path] = c
	}

	return nil
}

// Commit writes the current state of all canvases tracked by the repo to disk.
func (r *Repo) Commit() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for path, c := range r.canvases {
		// TODO: diff first
		if err := c.Validate(); err != nil {
			return fmt.Errorf("invalid canvas %s: %w", path, err)
		}

		if err := jsoncanvas.EncodeFile(c, path); err != nil {
			return fmt.Errorf("can't encode to file: %w", err)
		}
	}

	return nil
}

func (r *Repo) One() (*canvas.Canvas, error) {
	paths := r.Paths()
	if len(paths) < 1 {
		return nil, fmt.Errorf("repo contains %d canvases", len(paths))
	}
	return r.canvases[paths[0]], nil
}

func (r *Repo) Paths() []string {
	var keys []string
	for key := range r.canvases {
		keys = append(keys, key)
	}
	return keys
}