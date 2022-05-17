module github.com/zhanchangbao/gohbase

go 1.13

require (
	github.com/go-zookeeper/zk v1.0.2
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.6.0
	go.opentelemetry.io/otel/170 v1.7.0
	go.opentelemetry.io/otel/trace/170 v1.7.0
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
	modernc.org/b v1.0.0
	modernc.org/mathutil v1.1.1 // indirect
	modernc.org/strutil v1.1.0 // indirect
)

replace (
go.opentelemetry.io/otel/170 v1.7.0 => go.opentelemetry.io/otel v1.7.0
go.opentelemetry.io/otel/trace/170 v1.7.0 => go.opentelemetry.io/otel/trace v1.7.0
)