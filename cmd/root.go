package cmd

import (
	"fmt"
	"github.com/Lysander233/duck/logic"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	schema  string
	name    string
	mysql   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "duck",
	Short: "sql to struct with gorm tag",
	Long:  `Specify the table name, generate struct according to the field`,
	Run: func(cmd *cobra.Command, args []string) {
		model := logic.GenStruct(schema, name, mysql)
		fmt.Println(model)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./duck_config.toml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().StringVarP(&schema, "schema", "s", "world", "table_schema")
	rootCmd.Flags().StringVarP(&name, "name", "n", "country", "table_name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "duck_config.toml"
		viper.AddConfigPath(home)
		viper.SetConfigName("duck_config")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		mysql = viper.GetString("mysql")
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println()
	}
}
