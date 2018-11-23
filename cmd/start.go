package cmd

import (
	"LinkedinApiServer/api"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This starts the API server",
	Long:  `This command starts the Linkedin api server`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)
		for _,val := range args{
			switch val {
			case "bypass":
				fmt.Println("bypass login credential")
			default:
				fmt.Println("no such argument for start command")
			}
		}
		api.StartServer()
	},
}