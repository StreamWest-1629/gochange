package engine

import (
	"errors"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v2"
)

type (
	FreqConfig struct {
		MaxWaitMs int
		MinWaitMs int
	}

	TriggerConfig struct {
		IncludeExts []string    `yaml:"includeExts,omitempty"`
		ExcludeDirs []string    `yaml:"excludeDirs,omitempty"`
		Runners     []string    `yaml:"runners,omitempty"`
		Frequenty   *FreqConfig `yaml:"frequenty,omitempty"`
	}

	RunnerConfig struct {
		Cmds         []string          `yaml:"commands,omitempty"`
		Envs         map[string]string `yaml:"environments,omitempty"`
		Dir          *string           `yaml:"directory,omitempty"`
		DeleyMs      *uint             `yaml:"delayMs,omitempty"`
		KillAndRerun *bool             `yaml:"killAndRerun,omitempty"`
	}

	Config struct {
		Triggers map[string]*TriggerConfig `yaml:"triggers,omitempty"`
		Runners  map[string]*RunnerConfig  `yaml:"runners,omitempty"`
		RootDir  string                    `yaml:"root,omitempty"`
	}
)

func (fc *FreqConfig) UnmarshalYAML(b []byte) error {
	var dest interface{}
	if err := yaml.Unmarshal(b, &dest); err != nil {
		return err
	} else if dest == nil {
		return errors.New("null error")
	}

	switch v := dest.(type) {
	case int:
		*fc = FreqConfig{
			MaxWaitMs: v,
			MinWaitMs: v,
		}
		return nil
	default:
		return errors.New("invalid type error: (detected: " + reflect.TypeOf(dest).Name() + ")")
	}
}

func (c *Config) SetDefault(path string) error {

	if !filepath.IsAbs(path) {
		return errors.New("path isn't absolute path")
	}
	if !filepath.IsAbs(c.RootDir) {
		c.RootDir = filepath.Join(c.RootDir)
	}

	for id := range c.Runners {
		if err := c.Runners[id].SetDefault(*c); err != nil {
			return err
		}
	}

	for id := range c.Runners {
		if err := c.Triggers[id].SetDefault(*c); err != nil {
			return err
		}
	}

	return nil
}

func (rc *RunnerConfig) SetDefault(config Config) error {

	if len(rc.Cmds) == 0 {
		return errors.New("command is empty")
	}

	if rc.Dir == nil {
		rc.Dir = &config.RootDir
	} else if !filepath.IsAbs(*rc.Dir) {
		dir := filepath.Join(config.RootDir, *rc.Dir)
		rc.Dir = &dir
	}

	if rc.KillAndRerun == nil {
		killAndRerun := false
		rc.KillAndRerun = &killAndRerun
	}

	if rc.DeleyMs == nil {
		deleyMs := uint(300)
		rc.DeleyMs = &deleyMs
	}

	return nil
}

func (tc *TriggerConfig) SetDefault(config Config) error {

	if len(tc.IncludeExts) == 0 {
		return errors.New("extension is empty")
	}

	if len(tc.Runners) == 0 {
		return errors.New("runner is empty")
	}

	for i := range tc.ExcludeDirs {
		if !filepath.IsAbs(tc.ExcludeDirs[i]) {
			tc.ExcludeDirs[i] = filepath.Join(config.RootDir, tc.ExcludeDirs[i])
		}
	}

	for i := range tc.Runners {
		if _, exist := config.Runners[tc.Runners[i]]; !exist {
			return errors.New("unknown runner: " + tc.Runners[i])
		}
	}

	return nil
}
