package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"test-container-runner/cmd/internal"
)

func newStartCommand() *cobra.Command {
	var slice []string
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "This command will create corresponding test container",
		Long:  `This command will create corresponding test container`,
		Run: func(cmd *cobra.Command, args []string) {
			containerUrl, err := cmd.Flags().GetString("containerUrl")
			if err != nil {
				fmt.Println("container url parameter is not exist")
				return
			}
			exposePorts, err := cmd.Flags().GetStringSlice("exposePorts")
			if err != nil {
				fmt.Println("expose ports parameter is not exist")
				return
			}
			container := internal.NewContainer(containerUrl, exposePorts)
			container.Start()
		},
	}
	startCmd.Flags().String("containerUrl", "", "")
	startCmd.Flags().StringSliceVarP(&slice, "exposePorts", "", []string{}, "")
	return startCmd
}

func init() {
	startCmd := newStartCommand()
	rootCmd.AddCommand(startCmd)

}
