/*
Copyright © 2025 Ying Tu <yingtu35@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"example.com/toDoList/internal/task"
	"github.com/spf13/cobra"
)

var Name string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long: `Create a new task in the to-do list with a name.
	Name can be specified with an argument after the command, or with the --name flag.
	If neither is provided, an error will be returned.

	All available flags are:
	--name: name of the task, default is empty`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)

		var taskName string

		// if a name is provided as an argument, use it
		if len(args) > 0 {
			taskName = args[0]
		} else if Name != "" {
			taskName = Name
		} else {
			fmt.Println("Please provide a name for the task")
			return
		}

		service, err := task.NewService()
		if err != nil {
			log.Fatalf("could not create task service: %v", err)
		}

		task, err := service.Create(taskName)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		DisplayTask(task)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.
	createCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of the task")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
