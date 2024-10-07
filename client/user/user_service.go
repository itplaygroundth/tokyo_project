package user

import (
	//"context"
    //"crypto/sha256"
    //"database/sql"
    //"encoding/hex"
    // "fmt"
    // "log"
    // "net"

    // "github.com/go-sql-driver/mysql"
    // "google.golang.org/grpc"
    // "golang.org/x/crypto/bcrypt"
   // pb "tokyo/proto"
)

// protoc --go_out=../proto --go-grpc_out=../proto  --go_opt=Muser.proto=./ --go-grpc_opt=Muser.proto=./ user.proto

// type userServiceServer struct {
//     pb.UnimplementeduserServiceServer
// }

// func (s *userServiceServer) RegisterUser(ctx context.Context, req *wallet.RegisterUserRequest) (*wallet.RegisterUserResponse, error) {
//     // Logic สำหรับการลงทะเบียนผู้ใช้
//     walletID, err := RegisterUser(req.Username, req.Password)
//     if err != nil {
//         return &wallet.RegisterUserResponse{Message: "Registration failed"}, err
//     }
//     return &wallet.RegisterUserResponse{WalletId: walletID, Message: "Registration successful"}, nil
// }

// func (s *userServiceServer) LoginUser(ctx context.Context, req *wallet.LoginUserRequest) (*wallet.LoginUserResponse, error) {
//     // Logic สำหรับการล็อกอิน
//     walletID, err := LoginUser(req.Username, req.Password)
//     if err != nil {
//         return &wallet.LoginUserResponse{Message: "Login failed"}, err
//     }
//     return &wallet.LoginUserResponse{WalletId: walletID, Message: "Login successful"}, nil
// }

 
