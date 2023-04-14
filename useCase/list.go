package useCase

import (
	bp "GoTodoCLI/model/bypri"
	to "GoTodoCLI/service/todo"
	ut "GoTodoCLI/util"
	tm "GoTodoCLI/model/todo"
	t  "GoTodoCLI/cmd/types"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func ListRun(f *t.FBool) func(cobra *cobra.Command, args []string) {
	return func(cobra *cobra.Command, args []string) {
		items, err := to.ReadItems()

		if err != nil {
			fmt.Println(err)
			panic(1)
		}

		if len(items) == 0 {
			fmt.Println("No task's found, use command <add> to add a task or more.")
			return
		}

		sort.Sort(bp.ByPri(items))

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
		defer w.Flush()

		done, all := (*f)["done"], (*f)["all"]
		for _, i := range items {
			switch {
			case all:
				addLine(w, &i)
			case done == i.Done:
				addLine(w, &i)
			}
		}
	}
}

func addLine(w *tabwriter.Writer, i *tm.Item) {
	fmt.Fprintln(w, line(i))
}

func line(i *tm.Item) string {
	return ut.SConcat(
		ut.FTab(i.Label()),
		ut.FTab(i.PrettyDone()),
		ut.FTab(i.PrettyPriority()),
		ut.FTab(i.Text),
	)
}
