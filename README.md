Result in mine
```
~/Documents/playground/grpc-rest-benchmark(master ✔) go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: benchmark
BenchmarkGRPC-4   	   20000	     89584 ns/op	    9256 B/op	     173 allocs/op
BenchmarkREST-4   	   10000	    215445 ns/op	   19023 B/op	     143 allocs/op
PASS
ok  	benchmark	7.002s
~/Documents/playground/grpc-rest-benchmark(master ✔) go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: benchmark
BenchmarkGRPC-4   	   20000	     90188 ns/op	    9255 B/op	     173 allocs/op
BenchmarkREST-4   	    3000	    381953 ns/op	   19037 B/op	     143 allocs/op
PASS
ok  	benchmark	5.935s
~/Documents/playground/grpc-rest-benchmark(master ✔) go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: benchmark
BenchmarkGRPC-4   	   20000	     88909 ns/op	    9256 B/op	     173 allocs/op
BenchmarkREST-4   	    3000	    459965 ns/op	   19038 B/op	     143 allocs/op
PASS
ok  	benchmark	6.233s
```
