package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbSyncCmd = &cobra.Command{
	Use:   "sync_db",
	Short: "Syncs database",
	Long:  "Syncs city government database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("works")
	},
}

func init() {
	rootCommand.AddCommand(dbSyncCmd)
}
