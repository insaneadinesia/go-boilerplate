package cmd

import (
	"log"
	"strings"

	"github.com/insaneadinesia/go-boilerplate/cmd/server"
	"github.com/spf13/cobra"
)

func newServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "server command handler",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("unknown command %q", strings.Join(args, " "))
			}

			_ = cmd.Help()
		},
	}

	serverCmd.AddCommand(server.NewRestServer())
	serverCmd.AddCommand(server.NewWorkerServer())

	return serverCmd
}

func init() {
	rootCmd.AddCommand(newServerCmd())
}
