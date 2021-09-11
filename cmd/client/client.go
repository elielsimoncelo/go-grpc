package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/sousaeliel/go-fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	fmt.Println("\n### Unary RPCs")
	AddUser(client)

	fmt.Println("\n### Server streaming RPCs")
	AddUserVerbose(client)

	fmt.Println("\n### Client streaming RPCs")
	AddUsers(client)

	fmt.Println("\n### Bidirectional streaming RPCs")
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Printf("Added User: %v \n", res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}

		fmt.Printf("AddUser Stream => Status: %v | Data: %v \n", stream.GetStatus(), stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{Id: "w1", Name: "Jose", Email: "jose@jose.com"},
		{Id: "w2", Name: "Maria", Email: "m@m.com"},
		{Id: "w3", Name: "Luiza", Email: "l@l.com"},
		{Id: "w4", Name: "Fernanda", Email: "f@f.com"},
		{Id: "w5", Name: "Gabriel", Email: "g@g.com"},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println("Server response:", res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		{Id: "w1", Name: "Jose", Email: "jose@jose.com"},
		{Id: "w2", Name: "Maria", Email: "m@m.com"},
		{Id: "w3", Name: "Luiza", Email: "l@l.com"},
		{Id: "w4", Name: "Fernanda", Email: "f@f.com"},
		{Id: "w5", Name: "Gabriel", Email: "g@g.com"},
	}

	//ctx := stream.Context()
	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Printf("Sending user: %v \n", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}

			fmt.Printf("Receiving user %v with status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}

		close(wait)
	}()

	<-wait
}
