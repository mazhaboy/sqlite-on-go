package statement

import (
	"fmt"
	"sqlite/domain"
	"strconv"
	"strings"
)

type Executer interface {
	Insert(statement string)
	Select(statement string)
}

func Insert(statement string) {
	if domain.Tbl.RowNums <= domain.TableMaxRows {
		id, username, email := parseArgs(statement)
		if len(username) <= 32 && len(email) <= 255 {
			newRow := &domain.Row{
				ID:       id,
				Username: username,
				Email:    email,
			}
			domain.Tbl.RowNums++
			newPage := make(domain.Page, 0)
			newPage = append(newPage, newRow)
			domain.Tbl.Pages = append(domain.Tbl.Pages, &newPage)
		}
		fmt.Println("Executed")
	}

}

func Select(statement string) {
}

func parseArgs(stm string) (uint32, []byte, []byte) {
	args := strings.Split(stm, domain.WhiteSpace)
	id, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		fmt.Println("error can not parse to uint32: err")
	}
	username := []byte(args[2])
	email := []byte(args[3])
	return uint32(id), username, email
}
