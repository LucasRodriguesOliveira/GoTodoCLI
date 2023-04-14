package useCase

import (
	bp "GoTodoCLI/model/bypri"
	to "GoTodoCLI/service/todo"
	"fmt"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

func DoneRun() func(cobra *cobra.Command, args []string) {
	return func(cobra *cobra.Command, args []string) {
		items, err := to.ReadItems()

		if err != nil {
			fmt.Println(err)
		}

		if len(items) == 0 {
			fmt.Println("no items to mark as done, try adding one task first with `add`")
			return
		}

		i, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("%q is not a valid label\n", args[0])
			return
		}

		if i > 0 && i <= len(items) {
			items[i-1].Done = true
			fmt.Printf("%q marked as done\n", items[i-1].Text)

			sort.Sort(bp.ByPri(items))
			to.SaveItems(items)
			return
		}

		fmt.Printf("[%v] does not match any items", i)
	}
}
