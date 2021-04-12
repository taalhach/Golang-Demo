package internal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:  "velocity_works",
	Long: "velocity_works: velocity_worker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("works")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Printf("error while executing command: %v\n", err)
		os.Exit(1)
	}
}
