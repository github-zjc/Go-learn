package main

import (
	"fmt"
	"net"

	pb "grpc-demo/proto/hello" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "localhost:5000"
)

// 定义helloService并实现约定的接口
type helloService struct {
	pb.UnimplementedHelloServer
}

// SayHello 实现Hello服务接口
func (h *helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func newServer() *helloService {
	return &helloService{}
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
		return
	}

	// 实例化grpc Server
	s := grpc.NewServer()
	fmt.Println("实例化grpc server成功")
	// 注册HelloService
	pb.RegisterHelloServer(s, newServer())

	// grpclog.Println("Listen on " + Address)  //grpclog.Println打印不出来
	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}
