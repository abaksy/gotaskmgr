/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do tasknumber",
	Short: "Complete a task",
	Long: `Mark the task number supplied as completed in the database
	Post this the task will no longer be available in the task list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
