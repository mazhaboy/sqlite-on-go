package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type StatementType string
type Command string
type ErrorMessage string

type Row struct {
	ID       uint32
	Username [UsernameMaxSize]byte
	Email    [EmailMaxSize]byte
}

type Table struct {
	RowNums     uint64
	CurrentPage uint64
	Pages       [TableMaxPages]*Page
}

type Page struct {
	CurrentRow uint64
	PageSize   uint64
	Rows       []*Row
}

type Statement struct {
	Type StatementType
	Row  Row
}

var table Table

const (
	UsernameMaxSize   = 32
	EmailMaxSize      = 255
	UserNameOffset    = 4
	EmailOffset       = 36
	TotalRowSizeToAdd = UsernameMaxSize + EmailMaxSize + UserNameOffset + EmailOffset
	PageSize          = 4096
	TableMaxPages     = 100
	RowPerPage        = PageSize / TableMaxPages
	TableMaxRows      = RowPerPage * TableMaxPages
)

const WhiteSpace = " "

const (
	SELECT StatementType = "select"
	INSERT StatementType = "insert"
)

const (
	QuantityArgs       ErrorMessage = "wrong quantity of arguments"
	InvalidStatement   ErrorMessage = "there is no such statement"
	BytesEmailLimit    ErrorMessage = "overflow max size of bytes 255"
	BytesUsernameLimit ErrorMessage = "overflow max size of bytes 32"
	TableIsFull        ErrorMessage = "table is full"
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

func ParseArgs(stm string) (uint32, [UsernameMaxSize]byte, [EmailMaxSize]byte, error) {
	args := strings.Split(stm, WhiteSpace)
	var (
		id       int64
		err      error
		username [UsernameMaxSize]byte
		email    [EmailMaxSize]byte
	)
	id, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return uint32(id), username, email, err
	}
	username, err = toUsername([]byte(args[2]))
	if err != nil {
		fmt.Println("error can not parse to [32]byte: ", err.Error())
		return uint32(id), username, email, err
	}
	email, err = toEmail([]byte(args[3]))
	if err != nil {
		fmt.Println("error can not parse to [255]byte: ", err.Error())
		return uint32(id), username, email, err
	}
	return uint32(id), username, email, err
}

func toUsername(src []byte) ([UsernameMaxSize]byte, error) {
	var res [UsernameMaxSize]byte
	err := errors.New(string(BytesUsernameLimit))
	if len(src) > UsernameMaxSize {
		return res, err
	}
	for i := range src {
		res[i] = src[i]
	}
	return res, nil
}

func toEmail(src []byte) ([EmailMaxSize]byte, error) {
	var res [EmailMaxSize]byte
	err := errors.New(string(BytesEmailLimit))
	if len(src) > EmailMaxSize {
		return res, err
	}
	for i := range src {
		res[i] = src[i]
	}
	return res, nil
}

func GetTable() *Table {
	return &table
}

func (t *Table) AddRow(row *Row) error {
	if len(t.Pages) <= TableMaxPages {
		t.RowNums++
		//fmt.Println(t.Pages[t.CurrentPage])
		//if t.Pages[t.CurrentPage].PageSize+TotalRowSizeToAdd > PageSize {
		//	t.CurrentPage++
		//}
		t.Pages[t.CurrentPage].Rows[t.Pages[t.CurrentPage].CurrentRow] = row
		t.Pages[t.CurrentPage].PageSize += TotalRowSizeToAdd
		t.Pages[t.CurrentPage].CurrentRow++
		return nil
	}
	return errors.New(string(TableIsFull))
}

type Handler interface {
	HandlerInput()
}

func (c Command) HandlerInput() {
	fmt.Println(c, ": command not found")
}

func (c StatementType) HandlerInput(message ErrorMessage) {
	fmt.Println(c, ": syntax error statement not valid ", message)
}
