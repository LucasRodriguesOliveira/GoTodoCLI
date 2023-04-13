package useCase

import (
	"GoTodoCLI/model"
	serviceTodo "GoTodoCLI/service/todo"

	"github.com/spf13/cobra"
)

func AddRun(cmd *cobra.Command, args []string) {
	items := []model.Item{}
	for _, arg := range args {
		items = append(items, model.Item{ Text: arg })
	}

	serviceTodo.SaveItems(items)
}
