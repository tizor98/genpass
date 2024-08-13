package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tizor98/genpass/service"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new password",
	Long: `Generate a new password based on the options provided. 
Default options is a password with at least one capital letter, lower letter and at least one number:
`,
	Run: func(cmd *cobra.Command, args []string) {
		pass := service.NewPassword(cmd.Context(), service.PassTypeAll)
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
	newCmd.Flags().StringP("for", "f", "", "Indicate the entity for witch a password will be generated")
}
