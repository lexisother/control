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
	"reflect"

	"github.com/lexisother/control/lib"
	"github.com/spf13/cobra"
)

// Command layout/syntax: control configure <project> <key> <value>
// First of all, we should check if the project specified (first argument) is present in the `projects` map.
// Then, we verify if `key` is actually a valid property of the `Project` struct.
// If so, we set the value of the property to `value`.
// If not, we print an error message.
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure a project",
	Long:  `Configure a project`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("configure requires a project name, a key and a value")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		key := args[1]
		value := args[2]

		config := lib.ReadConfig()

		if key != "Location" {
			if loc := config.Projects[projectName].Location; loc == "" {
				fmt.Printf("Warning: Location for project '%s' isn't set! Most commands won't work...\n", projectName)
			}
		}

		if config.Projects == nil {
			config.Projects = make(map[string]lib.Project)
		}

		if _, ok := config.Projects[projectName]; !ok {
			fmt.Printf("Project %s not found, adding it...\n", projectName)
			config.Projects[projectName] = lib.Project{}
		}

		project := config.Projects[projectName]
		projectType := reflect.TypeOf(project)

		if _, ok := projectType.FieldByName(key); !ok {
			fmt.Printf("'%s' is not a valid configuration option!\n", key)
			return
		}

		projectValue := reflect.ValueOf(&project).Elem()
		projectField := projectValue.FieldByName(key)
		projectField.SetString(value)
		fmt.Printf("Setting %s to %s...\n", key, value)

		config.Projects[projectName] = project
		lib.WriteConfig(config)
		fmt.Println("Wrote config!")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
