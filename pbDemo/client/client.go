package main

import (
	"HelloGo/pbDemo/customer"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address = "localhost:20000"
)

func createCustomer(client pbDemo.CustomerClient, request *pbDemo.CustomerRequest) {
	response, e := client.CreateCustomer(context.Background(), request)
	if e != nil {
		log.Fatalf("Create customer err: %v", e)
		return
	}
	if response.Success {
		log.Printf("A new customer created success with id: %d", response.Id)
	}
}

func getCustomers(client pbDemo.CustomerClient, filter *pbDemo.CustomerFilter) {
	stream, e := client.GetCustomers(context.Background(), filter)
	if e != nil {
		log.Fatalf("Get customer err: %v", e)
		return
	}
	for {
		cus, e := stream.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("%v.GetCustomers() err: %v", client, e)
		}
		log.Printf("Customer: %v", cus)
	}
}

func main() {
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Connect to grpc failed: %v", e)
	}
	defer conn.Close()
	client := pbDemo.NewCustomerClient(conn)
	//cus1 := genCustomer(1, "pixel", "pixel@qq.com", "123456")
	//cus2 := genCustomer(2, "pig", "pig@qq.com", "666666")

	//createCustomer(client, cus1)
	//createCustomer(client, cus2)
	filter := &pbDemo.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}

func genCustomer(id int32, name string, email string, phone string) *pbDemo.CustomerRequest {
	return &pbDemo.CustomerRequest{
		Id:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
