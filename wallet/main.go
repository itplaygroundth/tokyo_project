package main

import (
    "context"
    "log"
    "net"

    pb "tokyo/proto" // เปลี่ยนเส้นทางนี้ให้ตรงกับเส้นทางจริงที่ proto ของคุณอยู่
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
    // implement your logic here
    return &pb.RegisterUserResponse{
        WalletId: "wallet123",
        Message:  "User registered successfully",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &server{})
    log.Println("gRPC server listening on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}