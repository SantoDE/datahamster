package rpc

import (
	"github.com/SantoDE/datahamster"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/connect"
	"google.golang.org/grpc"
	"net"
)

//Server struct to hold RPC Server Information
type Server struct {
	services *Services
}

//Services struct to hold RPC Services Information
type Services struct {
	AgentService *connect.AgentService
}

//NewServer function to create a new RPC Server
func NewServer(services *datahamster.Services) *Server {
	server := new(Server)
	server.services = new(Services)
	server.services.AgentService = connect.NewAgentService(services.AgentService)
	return server
}

//Start function to create a new RPC Server
func (r *Server) Start() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	connect.RegisterAgentConnectServer(server, r.services.AgentService)
	log.Debugf("Start Serve FileUpload RPC")
	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}