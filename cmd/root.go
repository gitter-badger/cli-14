package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "jt",
	Short:   "\033[92mCLI for JT Framework\033[0m",
	Long:    "\033[92mFast and powerfull cli for JT Framework\033[0m",
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		printLogo()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
