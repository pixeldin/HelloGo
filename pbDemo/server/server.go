package main

import (
	"HelloGo/pbDemo/customer"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const (
	port = ":20000"
)

type server struct {
	customers []*pbDemo.CustomerRequest
}

// 创建
func (s *server) CreateCustomer(ctx context.Context, request *pbDemo.CustomerRequest) (*pbDemo.CustomerResponse, error) {
	s.customers = append(s.customers, request)
	return &pbDemo.CustomerResponse{Id: request.Id, Success: true}, nil
}

// 查找
func (s *server) GetCustomers(filter *pbDemo.CustomerFilter, steam pbDemo.Customer_GetCustomersServer) error {
	for _, cut := range s.customers {
		if filter.Keyword != "" {
			if !strings.Contains(cut.Name, filter.Keyword) {
				continue
			}
		}
		//找到目标发送到流式管道中
		if err := steam.Send(cut); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listener, e := net.Listen("tcp", port)
	if e != nil {
		log.Fatalf("failed to listen: %v", e)
		return
	}
	s := grpc.NewServer()
	pbDemo.RegisterCustomerServer(s, &server{})

	e = s.Serve(listener)
	if e != nil {
		log.Fatalf("failed to serve: %v", e)
		return
	}
}
