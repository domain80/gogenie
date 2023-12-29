/*
Copyright © 2023 David Mainoo work.davidmainoo@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genie",
	Short: "A cli tool to quickly generate boilerplate code for developing webapps in golang ",
	Long: `

    /$$$$$$                      /$$          
   /$$__  $$                    |__/          
  | $$  \__/  /$$$$$$  /$$$$$$$  /$$  /$$$$$$ 
  | $$ /$$$$ /$$__  $$| $$__  $$| $$ /$$__  $$
  | $$|_  $$| $$$$$$$$| $$  \ $$| $$| $$$$$$$$
  | $$  \ $$| $$_____/| $$  | $$| $$| $$_____/
  |  $$$$$$/|  $$$$$$$| $$  | $$| $$|  $$$$$$$
   \______/  \_______/|__/  |__/|__/ \_______/

  A cli tool to quickly generate boilerplate code for developing 
  webapps in golang. The tool can be used as a standalone tool or 
  together with the go generate command to generate various parts 
  of the application such as models, routers etc`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

//go:embed assets
var assets embed.FS

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.genie.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
