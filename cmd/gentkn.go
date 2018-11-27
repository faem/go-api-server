package cmd

import (
	"LinkedinApiServer/api"
	"fmt"
	"github.com/spf13/cobra"
)
var user string
var exp int

func init() {
	tokenCmd.PersistentFlags().StringVarP(&user, "user","u", "admin", "This flag sets the username")
	tokenCmd.PersistentFlags().IntVarP(&exp, "exp","e", 10, "This flag sets the expiration time of the token, default is 10 minute")
	rootCmd.AddCommand(tokenCmd)
}

var tokenCmd = &cobra.Command{
	Use:   "gentkn",
	Short: "Generates a token",
	Long:  `This generates a token for the Linkedin api server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.GetTokenCmd(user,exp))
	},
}