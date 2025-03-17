/*
Copyright © 2025 Ying Tu <yingtu35@gmail.com>
*/
package cmd

import (
	"log"
	"strconv"

	"example.com/toDoList/internal/task"
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as undone",
	Long: `Undone command serves as the opposite of the complete command.
	Mark a task with the given ID as undone in the to-do list.
	If no ID is provided, an error will be returned.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)
		if len(args) == 0 {
			log.Printf("Please provide a task ID")
			return
		}
		service, err := task.NewService()
		if err != nil {
			log.Fatalf("could not create task service: %v", err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("could not convert task ID to integer: %v", err)
		}
		completedTask, err := service.Undone(id)
		if err != nil {
			log.Fatalf("could not undone task: %v", err)
		}
		DisplayTask(completedTask)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
