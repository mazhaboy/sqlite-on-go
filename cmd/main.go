package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sqlite/domain"
	"sqlite/statement"
	"strings"
)

func main() {
	// Hardcoded repl commands
	metaCommands := map[domain.Command]interface{}{
		domain.CommandEnum.HELP:  displayHelp,
		domain.CommandEnum.CLEAR: clearScreen,
	}

	statements := map[domain.StatementType]interface{}{
		domain.StatementEnum.SELECT: domain.StatementEnum.SELECT,
		domain.StatementEnum.INSERT: domain.StatementEnum.INSERT,
	}
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if strings.EqualFold(".", string(reader.Bytes()[0])) {
			cmd := domain.Command(text)
			if command, exists := metaCommands[domain.Command(text)]; exists {
				// Call a hardcoded function
				command.(func())()
			} else if strings.EqualFold(".exit", text) {
				// Close the program on the exit command
				return
			} else {
				// Pass the command to the parser
				cmd.HandlerInput()
			}
		} else {
			statementArgs := strings.Split(text, domain.WhiteSpace)
			stm := domain.StatementType(statementArgs[0])
			fmt.Println(statementArgs)
			if statementType, exists := statements[stm]; exists {
				// Call a hardcoded function
				switch statementType {
				case domain.INSERT:
					statement.Insert(text)
				case domain.SELECT:
					statement.Select(text)
				}
			} else {
				// Pass the statement to the parser
				stm.HandlerInput(domain.InvalidStatement)
			}

		}

		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

// cliName is the name used in the repl prompts
var cliName string = "sqlite"

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

// displayHelp informs the user about our hardcoded functions
func displayHelp() {
	fmt.Printf("Welcome to %v! These are the available commands: \n", cliName)
	fmt.Println(".help    - Show available commands")
	fmt.Println(".clear   - Clear the terminal screen")
	fmt.Println(".exit    - Closes your connection to", "sqlite")
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("error : ", err.Error())
	}
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)

	return output
}

func insertState() {
	fmt.Println("INSERT - insert to database")
}

func selectState() {
	fmt.Println("SELECT - select from database")
}
