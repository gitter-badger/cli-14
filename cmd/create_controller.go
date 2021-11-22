package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var name string
var path string
var ns string

func init() {
	createControllerCmd.Flags().StringVarP(&name, "name", "n", "HomeController", "--name HomeController")
	createControllerCmd.Flags().StringVarP(&path, "path", "p", "./src/controllers/", "--path ./src/controllers/")
	createControllerCmd.Flags().StringVar(&ns, "ns", "JtF\\Application", "--ns JtF\\Application")
	rootCmd.AddCommand(createControllerCmd)
}

var createControllerCmd = &cobra.Command{
	Use:     "create:controller",
	Aliases: []string{"cc"},
	Short:   "Create http controller",
	Long:    `Create http controller for JT Framework`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			name = args[0]
		}
		if len(args) > 1 {
			path = args[1]
		}
		if len(args) > 2 {
			ns = args[2]
		}
		createController(name, path, ns)
	},
}

func createController(name, path, ns string) {
	filePath := filepath.Join(path, name+".php")
	os.MkdirAll(path, os.ModePerm)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("\u274C Unable to create file...")
		os.Exit(1)
	} else {
		_, err := file.WriteString(generateController(name, ns))
		if err != nil {
			fmt.Println("\u274C Oops unknown error when writing to file...")
		} else {
			fmt.Printf("\u2705 Controller %s created\n", name)
		}
	}

	defer file.Close()
}
