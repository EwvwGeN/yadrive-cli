package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"

	"github.com/EwvwGeN/yadrive-cli/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userSecret string = "change by building"
	cfgFile string
 	saveToFile bool
)

var rootCmd = &cobra.Command{
	Use:   "yadrive-cli",
	Short: "CLI application for using yandex drive",
	Long: `CLI application for using yandex drive`,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if !saveToFile {
			return
		}
		keys := viper.AllKeys()
		if len(keys) == 0 {
			return
		}
		data := bytes.Buffer{}
		for i := 0; i < len(keys); i++ {
			data.WriteString(keys[i]+": ")
			data.WriteString(viper.GetString(keys[i]) + "\n")
		}

		block, err := aes.NewCipher([]byte(util.GetMD5Hash(userSecret)))
		cobra.CheckErr(err)

		gcmInstance, err := cipher.NewGCM(block)
		cobra.CheckErr(err)

		nonce := make([]byte, gcmInstance.NonceSize())
		io.ReadFull(rand.Reader, nonce)
		cipheredText := gcmInstance.Seal(nonce, nonce, data.Bytes(), nil)
		f, err := os.OpenFile(cfgFile, os.O_CREATE, 0666)
		cobra.CheckErr(err)
		_, err = f.Write(cipheredText)
		if err != nil {
			f.Close()
			cmd.OutOrStderr().Write([]byte("Error while saving data"))
			return
		}
		f.Close()
	  },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yadrive.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&saveToFile, "save", "s", false, "Save data to config file")
}

func initConfig() {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		cfgFile = home + string(os.PathSeparator) + ".yadrive-cli"
	}
	data, err := os.ReadFile(cfgFile)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	cobra.CheckErr(err)
	block, err := aes.NewCipher([]byte(util.GetMD5Hash(userSecret)))
	cobra.CheckErr(err)
	gcmInstance, err := cipher.NewGCM(block)
	cobra.CheckErr(err)
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	cobra.CheckErr(err)
	reader := bytes.NewReader(plainText)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(reader); err != nil {
		rootCmd.OutOrStderr().Write([]byte("Cant read config file"))
	}
}

