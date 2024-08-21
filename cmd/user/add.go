package user

import (
	"bufio"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/service"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

var (
	noActivate *bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user, and by default set it as the current active user.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 1 {
			username = args[0]
		} else {
			cmd.Print("Enter the new username: ")
			name, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				cmd.PrintErrln("An unexpected error happened.")
				os.Exit(1)
			}
			username = strings.TrimSpace(name)
		}

		if len(username) >= 20 {
			cmd.PrintErrln("Error: Username must be less than 20 characters.")
			os.Exit(1)
		}

		cmd.Print("Enter password for new user: ")
		bt, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			cmd.PrintErrln("An unexpected error happened.")
			os.Exit(1)
		}

		pass := string(bt)
		if len(pass) >= 64 {
			cmd.PrintErrln("Error: The password must be less than 64 characters.")
			os.Exit(1)
		}
		pass = strings.TrimSpace(pass)

		cmd.Printf("\n")

		user, err := service.NewUser(username, pass)
		if err != nil {
			cmd.PrintErrf("Error: %s.\n", err.Error())
			os.Exit(1)
		}

		cmd.Print("\nUser created successfully!!\n")

		if noActivate != nil && !*noActivate {
			err := service.SetActive(user.Username, pass)
			if err != nil {
				cmd.PrintErrf("Error: %s.\n", err.Error())
				os.Exit(1)
			}
			cmd.Printf("\n%s is now set as the current active user.\n", user.Username)
		}
	},
}

func init() {
	noActivate = addCmd.Flags().BoolP("no-activate", "n", false, "Indicate if new user is set as the current active user")
}
