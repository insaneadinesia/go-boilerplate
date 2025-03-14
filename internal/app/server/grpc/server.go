package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/handler/grpc/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/validator"
	"google.golang.org/grpc"
)

func StartGrpcService(container *container.Container) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", container.Config.AppGRPCPort))
	if err != nil {
		panic(err)
	}

	// inject dependency
	validator := validator.SetupValidator()
	userHandler := user.NewHandler().SetUserUsecase(container.UserUsecase).SetValidator(validator).Validate()

	// register service
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			TracingMiddleware(),
			LoggingMiddleware(container),
		),
	)

	user.RegisterUserServiceServer(s, userHandler)

	log.Printf("server listening at %v\n", lis.Addr())

	// Start server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("shutting down the server...")

	s.GracefulStop()
}
