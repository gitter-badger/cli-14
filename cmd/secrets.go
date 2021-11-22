package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(secretsCmd)
}

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Enable user secrets",
	Long:  "Enable user secrets for JT Framework",
	Run: func(cmd *cobra.Command, args []string) {
		enableUserSecrets()
	},
}
