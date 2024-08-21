package cmd

import (
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"os"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a password for the active user",
	Long:  `Remove a password for the active user. Example: 'genpass rm www.google.com' will remove the password saved for google if exists.`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		EnsureAUserIsActive(cmd)
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := cmd.Context().Value(utils.GeneralUser).(*entity.User)
		userPass := cmd.Context().Value(utils.GeneralPassword).(string)
		forEntity := args[0]

		err := service.DeletePassword(forEntity, user.Username, userPass)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		cmd.Printf("Password for %s deleted successfully.\n", forEntity)
	},
}
