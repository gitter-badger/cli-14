package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var port int
var folder string

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8000, "--port 1111")
	serveCmd.Flags().StringVarP(&folder, "folder", "f", "./public", "--folder ./public")
	serveCmd.MarkFlagDirname("folder")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:     "serve [port] [folder]",
	Aliases: []string{"ss", "server", "run"},
	Short:   "Serve your application",
	Long:    "Run your application based on JT Framework",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			port, _ = strconv.Atoi(args[0])
		}
		if len(args) > 1 {
			folder = args[1]
		}
		runServer(port, folder)
	},
}
