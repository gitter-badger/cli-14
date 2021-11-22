package cmd

import (
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
