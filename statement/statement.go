package statement

import (
	"sqlite/domain"
	"strings"
)

type Executer interface {
	Insert(statement string)
	Select(statement string)
}

func Insert(statement string) {
	stmArgs := strings.Split(statement, domain.WhiteSpace)

}

func Select(statement string) {
	stmArgs := strings.Split(statement, domain.WhiteSpace)
}
