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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the to-do list",
	Long: `Delete command will delete a task with the given ID from the to-do list.
	Delete command is different from the complete command, as it will remove the task from the list.
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

		if err := service.Delete(id); err != nil {
			log.Fatalf("could not delete task: %v", err)
		}
		log.Printf("Task with ID %d has been deleted", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
