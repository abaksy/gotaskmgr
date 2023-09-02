/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	utils "github.com/abaksy/gotaskmgr/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `List the tasks still pending along with their 
	unique IDs. This ID can be used with the do command to
	complete a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.ListTasks()
		if err != nil {
			log.Default().Fatalf("error while listing tasks: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
