package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "github.com/ncjain/gRPC"
)

const (
	port = ":50051"
)

type UUser struct{
	Id			int32
	Username	string
	Email		string
}

var uusers []UUser

type server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	u1 := UUser{
		Id:1,
		Username:"Kjain1",
		Email:"kjain1@gmail.com",
	}
	u2 := UUser{
		Id:2,
		Username:"Kjain2",
		Email:"kjain2@gmail.com",
	}
	uusers = append(uusers, u1, u2)
	fmt.Println("Listening on port 50051")
	
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("Received: Request")
	return &pb.HealthCheckResponse{Msg: "Hello Kapil"}, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.UsersResponse, error) {
	log.Printf("Received: Get Users Request")
	u := []*pb.User{}
	for _, user := range uusers{
		u1 := & pb.User{
			Id: user.Id,
			Username: user.Username,
			Email: user.Email,
		}
		u = append(u,u1)
	}

	return &pb.UsersResponse{
		Sucess: true,
		Error: "No Error",
		Users: u,
		}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: Get User Request")
	u := pb.User{}
	for _, user := range uusers{
		if in.GetId() == user.Id{
			u.Id = user.Id
			u.Username = user.Username
			u.Email = user.Email
			break
		}
	}
	return &pb.UserResponse{
		Sucess: true,
		Error: "No Error",
		User: &u,
		}, nil
}


func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UsersResponse, error) {
	log.Printf("Received: Create User Request")
	uuser := UUser{
		Id : in.GetId(),
		Username : in.GetUsername(),
		Email : in.GetEmail(),
	}
	uusers = append(uusers, uuser)

	u := []*pb.User{}
	for _, user := range uusers{
		u1 := & pb.User{
			Id: user.Id,
			Username: user.Username,
			Email: user.Email,
		}
		u = append(u,u1)
	}

	return &pb.UsersResponse{
		Sucess: true,
		Error: "No Error",
		Users: u,
		}, nil
}


func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.UsersResponse, error) {
	log.Printf("Received: Delete User Request")
	for index, user := range uusers{
		if in.GetId() == user.Id{
			uusers = append(uusers[:index], uusers[index+1:]...)
			break
		}
	}

	u := []*pb.User{}
	for _, user := range uusers{
		u1 := & pb.User{
			Id: user.Id,
			Username: user.Username,
			Email: user.Email,
		}
		u = append(u,u1)
	}

	return &pb.UsersResponse{
		Sucess: true,
		Error: "No Error",
		Users: u,
		}, nil
}


func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: Update User Request")
	for index, user := range uusers{
		if in.GetId() == user.Id{
			uusers = append(uusers[:index], uusers[index+1:]...)
			uuser := UUser{
				Id : in.User.Id,
				Username : in.User.Username,
				Email : in.User.Email,
			}
			uusers = append(uusers, uuser)
			break
		}
	}

	return &pb.UserResponse{
		Sucess: true,
		Error: "No Error",
		User: in.GetUser(),
		}, nil
}

/*
type server struct {
	pb.UnimplementedUserServiceServer
}
lis, err := net.Listen("tcp", port)
s := grpc.NewServer()
pb.RegisterUserServiceServer(s, &server{})
err := s.Serve(lis)
*/
