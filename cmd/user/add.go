/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
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

		user := new(entity.User)
		if noActivate != nil && !*noActivate {
			service.SetActive(user.Username)
			cmd.Printf("Username %s is now set as the current active user.\n", user.Username)
		}
	},
}

func init() {

	noActivate = addCmd.Flags().BoolP("no-activate", "n", false, "Indicate if new user is set as the current active user")
}
