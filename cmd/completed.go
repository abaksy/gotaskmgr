/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	constants "github.com/abaksy/gotaskmgr/const"
	db "github.com/abaksy/gotaskmgr/db"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List completed tasks",
	Long: `List only tasks that were completed today.
	Use "task do" to mark a task as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.ListTasks(constants.COMPLETE)
		if err != nil {
			fmt.Printf("error while listing tasks: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
