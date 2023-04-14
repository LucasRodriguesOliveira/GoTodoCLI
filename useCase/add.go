package useCase

import (
	mo "GoTodoCLI/model/todo"
	to "GoTodoCLI/service/todo"
	ty "GoTodoCLI/service/types"
	"fmt"

	"github.com/spf13/cobra"
)

func AddRun(dbFile ty.FileData, priority *int) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No item to add")
		}

		items := []mo.Item{}
		for _, arg := range args {
			item := mo.Item{Text: arg}
			item.SetPriority(*priority)

			items = append(items, item)
		}

		dbItems, err := to.ReadItems(dbFile)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(dbItems) > 0 {
			items = append(dbItems, items...)
		}


		to.SaveItems(items, dbFile)
	}
}
