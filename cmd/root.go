/*
Copyright © 2021 Dan Webb<dan.webb@damacus.io>

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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/damacus/clone-org-repos/checkout"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clone-org-repos",
	Short: "A tool to clone all repositories in a github org",
	Long:  `clone-org-repos allows you to clone all repositories and update them within a given file path`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN")
		org := getStringFlag("org", cmd)
		path := getStringFlag("path", cmd)
		home, _ := os.UserHomeDir()
		if path == "" {
			path = home
		}
		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("%s/%s", home, path)
		}
		checkout.Checkout(token, org, path)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", "", "Path to checkout repositories to, defaults to user's home directory")
	rootCmd.PersistentFlags().StringP("org", "o", "", "Name of the org you wish to checkout")
	rootCmd.MarkPersistentFlagRequired("org")
}

func getStringFlag(flagName string, cmd *cobra.Command) string {
	name, err := cmd.Flags().GetString(flagName)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
