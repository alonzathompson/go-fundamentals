package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/dgrijalva/jwt-go"

	auth "github.com/alonzathompson/go-fundamentals/goGRPC/authTokenHandler"
	pb "github.com/alonzathompson/go-fundamentals/goGRPC/customer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port      = ":50051"
	connector = "localhost:50052"
)

// server is used to implement customer.CustomerServer.
/*
	Our server struct contains a slice of Pointer to Customer request
*/
type server struct {
	savedCustomers []*pb.CustomerRequest
}

var msg *pb.CustomerRequest

// CreateCustomer creates a new Customer
/*
Notice that we start from a struct and then attatch
reciever functions to it. Notice that the first param
is the context object. Second param is a Pointer to our
customer object in proto. It retruns a customer id and bool
*/
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("Something went wrong with parsing incoming request")
	}

	//Parsed the jwt
	t, err := auth.ParseJWT(md)
	if err != nil {
		fmt.Println(err)
	}

	//Must map the claims in order to get value by key
	claims := t.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	ctx = metadata.AppendToOutgoingContext(context.Background(), "user", claims["usr"].(string))
	s.savedCustomers = append(s.savedCustomers, in)
	testRun(ctx, in)

	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *server) ForwardCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("Something went wrong with parsing incoming request")
	}
	fmt.Println(in, md)

	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

//ForwardCustomer request to another server
func forwardCustomer(ctx context.Context, client pb.CustomerClient, in *pb.CustomerRequest) {
	// Creates a new CustomerClient
	fmt.Println("ForwardCustomer Forwarding...")
	//End of important part
	resp, err := client.ForwardCustomer(context.Background(), in)
	if err != nil {
		log.Fatalf("Could not Forward: %v", err)
	}
	if resp.Success {
		log.Printf("Forwarded new request")
	}
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

func testRun(c context.Context, m *pb.CustomerRequest) {
	fmt.Println("Connecting from testRun")
	conn, err := grpc.Dial(connector, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)
	//End of important part

	//Declaring customer at the addreess of customer request object
	if m != nil {
		customer := m
		forwardCustomer(c, client, customer)
	} else {
		customer := &pb.CustomerRequest{
			Id:    101,
			Name:  "Shiju Varghese",
			Email: "shiju@xyz.com",
			Phone: "732-757-2923",
			Addresses: []*pb.CustomerRequest_Address{
				&pb.CustomerRequest_Address{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: false,
				},
				&pb.CustomerRequest_Address{
					Street:            "Greenfield",
					City:              "Kochi",
					State:             "KL",
					Zip:               "68356",
					IsShippingAddress: true,
				},
			},
		}
		forwardCustomer(c, client, customer)
	}

	// Create a new customer
	//createCustomer(client, customer)
	// Filter with an empty Keyword
	//filter := &pb.CustomerFilter{Keyword: ""}

}

func runSrv() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	/*
		--IMPORTANT The Rester and Server words are added to our customer
	*/
	pb.RegisterCustomerServer(s, &server{})
	s.Serve(lis)
}

func main() {
	runSrv()

	//close(msgCh)
}
