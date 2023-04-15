package ops

import (
	"net"

	"github.com/DavisRayM/integration-helper/internal/config"
	pb "github.com/DavisRayM/integration-helper/proto"
	"google.golang.org/grpc"
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
