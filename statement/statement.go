package statement

import (
	"fmt"
	"sqlite/domain"
)

type Executer interface {
	Insert(statement string)
	Select(statement string)
}

func Insert(statement string) {
	table := domain.GetTable()
	if table.RowNums <= domain.TableMaxRows {
		id, username, email, err := domain.ParseArgs(statement)
		if err != nil {
			return
		}
		newRow := &domain.Row{
			ID:       id,
			Username: username,
			Email:    email,
		}
		if err = table.AddRow(newRow); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Executed")
	}
}

func Select(statement string) {
	table := domain.GetTable()
	for i := range table.Pages {
		for j := range table.Pages[i].Rows {
			fmt.Println(table.Pages[i].Rows[j].ID, table.Pages[i].Rows[j].Username, table.Pages[i].Rows[j].Email)
		}
	}
}
