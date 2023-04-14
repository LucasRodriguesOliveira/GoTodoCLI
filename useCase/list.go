package useCase

import (
	bp "GoTodoCLI/model/bypri"
	to "GoTodoCLI/service/todo"
	ty "GoTodoCLI/service/types"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func ListRun(dbFile ty.FileData) func(cobra *cobra.Command, args []string) {
	return func(cobra *cobra.Command, args []string) {
		items, err := to.ReadItems(dbFile)

		if err != nil {
			fmt.Println(err)
			panic(1)
		}

		sort.Sort(bp.ByPri(items))

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
		defer w.Flush()

		for _, i := range items {
			fmt.Fprintln(w,i.Label()+"\t"+i.PrettyPriority()+"\t"+i.Text+"\t")
		}
	}
}
