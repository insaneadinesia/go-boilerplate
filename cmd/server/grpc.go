package server

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/server/grpc"
	"github.com/spf13/cobra"
)

func NewGrpcServer() *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "Run GRPC Server",
		Long:  "Run GRPC Server",
		Run: func(cmd *cobra.Command, args []string) {
			container := container.Setup()
			grpc.StartGrpcService(container)
		},
	}
}
