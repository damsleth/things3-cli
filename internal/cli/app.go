package cli

import (
	"io"
	"os"

	"github.com/ossianhempel/things3-cli/internal/open"
)

// Launcher defines the interface for opening Things URLs.
type Launcher interface {
	Open(args ...string) error
}

// App holds shared dependencies for commands.
type App struct {
	In         io.Reader
	Out        io.Writer
	Err        io.Writer
	Launcher   Launcher
	Debug      bool
	Foreground bool
	DryRun     bool
}

// NewApp builds the default application wiring.
func NewApp() *App {
	debug := os.Getenv("DEBUG") != ""
	return &App{
		In:         os.Stdin,
		Out:        os.Stdout,
		Err:        os.Stderr,
		Launcher:   open.NewFromEnv(os.Stdout, os.Stderr),
		Debug:      debug,
		Foreground: false,
		DryRun:     false,
	}
}
