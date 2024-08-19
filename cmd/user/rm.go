package user

import (
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/service"
	"golang.org/x/term"
	"os"
	"syscall"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete user",
	Long: `Delete a user, with all the associate passwords.

If the user is the current active user. You should indicate the new active user with 'password user [new active user]' command`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			cmd.PrintErrln("Error: You must specify a username.")
			os.Exit(1)
		}

		username := args[0]

		cmd.Print("Enter the user password: ")
		bt, err := term.ReadPassword(int(syscall.Stdin))
		cmd.Print("\n")
		if err != nil {
			cmd.PrintErrln("An unexpected error happened.")
			os.Exit(1)
		}

		pass := string(bt)

		err = service.RemoveUser(username, pass)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			os.Exit(1)
		}

		cmd.Print("\nUser remove successfully!!\n")
	},
}
