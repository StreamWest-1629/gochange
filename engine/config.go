package model

import (
	"errors"
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
		Triggers map[string]TriggerConfig `yaml:"triggers,omitempty"`
		Runners  map[string]RunnerConfig  `yaml:"runners,omitempty"`
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
		fc.MaxWaitMs, fc.MinWaitMs = v, v
		return nil
	default:
		return errors.New("invalid type error: (detected: " + reflect.TypeOf(dest).Name() + ")")
	}
}
