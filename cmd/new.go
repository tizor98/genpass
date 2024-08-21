package cmd

import (
	"bufio"
	"context"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
	"os"
	"strings"
)

var (
	passType   *string
	passLength *int
)
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new password",
	Long: `Generate a new password based on the options provided. 

If you have setup a user. You can optionally specified an entity for which the password must be generated.
The new password will be associate with the entity and the user in encrypted form for further consult.

Example: genpass new -t n -l 30 www.google.com - Will generate a 30 length password containing capital letters, lower letters and numbers.
And in case there is a user setup, will save the generated password for www.google.com entity.
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		AskForPassword(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		if len(args) > 0 {
			if nil == ctx.Value(utils.GeneralUser) {
				cmd.PrintErr("Error: If you specified an entity, you must setup first a user. Try 'genpass help user' for more info")
				os.Exit(1)
			}

			ctx = context.WithValue(ctx, utils.NewArgForEntity, args[0])
		}

		selectedPassType := service.PassTypeAll
		if passType != nil {
			switch *passType {
			case "s":
				selectedPassType = service.PassTypeCapitalAndLower
				break
			case "n":
				selectedPassType = service.PassTypeCapitalAndLowerAndNumber
				break
			}
		}
		ctx = context.WithValue(ctx, utils.NewFlagPassType, selectedPassType)

		if passLength != nil {
			ctx = context.WithValue(ctx, utils.NewFlagPassLength, *passLength)
		}

		pass := service.NewPassword(ctx)

		user := ctx.Value(utils.GeneralUser)
		forEntity := ctx.Value(utils.NewArgForEntity)

		if user == nil || forEntity == nil || len(forEntity.(string)) == 0 {
			cmd.Println(pass)
			os.Exit(0)
		}

		userPass := cmd.Context().Value(utils.GeneralPassword).(string)

		if hasOne := service.HasPassword(forEntity.(string), user.(*entity.User).Username, userPass); !hasOne {
			service.SaveNewPassword(pass, forEntity.(string), user.(*entity.User), userPass)
			cmd.Println(pass)
			os.Exit(0)
		}

		var decision string

		cmd.Printf("It already exists a password for %s. Are you sure you want to replace it? (Y/n): ", forEntity)
		decision, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			cmd.PrintErrln("An unexpected error happened.")
			os.Exit(1)
		}

		decision = strings.TrimSpace(decision)
		if decision != "" && strings.ToLower(decision) != "y" {
			os.Exit(0)
		}
		service.UpdatePassword(pass, forEntity.(string), user.(*entity.User), userPass)
		cmd.Println(pass)
	},
}

func init() {
	passType = newCmd.Flags().StringP("type", "t", "a", "Indicate the password type to generate. Options: a=All, s=Cap and lower letters, n=Same as s but with numbers.")
	passLength = newCmd.Flags().IntP("length", "l", 20, "Indicate the length of the password")
}
