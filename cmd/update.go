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

var NewName string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task in the to-do list",
	Long: `Update command allows you to update a task in the to-do list.
	You can update a task by providing the task id and the new name.
	Name can be specified with an argument after id, or with the --name flag.
	If no id is provided, an error will be returned.
	If no name is provided, an error will be returned.
	
	All available flags are:
	--name: new name of the task`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)
		if len(args) == 0 {
			log.Printf("Please provide a task ID")
			return
		}

		if len(args) > 1 {
			NewName = args[1]
		}
		if NewName == "" {
			log.Printf("Please provide a new name for the task")
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

		updatedTask, err := service.Update(id, NewName)
		if err != nil {
			log.Fatalf("could not delete task: %v", err)
		}
		DisplayTask(updatedTask)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
	updateCmd.Flags().StringVarP(&NewName, "name", "n", "", "New name of the task")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
