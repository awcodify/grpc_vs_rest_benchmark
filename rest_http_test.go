package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"benchmark/rest"
)

func init() {
	go rest.Serve()
	time.Sleep(time.Second)
}

func BenchmarkREST(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doREST(client, b)
	}
}

func doREST(client *http.Client, b *testing.B) {
	request := &rest.Request{
		Id:      1,
		Message: "Hello",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(request)

	resp, err := client.Post("http://127.0.0.1:8888/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()
}
