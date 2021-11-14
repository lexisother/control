/*
Copyright © 2021 Alyxia Sother <lexisoth2005@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/lexisother/control/lib"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project by running the configured update command.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("update requires a project name")
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		projects := lib.ReadConfig().Projects
		names := make([]string, len(projects))

		i := 0
		for k := range projects {
			names[i] = k
			i++
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if val, ok := lib.ReadConfig().Projects[projectName]; ok {
			if val.UpdateCommand != "" {
				command := strings.Split(val.UpdateCommand, " ")
				err := lib.RunCmd(val.Location, command[0], command[1:])
				if err != nil {
					fmt.Printf("Something went wrong while running the command: %s", err)
					return
				}
			} else {
				fmt.Println("No update command configured for this project.")
			}
		} else {
			fmt.Printf("Project '%s' not found in config, use `control configure` to add it!\n", projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
