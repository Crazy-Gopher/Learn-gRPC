package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/ncjain/gRPC"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	/*
	
	// Healthcheck
	res, err := c.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.GetMsg())
	
	// GetUsers
	res, err := c.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("gRPC Response : %s", res.GetUsers())
	
	// GetUser
	res, err := c.GetUser(ctx, &pb.GetUserRequest{Id:3})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res)
	
	// CreateUser
	res, err := c.CreateUser(ctx, &pb.CreateUserRequest{Id:3, Username:"Kjain3", Email:"kjain3@gmail.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res)

	// DeleteUser
	res, err := c.DeleteUser(ctx, &pb.DeleteUserRequest{Id:1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res)


	*/
	
	// UpdateUser
	res, err := c.UpdateUser(ctx, &pb.UpdateUserRequest{Id:1, User: &pb.User{Id:1, Username:"Kjain", Email:"kjain@gmail.com"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res)

	res1, err1 := c.GetUsers(ctx, &pb.GetUsersRequest{})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err1)
	}
	log.Printf("gRPC Response : %s", res1.GetUsers())
}

/*
conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
defer conn.Close()
c := pb.NewUserServiceClient(conn)
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
res, err := c.HealthCheck(ctx, &pb.HealthCheckRequest{})
*/