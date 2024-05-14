package cmd

import (
	"fmt"

	"github.com/EwvwGeN/yadrive-cli/internal/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication for yandex disk",
	Long: `Authentication for yandex disk`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().String(constant.IdFlag, "", "direct pass of the client id (by default will be get from config)")
	authCmd.Flags().String(constant.SecretFlag, "", "direct pass of the client secret (by default will be get from config)")
	viper.BindPFlag(constant.IdFlag, authCmd.Flag(constant.IdFlag))
	viper.BindPFlag(constant.SecretFlag, authCmd.Flag(constant.SecretFlag))
}
