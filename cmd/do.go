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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do tasknumber",
	Short: "Complete a task",
	Long: `Mark the task number supplied as completed in the database
	Post this the task will no longer be available in the task list`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("error running task do: could not parse task ID")
		}
		err = db.DoTask(taskID)
		if err != nil {
			fmt.Printf("failed to complete task: %s\n", err)
		} else {
			fmt.Println("completed task!")
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
