/*
Package user
*/
package user

import (
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/service"
)

const (
	headerUsername = "Username"
	headerIsActive = "Is active"
	isActiveText   = "Yes"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all users.",
	Long: `List all users registered with the information of who (if any) is the active user

If you don´t have any user registered use 'genpass user add "username"' to registered one
or 'genpass user "username"' to setup an active user for which password will be saved.`,
	Run: func(cmd *cobra.Command, args []string) {
		users := service.GetUsers()

		tb := table.New(headerUsername, headerIsActive)
		tb.WithHeaderSeparatorRow('-')

		for username, isActive := range users {
			var textIsActive string
			if isActive {
				textIsActive = isActiveText
			}
			tb.AddRow(username, textIsActive)
		}
		tb.Print()
	},
}
