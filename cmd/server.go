package cmd

import (
	"echo_shop/bootstrap"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run app server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
