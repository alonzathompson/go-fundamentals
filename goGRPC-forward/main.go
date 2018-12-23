package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/alonzathompson/go-fundamentals/goGRPC/customer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50052"
)

type server struct {
	savedCustomers []*pb.CustomerRequest
}

func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("Something went wrong with parsing incoming request")
	}

	//Must map the claims in order to get value by key
	s.savedCustomers = append(s.savedCustomers, in)

	fmt.Println(md, s.savedCustomers)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ForwardCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("Something went wrong with parsing incoming request")
	}
	fmt.Println(md)

	return &pb.CustomerResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening")
	// Creates a new gRPC server
	s := grpc.NewServer()
	/*
		--IMPORTANT The Rester and Server words are added to our customer
	*/
	pb.RegisterCustomerServer(s, &server{})
	s.Serve(lis)
}
