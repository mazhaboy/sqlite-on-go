package domain

import "fmt"

type Statement string
type Command string

const (
	SELECT Statement = "select"
	INSERT Statement = "insert"
)

var StatementEnum = struct {
	SELECT Statement
	INSERT Statement
}{SELECT: SELECT, INSERT: INSERT}

const (
	HELP  Command = ".help"
	CLEAR Command = ".clear"
	EXIT  Command = ".exit"
)

var CommandEnum = struct {
	HELP  Command
	CLEAR Command
	EXIT  Command
}{HELP: HELP, CLEAR: CLEAR, EXIT: EXIT}

type Handler interface {
	HandlerInput()
}

func (c Command) HandlerInput() {
	fmt.Println(c, ": command not found")
}

func (c Statement) HandlerInput() {
	fmt.Println(c, ": syntax error statement not valid")
}
