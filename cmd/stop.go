package cmd

import (
	"LinkedinApiServer/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "This starts the API server",
	Long:  `This command starts the Linkedin api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Stop <- false
	},
}