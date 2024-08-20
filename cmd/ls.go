package cmd

import (
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"os"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all passwords for the active user",
	Long:  `List all passwords for the active user.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := cmd.Context().Value(utils.GeneralUser)

		if user == nil || user.(*entity.User).Id == 0 {
			cmd.PrintErrln("Error: You must set an active user first. Type 'genpass help user' to get started.")
			os.Exit(1)
		}

		if len(args) > 0 {
			cmd.PrintErrln("Error: ls command does not accept any arguments.")
			os.Exit(1)
		}

		userPass := cmd.Context().Value(utils.GeneralPassword).(string)

		passList := service.GetAllPasswords(user.(*entity.User).Username, userPass)
		if len(passList) == 0 {
			cmd.PrintErrln("Error: There are no passwords for the user.")
			os.Exit(1)
		}

		cmd.Printf("\nThere are %v passwords:\n", len(passList))
		for _, pass := range passList {
			cmd.Println(pass)
		}
	},
}
