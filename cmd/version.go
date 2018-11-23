package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Linkedin API server",
	Long:  `This command shows the version no of the Linkedi api server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.3")
	},
}