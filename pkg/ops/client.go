package ops

import (
	"context"
	"fmt"

	"github.com/DavisRayM/ops-gnome/pkg/config"
	"github.com/DavisRayM/ops-gnome/pkg/helm"
	pb "github.com/DavisRayM/ops-gnome/proto/ops/v1"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.OpsServiceClient
	conn   *grpc.ClientConn
}

type CallOptions struct {
	task       *taskCallOptions
	deployment *deploymentCallOptions
}

type taskCallOptions struct {
	name string
}

type deploymentCallOptions struct {
	name      string
	namespace string
}

type CallOption func(o *CallOptions) error

func NewClient(config *config.AddressConfig) (*Client, error) {
	conn, err := grpc.Dial(config.StringRep(), grpc.WithInsecure())
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

func (c *Client) TriggerTask(ctx context.Context, options ...CallOption) (string, error) {
	opts := CallOptions{task: &taskCallOptions{}}
	for _, o := range options {
		o(&opts)
	}

	req := &pb.TriggerTaskRequest{Name: opts.task.name}
	resp, err := c.client.TriggerTask(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.TimeTaken, nil
}

func (c *Client) GetDeploymentStatus(ctx context.Context, options ...CallOption) (*helm.ReleaseStatus, error) {
	opts := CallOptions{deployment: &deploymentCallOptions{}}
	for _, o := range options {
		o(&opts)
	}

	req := &pb.GetDeploymentStatusRequest{Release: opts.deployment.name, Namespace: opts.deployment.namespace}
	resp, err := c.client.GetDeploymentStatus(ctx, req)
	if err != nil {
		return &helm.ReleaseStatus{}, err
	}

	return &helm.ReleaseStatus{
		Name:              opts.deployment.name,
		Namespace:         opts.deployment.namespace,
		Status:            resp.Status,
		StatusDescription: resp.StatusDescription,
		LastDeployed:      resp.LastDeployed,
		FirstDeployed:     resp.FirstDeployed,
	}, nil
}

func WithNamespace(s string) CallOption {
	return func(o *CallOptions) error {
		if o.deployment == nil {
			return fmt.Errorf("WithNamespace can only be used on Deployment related functions.")
		}
		o.deployment.namespace = s
		return nil
	}
}

func WithRelease(s string) CallOption {
	return func(o *CallOptions) error {
		if o.deployment == nil {
			return fmt.Errorf("WithRelease can only be used on Deployment related functions.")
		}
		o.deployment.name = s
		return nil
	}
}

func WithName(s string) CallOption {
	return func(o *CallOptions) error {
		if o.task == nil {
			return fmt.Errorf("WithName can only be used on Task related functions.")
		}
		o.task.name = s
		return nil
	}
}
