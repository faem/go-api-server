package cmd

import (
	"LinkedinApiServer/api"
	"github.com/spf13/cobra"
)

var port string //set port
var	bpa bool //bypass authentication
var stopTime int8 //stop the server after a definite time
func init() {
	startCmd.PersistentFlags().StringVarP(&port, "port", "p","8080", "This flag sets the port of our API server")
	startCmd.PersistentFlags().BoolVarP(&bpa, "bpl", "b", false, "This flag allows to bypass the authentication ")
	startCmd.PersistentFlags().Int8VarP(&stopTime,"shutdown","s",0,"This will be used for stopping the server after a definite time.")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This starts the API server",
	Long:  `This command starts the Linkedin api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.SetValues(port, bpa, stopTime)
		api.StartServer()
	},
}