/*
Copyright © 2025 Ying Tu <yingtu35@gmail.com>
*/
package cmd

import (
	"log"

	"example.com/toDoList/internal/task"
	"github.com/spf13/cobra"
)

var Include string

var OrderBy string
var Desc bool
var Asc bool

var Done string

var Limit int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the to-do list",
	Long: `List command displays tasks in the to-do list.
	You can view all tasks or filter them by name, status.
	All available flags are:
	--include: filter tasks by id, name that includes the given string, default is id
	--done: filter tasks by status that is done, true, false, or all
	--orderBy: order tasks by name or createdAt, default is empty, which will sort by id
	--desc: sort tasks in descending order
	--limit: limit the number of tasks to display, default is 10`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)

		filter := task.Filter{
			Include: Include,
			OrderBy: OrderBy,
			Desc:    Desc,
			Done:    Done,
			Limit:   Limit,
		}

		service, err := task.NewService()
		if err != nil {
			log.Fatalf("could not create task service: %v", err)
		}

		tasks, err := service.List(filter)
		if err != nil {
			log.Fatalf("could not list tasks: %v", err)
		}

		DisplayTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().StringVarP(&Include, "include", "i", "", "Filter tasks by name that includes the given string")
	listCmd.Flags().StringVarP(&OrderBy, "orderBy", "o", "id", "Order tasks by id, name or createdAt")
	listCmd.Flags().BoolVar(&Desc, "desc", false, "Sort tasks in descending order")
	listCmd.Flags().StringVar(&Done, "done", "all", "Filter tasks by status that is done, true, false, or all")
	listCmd.Flags().IntVarP(&Limit, "limit", "l", 10, "Limit the number of tasks to display, default is 10")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
