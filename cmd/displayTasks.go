package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"example.com/toDoList/internal/task"
	"github.com/mergestat/timediff"
)

func DisplayTasks(tasks []task.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tName\tDone\tCreatedAt\tModifiedAt")
	for _, t := range tasks {
		createdAt := timediff.TimeDiff(t.CreatedAt)
		modifiedAt := timediff.TimeDiff(t.ModifiedAt)

		fmt.Fprintf(w, "%d\t%s\t%t\t%s\t%s\n", t.ID, t.Name, t.Done, createdAt, modifiedAt)
	}
	w.Flush()
}

func DisplayTask(task task.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tName\tDone\tCreatedAt\tModifiedAt")
	createdAt := timediff.TimeDiff(task.CreatedAt)
	modifiedAt := timediff.TimeDiff(task.ModifiedAt)

	fmt.Fprintf(w, "%d\t%s\t%t\t%s\t%s\n", task.ID, task.Name, task.Done, createdAt, modifiedAt)
	w.Flush()
}
