package trigger

import (
	"crypto/md5"
	"time"

	"github.com/streamwest-1629/gochange/engine"
)

type (
	Runner interface {
		OnChanged(mode, path string)
	}
	FileInfo struct {
		freqConfig   *engine.FreqConfig
		prevCheckSum [md5.BlockSize]byte
		timer        time.Timer
		runners      []Runner
	}
	DirInfo struct {
		prevDirs []string
		timer    time.Timer
	}
	Triggers struct {
		configs  []engine.TriggerConfig
		fileInfo map[string]FileInfo
		dirInfo  map[string]DirInfo
	}
)

func NewTrigger(configs ...engine.TriggerConfig) *Triggers {
	for _, config := range configs {

	}
}
