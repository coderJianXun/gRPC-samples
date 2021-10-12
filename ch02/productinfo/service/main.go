package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "productinfo/service/ecommerce"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() // 调用 gRPC API 创建新的 gRPC 服务器实例
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil { // 在指定端口上监听传入的消息
		log.Fatalf("failed to serve: %v", err)
	}
}
