package user

import (
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"golang.org/x/term"
	"os"
	"syscall"
)

var (
	deactivate *bool
)

var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Select an active user or get the current user",
	Long: `Select an active user passing the username as 'genpass user "user1"'
Or get the current user without any args as 'genpass user'".

Optionally you can also deactivate a user passing -d flag like 'genpass user -d "user1"'. This flag only works when a username is passed.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			user := cmd.Context().Value(utils.GeneralUser)

			if user == nil || user.(*entity.User).Id == 0 {
				cmd.PrintErrln("There is no active user. You can select an active user passing a username. Or first create a user if there are no users created using. Type 'genpass user help add' for more info.")
				os.Exit(0)
			}

			cmd.Printf("Active user: %s\n", user.(*entity.User).Username)
			os.Exit(0)
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

		if deactivate != nil && *deactivate {
			err = service.SetNonActive(username, pass)
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("User is set to non active state: %s\n", username)
		} else {
			err = service.SetActive(username, pass)
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("Active user: %s\n", username)
		}
	},
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(lsCmd)
	Cmd.AddCommand(rmCmd)

	deactivate = Cmd.Flags().BoolP("deactivate", "d", false, "Indicate if username must be deactivated")
}
