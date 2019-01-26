// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
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
	"log"
)

// testConfigCmd represents the testConfig command
var testConfigCmd = &cobra.Command{
	Use:   "testConfig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testConfig called")
		viper.SetConfigFile("./configs/env.json")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
		port := viper.Get("prod.port")
		fmt.Printf("Port: %v, Type: %T\n", port, port)
		host := viper.Get("prod.host") // returns string
		fmt.Printf("Host: %v, Type: %T\n", host, host)
	},
}

func init() {
	rootCmd.AddCommand(testConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
