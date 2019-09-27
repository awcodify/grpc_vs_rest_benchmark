package grpc

import (
	"benchmark/grpc/proto"
	"google.golang.org/grpc"

	"context"
	"log"
	"net"
)

type Server struct{}

func (s *Server) Request(ctx context.Context, in *proto.RequestInterface) (*proto.Response, error) {
	return &proto.Response{
		Code: 200,
		Data: in,
	}, nil
}

func Serve() {
	listen, err := net.Listen("tcp", ":8889")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterAPIServer(server, &Server{})
	log.Println(server.Serve(listen))
}
