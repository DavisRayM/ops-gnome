package ops

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/DavisRayM/integration-helper/pkg/config"
	"github.com/DavisRayM/integration-helper/pkg/helm"
	pb "github.com/DavisRayM/integration-helper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"helm.sh/helm/v3/pkg/action"
)

// The API struct is a representation of the Ops Servers API.
type API struct {
	pb.UnimplementedOpsServer
	grpcServer *grpc.Server
	config     *config.OpsConfig
}

// New creates a new API object and registers the Ops Service on the
// GRPC Server
func New(config *config.OpsConfig) (*API, error) {
	var opts []grpc.ServerOption

	a := &API{
		config:     config,
		grpcServer: grpc.NewServer(opts...),
	}
	a.grpcServer.RegisterService(&pb.Ops_ServiceDesc, a)
	return a, nil
}

// Start binds the API server to listen on a specific port
func (a *API) Start() error {
	lis, err := net.Listen("tcp", a.config.Server.StringRep())
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(lis)
}

// Stop gracefully shutdown the API services
func (a *API) Stop() {
	a.grpcServer.GracefulStop()
}

func (a *API) ListSupportedDeployments(ctx context.Context, req *pb.ListSupportedDeploymentsReq) (*pb.ListSupportedDeploymentsResp, error) {
	var deployments []string

	for _, deployment := range a.config.SupportedDeployments {
		deployments = append(deployments, fmt.Sprintf("Release %s on %s", deployment.Name, deployment.Cluster))
	}

	return &pb.ListSupportedDeploymentsResp{Deployment: deployments}, nil
}

func (a *API) retrieveDeployment(name string, namespace string) (helm.Release, error) {
	for _, r := range a.config.SupportedDeployments {
		if r.Name == name && r.Namespace == namespace {
			return r, nil
		}
	}

	return helm.Release{}, fmt.Errorf("deployment %s is currently not supported", name)
}

func (a *API) GetDeploymentStatus(ctx context.Context, req *pb.GetDeploymentStatusReq) (*pb.GetDeploymentStatusResp, error) {
	deployment, err := a.retrieveDeployment(req.Release, req.Namespace)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	config, err := deployment.ActionConfig()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	client := action.NewStatus(config)
	client.ShowResources = false
	client.ShowDescription = true

	release, err := client.Run(deployment.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetDeploymentStatusResp{
		Status:            release.Info.Status.String(),
		StatusDescription: release.Info.Description,
		FirstDeployed:     release.Info.FirstDeployed.Format(time.UnixDate),
		LastDeployed:      release.Info.LastDeployed.Format(time.UnixDate),
	}, nil
}
