/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	constants "github.com/abaksy/gotaskmgr/const"
	db "github.com/abaksy/gotaskmgr/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `List the tasks still pending along with their 
	unique IDs. This ID can be used with the do command to
	complete a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.ListTasks(constants.IN_PROGRESS)
		if err != nil {
			fmt.Printf("error while listing tasks: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
