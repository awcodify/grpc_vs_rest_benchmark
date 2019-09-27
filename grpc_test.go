package benchmarks

import (
	"context"
	"testing"
	"time"

	"benchmark/grpc"
	"benchmark/grpc/proto"
	g "google.golang.org/grpc"
)

func init() {
	go grpc.Serve()
	time.Sleep(time.Second)
}

func BenchmarkGRPC(b *testing.B) {
	conn, _ := g.Dial("127.0.0.1:8889", g.WithInsecure())

	client := proto.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client proto.APIClient, b *testing.B) {
	_, err := client.Request(context.Background(), &proto.RequestInterface{
		Id:      1,
		Message: "Hello",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}
}
