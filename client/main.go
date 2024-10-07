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
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := wallet.NewWalletServiceClient(conn)

    // ตัวอย่างเรียกใช้ RegisterUser
    res, err := c.RegisterUser(context.Background(), &wallet.RegisterUserRequest{
        Username: "testuser",
        Password: "testpassword",
    })
    if err != nil {
        log.Fatalf("could not register user: %v", err)
    }
    log.Printf("RegisterUser Response: %s", res.Message)
}