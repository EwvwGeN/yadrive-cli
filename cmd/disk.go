package cmd

import (
	"context"
	"os"

	"github.com/EwvwGeN/yadrive-cli/internal/constant"
	"github.com/EwvwGeN/yadrive-cli/internal/util"

	"github.com/EwvwGeN/yadrive-cli/internal/disk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(diskCmd)
	diskCmd.AddCommand(diskInfoCmd)
	diskCmd.PersistentFlags().String(constant.OauthFlag, "", "direct pass of the oauth token (by default will be get from config)")
	viper.BindPFlag(constant.OauthFlag, diskCmd.Flag(constant.OauthFlag))
}

var diskCmd = &cobra.Command{
	Use:   "disk [command]",
	Short: "root command for disk operations",
	Long: `root command for disk operations`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		oauthToken := viper.GetString(constant.OauthFlag)
		if len(oauthToken) == 0 {
			_, err := cmd.ErrOrStderr().Write([]byte("oauth token len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		ctx := cmd.Context()
		ctx = context.WithValue(ctx, util.ContextKey(constant.OauthFlag), oauthToken)
		cmd.SetContext(ctx)
	},
}

var diskInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "get info about current disk",
	Long: `get info about current disk`,
	Run: func(cmd *cobra.Command, args []string) {
		oauthToken, ok := cmd.Context().Value(util.ContextKey(constant.OauthFlag)).(string)
		if !ok {
			cmd.ErrOrStderr().Write([]byte("Something went wrong"))
			os.Exit(1)
		}
		err := disk.GetDiskInfo(cmd.OutOrStdout(), oauthToken)
		cobra.CheckErr(err)
	},
	
}
