/*
Package cmd
Copyright © 2024 Alberto Ortiz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/cmd/user"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"golang.org/x/term"
	"os"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use:   "genpass",
	Short: "Passwords generator and management system",
	Long: `
A cli tool to generate and manage passwords for all your services:
This is a non-profit project and you can use it as you want as long as the licence allow it.

To start try using 'genpass new -t=s' or 'genpass help new' for more info.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if user, ok := service.IsAuth(); ok {
			cmd.SetContext(context.WithValue(cmd.Context(), utils.GeneralUser, user))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func AskForPassword(cmd *cobra.Command) {
	u := cmd.Context().Value(utils.GeneralUser)

	if u == nil || u.(*entity.User).Id == 0 {
		return
	}

	cmd.Print("Enter the user password: ")
	bt, err := term.ReadPassword(int(syscall.Stdin))
	cmd.Print("\n")
	if err != nil {
		cmd.PrintErrln("An unexpected error happened.")
		os.Exit(1)
	}

	pass := string(bt)

	err = service.VerifyUserPassword(pass, u.(*entity.User).Password)
	if err != nil {
		cmd.PrintErrf("Error: %v\n", err)
		os.Exit(1)
	}

	cmd.SetContext(context.WithValue(cmd.Context(), utils.GeneralPassword, pass))
}

func EnsureAUserIsActive(cmd *cobra.Command) {
	u := cmd.Context().Value(utils.GeneralUser)

	if u == nil || u.(*entity.User).Id == 0 {
		cmd.PrintErrln("Error: You must set an active user first. Type 'genpass help user' to get started.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(user.Cmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(rmCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
