package main

import (
	"context"
	pb "github.com/xautjzd/grpc-template/structure"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPersonServiceClient(conn)
	person := &pb.PersonCreateRequest{
		Name: "jzd",
		Age:  18,
		Sex:  true,
	}
	r, err := c.AddPerson(context.Background(), person)
	if err != nil {
		log.Fatal("err: %v", err)
	}
	log.Printf("response: %v", r)
}
