package domain

import (
	"fmt"
)

type StatementType string
type Command string
type ErrorMessage string

type Row struct {
	ID       uint32
	Username []byte
	Email    []byte
}

type Table struct {
	RowNums uint32
	Pages   []*Page
}

type Statement struct {
	Type StatementType
	Row  Row
}

type Page []*Row

var Tbl Table

const (
	UsernameMaxSize = 32
	EmailMaxSize    = 255
	UserNameOffset  = 4
	EmailOffset     = 36
	PageSize        = 4096
	TableMaxPages   = 100
	RowPerPage      = PageSize / TableMaxPages
	TableMaxRows    = RowPerPage * TableMaxPages
)

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
