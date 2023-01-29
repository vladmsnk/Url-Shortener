package grpcserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
	"vladmsnk/taskrec/config"
)

type GRPCConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type GRPCServer struct {
	Ser  *grpc.Server
	Addr string
	lis  net.Listener
}

func NewGRPCServer(cfg *config.GRPc, opts ...grpc.ServerOption) (*GRPCServer, error) {

	comp := &GRPCServer{
		Ser:  grpc.NewServer(opts...),
		Addr: fmt.Sprintf(":%s", cfg.Port),
	}
	var err error
	if comp.lis, err = net.Listen("tcp", comp.Addr); err != nil {
		return nil, err
	}
	return comp, nil
}

func ResponseHeaderMatcher(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	headers := w.Header()
	if location, ok := headers["Grpc-Metadata-Location"]; ok {
		w.Header().Set("Location", location[0])
		w.WriteHeader(http.StatusFound)
	}

	return nil
}

func (ser *GRPCServer) Run() {
	go func() {
		_ = ser.Ser.Serve(ser.lis)
	}()
}

func (ser *GRPCServer) Close() {
	if ser.lis == nil {
		return
	}
	ser.Ser.GracefulStop()
}
