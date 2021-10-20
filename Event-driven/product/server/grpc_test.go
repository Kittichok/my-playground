package server_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/kittichok/event-driven/product/proto"

	"google.golang.org/grpc"
)

func TestGrpc(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		t.Fatalf(`Hello("") = %q, want "", error`, err)
	}
	defer conn.Close()

	client := pb.NewProductClient(conn)
	r, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "test"})
	if err != nil {
		t.Fatalf(`%q`, err)
	}
	fmt.Printf(r.GetMessage())
}
