package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "LinkedinApiServer",
	Short: "This is the main command",
	Long: `This is the main command of our Linkedin API server`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("demo works")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
