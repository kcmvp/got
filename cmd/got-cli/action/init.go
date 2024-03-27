package action

import "github.com/spf13/cobra"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project layout, configuration for project",
	Long:  `got-cli init`,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
