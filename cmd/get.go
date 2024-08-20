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
	Long:  `Get a password for the active user.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := cmd.Context().Value(utils.GeneralUser)

		if user == nil || user.(*entity.User).Id == 0 {
			cmd.PrintErrln("Error: You must set an active user first. Type 'genpass help user' to get started.")
			os.Exit(1)
		}

		if len(args) != 1 {
			cmd.PrintErrln("Error: you must specified an entity to retrieved the password.")
			os.Exit(1)
		}

		userPass := cmd.Context().Value(utils.GeneralPassword).(string)
		forEntity := args[0]

		pass, err := service.GetPassword(forEntity, user.(*entity.User).Username, userPass)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		cmd.Println(pass)
	},
}
