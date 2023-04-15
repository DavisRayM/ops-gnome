package helm

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

// Release is a representation of a helm release object.
type Release struct {
	Name      string `yaml:"release"`
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	ChartPath string `yaml:"chartPath"`
}

type ReleaseStatus struct {
	Name              string
	Namespace         string
	Status            string
	StatusDescription string
	FirstDeployed     string
	LastDeployed      string
}

// ActionConfig returns a configuration struct that can be used to perform
// actions using the helm package
func (r *Release) ActionConfig() (*action.Configuration, error) {
	envSettings := cli.New()
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(envSettings.RESTClientGetter(), r.Namespace, "secret", log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize helm configuration: %w", err)
	}
	return actionConfig, nil
}

// Chart loads a chart from the filesystem into a helm/pkg/chart.Chart struct
// the returns struct can be used alongside the helm.sh/helm/v3 package
func (r *Release) Chart() (*chart.Chart, error) {
	return loader.Load(r.ChartPath)
}

// Path returns a string representating the path of a releases value directory
// the format returned is `helm/releases/<cluster>/<namespace>/<release-name>
func (r *Release) Path() string {
	return fmt.Sprintf("helm/releases/%s/%s/%s", r.Cluster, r.Namespace, r.Name)
}

// String returns a string representation of the ReleaseStatus Struct
func (rs *ReleaseStatus) String() string {
	return fmt.Sprintf(
		"Name: %s\nNamespace: %s\nStatus: %s\nStatus Description: %s\nLast Deployed: %s\nFirst Deployed: %s",
		rs.Name, rs.Namespace, rs.Status, rs.StatusDescription, rs.LastDeployed, rs.FirstDeployed,
	)
}
