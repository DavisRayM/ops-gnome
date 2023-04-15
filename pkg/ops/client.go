package ops

import (
	"context"

	"github.com/DavisRayM/integration-helper/pkg/config"
	"github.com/DavisRayM/integration-helper/pkg/helm"
	pb "github.com/DavisRayM/integration-helper/proto/ops/v1"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.OpsServiceClient
	conn   *grpc.ClientConn
}

func NewClient(config *config.OpsConfig) (*Client, error) {
	conn, err := grpc.Dial(config.Server.StringRep(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: pb.NewOpsServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *Client) ListSupportedDeployment(ctx context.Context) ([]string, error) {
	req := &pb.ListSupportedDeploymentsRequest{}
	resp, err := c.client.ListSupportedDeployments(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Deployments, nil
}

func (c *Client) ListSupportedTasks(ctx context.Context) ([]*pb.Task, error) {
	req := &pb.ListSupportedTasksRequest{}
	resp, err := c.client.ListSupportedTasks(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Tasks, nil
}

func (c *Client) TriggerTask(ctx context.Context, name string) (string, error) {
	req := &pb.TriggerTaskRequest{Name: name}
	resp, err := c.client.TriggerTask(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.TimeTaken, nil
}

func (c *Client) GetDeploymentStatus(ctx context.Context, name, namespace string) (*helm.ReleaseStatus, error) {
	req := &pb.GetDeploymentStatusRequest{Release: name, Namespace: namespace}
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
