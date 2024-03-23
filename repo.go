package canvas_tools

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog/log"

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
		c := new(canvas.Canvas)

		c, err = jsoncanvas.DecodeFile(path)
		if err != nil {
			return nil, fmt.Errorf("can't parse canvas file %s: %w", path, err)
		}
		log.Trace().Str("path", path).Msg("found canvas")

		r.canvases[path] = c
	}
	log.Debug().Int("n", len(r.canvases)).Msg("found canvases")

	return &r, nil
}

func NewRepoFromCanvas(c *canvas.Canvas, path string, overwrite bool) (*Repo, error) {
	info, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("can't stat file: %w", err)
	}
	if info != nil && !overwrite {
		return nil, fmt.Errorf("file already exists: %s", path)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %w", err)
	}
	f.Close()

	r := Repo{
		mutex:    new(sync.RWMutex),
		canvases: make(map[string]*canvas.Canvas),
	}
	r.canvases[path] = c

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

		log.Trace().Str("path", path).Msg("reset canvas")
		r.canvases[path] = c
	}
	log.Debug().Int("n", len(r.canvases)).Msg("reset canvases")

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
		} else {
			log.Trace().Str("path", path).Msg("successfully saved canvas")
		}
	}
	log.Debug().Int("n", len(r.canvases)).Msg("successfully saved canvases")

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
