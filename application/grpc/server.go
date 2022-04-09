package grpc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf( format: "0.0.0.0:#{port}" )
	listener, err := net.Listen( network: "tcp", address )
	if err != nil {
		log.Fatal( v...: "cannot start grpc server", err )
	} 

	log.Printf( format: "gRPC server has been started on port %d", port )
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal( v...: "cannot start grpc server", err )
	}
}