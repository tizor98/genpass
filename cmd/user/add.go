/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user, and by default set it as the current active user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.PrintErrln("Error: Optionally you can enter only the new username.")
			os.Exit(1)
		}

		var username string
		if len(args) == 1 {
			username = args[0]
		} else {
			cmd.Print("Enter the new username: ")
			name, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				cmd.PrintErr("An unexpected error happened.")
				os.Exit(1)
			}
			username = name
		}

		if len(username) >= 20 {
			cmd.PrintErrln("Error: Username must be less than 20 characters.")
			os.Exit(1)
		}

		cmd.Print("Enter password for new user: ")
		bt, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			cmd.PrintErr("An unexpected error happened.")
			os.Exit(1)
		}

		pass := string(bt)
		if len(pass) >= 64 {
			cmd.PrintErr("Error: The password must be less than 64 characters.")
			os.Exit(1)
		}

		username = strings.TrimSpace(strings.Trim(username, "\n"))
		pass = strings.TrimSpace(strings.Trim(pass, "\n"))

		user := service.NewUser(username, pass)

		cmd.Print("\n\nUser created successfully!!")

		if noActivate != nil && !*noActivate {
			service.SetActive(user.Username)
			cmd.Printf("\n\n%s is now set as the current active user.\n", user.Username)
		}
	},
}

func init() {
	noActivate = addCmd.Flags().BoolP("no-activate", "n", false, "Indicate if new user is set as the current active user")
}
