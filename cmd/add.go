/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	constants "github.com/abaksy/gotaskmgr/const"
	utils "github.com/abaksy/gotaskmgr/utils"
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

		task := utils.Task{Name: taskName, CreationDate: creationDate}

		err := utils.AddTask(task)
		if err != nil {
			log.Default().Fatalf("Error running task add: %s", err)
		}
		fmt.Println("Added task!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
