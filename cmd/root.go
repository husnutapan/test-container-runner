package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "test-container-runner",
	Short: "Test container runner CLI in Go",
	Long:  `Test container runner CLI in Go`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
