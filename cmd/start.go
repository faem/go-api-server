package cmd

import (
	"LinkedinApiServer/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd2)
}

var versionCmd2 = &cobra.Command{
	Use:   "strtsrvr",
	Short: "This starts the API server",
	Long:  `This command shows the version no of the Linkedi api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer()
	},
}