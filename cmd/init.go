package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate config file",
	Long:  `Generate configuration file ~/duck_config.toml`,
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig()
		fmt.Println("Generate ~/duck_config.toml success!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func writeConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName("duck_config")
	viper.SetConfigType("toml")

	/*	viper.Set("app_name", "awesome web")
		viper.Set("log_level", "DEBUG")
		viper.Set("mysql.ip", "127.0.0.1")
		viper.Set("mysql.port", 3306)
		viper.Set("mysql.user", "root")
		viper.Set("mysql.password", "123456")
		viper.Set("mysql.database", "awesome")

		viper.Set("redis.ip", "127.0.0.1")
		viper.Set("redis.port", 6381)*/

	mysql := `root:123456@(127.0.0.1:3306)/world?charset=utf8&parseTime=True&loc=Local`
	viper.Set("mysql", mysql)

	err2 := viper.SafeWriteConfig()
	if err2 != nil {
		log.Fatal("write config failed: ", err)
	}
}
