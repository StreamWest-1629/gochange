package runner

import (
	"os/exec"
	"time"

	"github.com/streamwest-1629/gochange/engine"
)

type (
	Runner struct {
		config  *engine.RunnerConfig
		command *exec.Cmd
		timer   time.Timer
	}
)
