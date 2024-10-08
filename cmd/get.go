package cmd

import (
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password for the active user",
	Long:  `Get a password for the active user. Example 'genpass get www.google.com'`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		EnsureAUserIsActive(cmd)
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := cmd.Context().Value(utils.GeneralUser).(*entity.User)
		userPass := cmd.Context().Value(utils.GeneralPassword).(string)
		forEntity := args[0]

		pass, err := service.GetPassword(forEntity, user.Username, userPass)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		cmd.Println(pass)
	},
}
