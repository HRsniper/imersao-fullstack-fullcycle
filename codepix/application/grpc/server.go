package grpc

import (
  "fmt"
  "github.com/HRsniper/imersao-fullstack-fullcycle/application/grpc/pb"
  "github.com/HRsniper/imersao-fullstack-fullcycle/application/usecase"
  "github.com/HRsniper/imersao-fullstack-fullcycle/infrastructure/repository"
  "github.com/jinzhu/gorm"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "log"
  "net"
)

func StartGrpcServer(database *gorm.DB, port int) {
  grpcServer := grpc.NewServer()
  reflection.Register(grpcServer)

  pixRepository := repository.PixKeyRepositoryDb{Db: database}
  pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
  pixGrpcService := NewPixGrpcService(pixUseCase)
  pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

  address := fmt.Sprintf("0.0.0.0:%d", port)
  listener, err := net.Listen("tcp", address)
  if err != nil {
    log.Fatal("cannot start grpc server", err)
  }

  log.Printf("gRPC server has been started on port %d", port)
  err = grpcServer.Serve(listener)
  if err != nil {
    log.Fatal("cannot start grpc server", err)
  }
}
