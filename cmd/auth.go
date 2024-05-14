package cmd

import (
	"os"

	"github.com/EwvwGeN/yadrive-cli/internal/constant"
	"github.com/EwvwGeN/yadrive-cli/internal/token"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().String(constant.IdFlag, "", "direct pass of the client id (by default will be get from config)")
	authCmd.Flags().String(constant.SecretFlag, "", "direct pass of the client secret (by default will be get from config)")
	viper.BindPFlag(constant.IdFlag, authCmd.Flag(constant.IdFlag))
	viper.BindPFlag(constant.SecretFlag, authCmd.Flag(constant.SecretFlag))
}


var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication for yandex disk",
	Long: `Authentication for yandex disk`,
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString(constant.IdFlag)
		if len(id) != 32 {
			_, err := cmd.ErrOrStderr().Write([]byte("client id len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		secret := viper.GetString(constant.SecretFlag)
		if len(secret) != 32 {
			_, err := cmd.OutOrStderr().Write([]byte("client secret len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		err := token.GetAccessTokenForDevice(cmd.OutOrStdout(), id, secret)
		cobra.CheckErr(err)
	},
}

