package config

import (
	"fmt"
	"io/ioutil"

	"github.com/DavisRayM/ops-gnome/pkg/helm"
	"gopkg.in/yaml.v3"
)

type OpsConfig struct {
	Server AddressConfig `yaml:"ops"`

	SupportedDeployments Deployments `yaml:"deployments"`
	SupportedTasks       []Task      `yaml:"tasks"`
}

type Deployments []helm.Release

type Task struct {
	Command []string `yaml:"command"`
	Name    string   `yaml:"name"`
	Repo    TaskRepo `yaml:"repo,omitempty"`
}

type TaskRepo struct {
	URL       string `yaml:"url,omitempty"`
	DeployKey string `yaml:"key,omitempty"`
	Branch    string `yaml:"branch,omitempty"`
}

type AddressConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func (a *AddressConfig) StringRep() string {
	return fmt.Sprintf("%s:%d", a.Address, a.Port)
}

func LoadOpsConfig(path string) (*OpsConfig, error) {
	var c OpsConfig

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading ops configuration: %w", err)
	}

	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	return &c, nil
}
