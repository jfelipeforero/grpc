package grpc

import (
 "fmt"
 "net"
 "log"
 "github.com/jfelipeforero/grpc-proto/golang/order"
 "github.com/jfelipeforero/grpc/order/config"
 "github.com/jfelipeforero/grpc/order/internal/ports"
 "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 "google.golang.org/grpc/reflection"	
 "google.golang.org/grpc"
)

type Adapter struct {
  api ports.APIPort
  port int
  order.UnimplementedOrderServer
} 

func NewAdapter(api ports.APIPort, port int) {
  return &Adapter{api: api, port: port}
}

func (a Adapter) Run(){
  var err error

  listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
  if err != nil {
    log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
  }

  grpcServer := grpc.NewServer()
  order.RegisterOrderServer(grpcServer, a) 
  if config.GetEnv() ==  "development" {
      reflection.Register(grpcServer)
  }

  if err := grpcServer.Serve(listen); err != nil {
      log.Fatalf("failed to serve grpc on port")
  }
} 
