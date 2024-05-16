package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/EwvwGeN/yadrive-cli/internal/constant"
	"github.com/EwvwGeN/yadrive-cli/internal/resource"
	"github.com/EwvwGeN/yadrive-cli/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(resourceCmd)
	resourceCmd.AddCommand(resourceDownloadCmd)
	resourceCmd.PersistentFlags().String(constant.OauthFlag, "", "direct pass of the oauth token (by default will be get from config)")
	resourceCmd.PersistentFlags().String(constant.PathFlag, "", "path to the resource in the disk")
	viper.BindPFlag(constant.OauthFlag, resourceCmd.Flag(constant.OauthFlag))

	resourceDownloadCmd.Flags().String(constant.PathToFlag, "./", "destination path")
	resourceDownloadCmd.MarkFlagRequired(constant.PathFlag)
}

var resourceCmd = &cobra.Command{
	Use:   "resource [command]",
	Short: "root command for resource operations",
	Long: `root command for resource operations`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		oauthToken := viper.GetString(constant.OauthFlag)
		if len(oauthToken) == 0 {
			_, err := cmd.ErrOrStderr().Write([]byte("oauth token len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		var path *string
		if cmd.Flag(constant.PathFlag).Changed {
			parsedPath := cmd.Flag(constant.PathFlag).Value.String()
			path = &parsedPath
		}
		if path != nil && *path == "" {
			_, err := cmd.OutOrStderr().Write([]byte("path len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		ctx := cmd.Context()
		ctx = context.WithValue(ctx, util.ContextKey(constant.OauthFlag), oauthToken)
		ctx = context.WithValue(ctx, util.ContextKey(constant.PathFlag), path)
		cmd.SetContext(ctx)
	},
}

var resourceDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download resource by path",
	Long: `download resource by path`,
	Run: func(cmd *cobra.Command, args []string) {
		oauthToken, ok := cmd.Context().Value(util.ContextKey(constant.OauthFlag)).(string)
		if !ok {
			cmd.ErrOrStderr().Write([]byte("Something went wrong"))
			os.Exit(1)
		}
		path, ok := cmd.Context().Value(util.ContextKey(constant.PathFlag)).(*string)
		if !ok {
			cmd.ErrOrStderr().Write([]byte("Something went wrong"))
			os.Exit(1)
		}
		// should check path is valid?
		pathTo := cmd.Flag(constant.PathToFlag).Value.String()
		if len(strings.Split(pathTo, "/")) < 2 {
			_, err := cmd.OutOrStderr().Write([]byte("destination path len is not valid"))
			cobra.CheckErr(err)
			os.Exit(1)
		}
		err := resource.DownloadFileByPath(cmd.OutOrStdout(), oauthToken, *path, pathTo)
		cobra.CheckErr(err)
	},
	
}
