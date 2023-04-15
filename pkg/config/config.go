package config

import (
	"fmt"
	"io/ioutil"

	"github.com/DavisRayM/integration-helper/pkg/helm"
	"gopkg.in/yaml.v3"
)

type OpsConfig struct {
	Server AddressConfig `yaml:"ops"`

	SupportedDeployments Deployments `yaml:"deployments"`
}

type Deployments []helm.Release

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
