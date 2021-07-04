package config

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

//go:embed defaults.json
var defaults []byte

var entries []ClipboardEntry

type ClipboardEntry struct {
	Name   string   `json:"name"`
	GetCmd []string `json:"get"`
	SetCmd []string `json:"set"`
}

func init() {
	if err := json.Unmarshal(defaults, &entries); err != nil {
		panic(fmt.Errorf("unmarshal: %v", err))
	}
}

func (c *ClipboardEntry) Get() ([]byte, error) {
	out, err := exec.Command(c.GetCmd[0], c.GetCmd[1:]...).Output()
	if err != nil {
		return nil, fmt.Errorf("get '%s': %v", c.Name, err)
	}

	return out, nil
}

func (c *ClipboardEntry) Set(content []byte) error {
	cmd := exec.Command(c.SetCmd[0], c.SetCmd[1:]...)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("stdin pipe '%s': %v", c.Name, err)
	}

	if _, err := pipe.Write(content); err != nil {
		return fmt.Errorf("write '%s': %v", c.Name, err)
	}

	if err := pipe.Close(); err != nil {
		return fmt.Errorf("close '%s': %v", c.Name, err)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run'%s': %v", c.Name, err)
	}

	return nil
}
