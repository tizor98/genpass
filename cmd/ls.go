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
	Args:  cobra.ExactArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		EnsureAUserIsActive(cmd)
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := cmd.Context().Value(utils.GeneralUser).(*entity.User)
		userPass := cmd.Context().Value(utils.GeneralPassword).(string)

		passList := service.GetAllPasswords(user.Username, userPass)
		if len(passList) == 0 {
			cmd.Println("There are no passwords for the user.")
			os.Exit(0)
		}

		cmd.Printf("There are %v passwords:\n", len(passList))
		for _, pass := range passList {
			cmd.Println(pass)
		}
	},
}
