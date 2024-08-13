package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/service"
	"github.com/tizor98/genpass/utils"
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

Example: genpass new -t=n -l=30 www.google.com - Will generate a 30 length password containing capital letters, lower letters and numbers.
And in case there is a user setup, will save the generated password for www.google.com entity.

`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		if len(args) > 0 {
			if nil == ctx.Value(utils.GeneralUser) {
				fmt.Println("ERROR: If you specified an entity, you must setup first a user. Try 'genpass help user' for more info")
				return
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
		fmt.Println(pass)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	passType = newCmd.Flags().StringP("type", "t", "a", "Indicate the password type to generate. Options: a=All, s=Cap and lower letters, n=Same as s but with numbers.")
	passLength = newCmd.Flags().IntP("length", "l", 20, "Indicate the length of the password")
}
