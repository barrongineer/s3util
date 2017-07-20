// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "s3util",
	Short: "A command line utility for interacting with S3",
	Long: `Examples:
	s3util settings
	s3util put --key path/to/something/cool.exe --file ./cool.exe`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s3util.toml)")

	RootCmd.PersistentFlags().StringP("endpoint", "e", "", "endpoint to use instead of the AWS default")
	viper.BindPFlag("endpoint", RootCmd.PersistentFlags().Lookup("endpoint"))

	RootCmd.PersistentFlags().StringP("bucket", "b", "", "bucket to use for the s3 connection")
	viper.BindPFlag("bucket", RootCmd.PersistentFlags().Lookup("bucket"))

	RootCmd.PersistentFlags().StringP("id", "i", "", "id to use for authentication")
	viper.BindPFlag("id", RootCmd.PersistentFlags().Lookup("id"))

	RootCmd.PersistentFlags().StringP("secret", "s", "", "secret to use for authentication")
	viper.BindPFlag("secret", RootCmd.PersistentFlags().Lookup("secret"))

	RootCmd.PersistentFlags().StringP("region", "r", "us-east-1", "region to use for the s3 connection")
	viper.BindPFlag("region", RootCmd.PersistentFlags().Lookup("region"))
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

		// Search config in home directory with name ".s3util" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".s3util")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
