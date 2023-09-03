/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	constants "github.com/abaksy/gotaskmgr/const"
	db "github.com/abaksy/gotaskmgr/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add taskname",
	Short: "Add a task",
	Long: `The add command adds a task to be done to the database
	`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, constants.SPACE)
		creationDate := time.Now()

		task := db.Task{Name: taskName, CreationDate: creationDate, CompletionStatus: constants.IN_PROGRESS}

		err := db.AddTask(task)
		if err != nil {
			fmt.Printf("error running task add: %s\n", err)
		} else {
			fmt.Println("Added task!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
