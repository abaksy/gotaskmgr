/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/abaksy/gotaskmgr/db"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long: `Remove a task from the database
	After this, the task will not be listed, and cannot be marked as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("error running task do: could not parse task ID")
		}
		err = db.RemoveTask(taskID)
		if err != nil {
			fmt.Printf("failed to remove task: %s\n", err)
		} else {
			fmt.Printf("deleted task")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
