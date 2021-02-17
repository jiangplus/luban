package cmd

import (
"fmt"

"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(submitCmd)
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit luban job to server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("luban job")
	},
}
