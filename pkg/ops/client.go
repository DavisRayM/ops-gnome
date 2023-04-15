package ops

import (
	"context"

	"github.com/DavisRayM/integration-helper/pkg/config"
	"github.com/DavisRayM/integration-helper/pkg/helm"
	pb "github.com/DavisRayM/integration-helper/proto"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.OpsClient
	conn   *grpc.ClientConn
}

func NewClient(config *config.OpsConfig) (*Client, error) {
	conn, err := grpc.Dial(config.Server.StringRep(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: pb.NewOpsClient(conn),
		conn:   conn,
	}, nil
}

func (c *Client) ListSupportedDeployment(ctx context.Context) ([]string, error) {
	req := &pb.ListSupportedDeploymentsReq{}
	resp, err := c.client.ListSupportedDeployments(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Deployment, nil
}

func (c *Client) GetDeploymentStatus(ctx context.Context, name, namespace string) (*helm.ReleaseStatus, error) {
	req := &pb.GetDeploymentStatusReq{Release: name, Namespace: namespace}
	resp, err := c.client.GetDeploymentStatus(ctx, req)
	if err != nil {
		return &helm.ReleaseStatus{}, err
	}

	return &helm.ReleaseStatus{
		Name:              name,
		Namespace:         namespace,
		Status:            resp.Status,
		StatusDescription: resp.StatusDescription,
		LastDeployed:      resp.LastDeployed,
		FirstDeployed:     resp.FirstDeployed,
	}, nil
}
