// lint:nolint
package action

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate schema for entity",
	Long: `Generate schema for entity struct in package entity.
A entity struct must implement github.com/kcmvp/got/dbx/IEntity
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("so nice to meet you!")
		return nil
	},
	ValidArgs: []string{"schema", "repository", "service"},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
