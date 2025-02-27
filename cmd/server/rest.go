package server

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/server/rest"
	"github.com/spf13/cobra"
)

func NewRestServer() *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run Rest HTTP Server",
		Long:  "Run Rest HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			container := container.Setup()
			rest.StartRestHttpService(container)
		},
	}
}
