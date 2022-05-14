package domain

import (
	"fmt"
)

type StatementType string

type Row struct {
	ID       uint32
	Username [UsernameMaxSize]byte
	Email    [EmailMaxSize]byte
}

type Page [RowMaxSize]*Row

type Table struct {
	RowNums uint32
	Pages   Page
}

const (
	UsernameMaxSize = 32
	EmailMaxSize    = 255
	TableMaxPages   = 100
	RowMaxSize      = 12
)

type Statement struct {
	Type StatementType
	Row  Row
}

type Command string
type ErrorMessage string

const WhiteSpace = " "

const (
	SELECT StatementType = "select"
	INSERT StatementType = "insert"
)

const (
	QuantityArgs     ErrorMessage = "wrong quantity of arguments"
	InvalidStatement ErrorMessage = "there is no such statement"
)

var StatementEnum = struct {
	SELECT StatementType
	INSERT StatementType
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

func (c StatementType) HandlerInput(message ErrorMessage) {
	fmt.Println(c, ": syntax error statement not valid ", message)
}
