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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Output the global settings from the config file and flags",
	Long: `Example:
	s3util settings`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("endpoint: %s\n", viper.GetString("endpoint"))
		fmt.Printf("bucket: %s\n", viper.GetString("bucket"))
		fmt.Printf("region: %s\n", viper.GetString("region"))
		fmt.Printf("id: %s\n", viper.GetString("id"))
		fmt.Printf("secret: %s\n", viper.GetString("secret"))
	},
}

func init() {
	RootCmd.AddCommand(settingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// settingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// settingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
