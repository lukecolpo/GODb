package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MetaCommand string

const (
	Exit MetaCommand = ".exit\n"
	Help MetaCommand = ".help\n"
)

type QueryType int

const (
	Insert QueryType = iota
	Select
)

type Statement struct {
	querytype QueryType
}

func prepareStatement(input string) (statement Statement, err error) {
	err = nil
	if strings.HasPrefix(input, "insert") {
		statement.querytype = Insert
	} else if strings.HasPrefix(input, "select") {
		statement.querytype = Select
	} else {
		end := 10
		if len(input) < 10 {
			end = len(input)
		}
		err = fmt.Errorf("Unrecognized query type for statement: %s", input[:end])
	}
	return

}

func executeStatement(statement Statement) error {
	switch statement.querytype {
	case Insert:
		fmt.Print("This is where we would do an insert.\n")
	case Select:
		fmt.Print("This is where we would do a select.\n")
	default:
		return nil
	}
	return nil
}

func main() {
	fmt.Print("GODb version 0.0.1\n")
	fmt.Print("Enter \".help\" for instructions\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("GODb> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if strings.HasPrefix(input, ".") {
			switch meta_command := MetaCommand(input); meta_command {
			case Help:
				fmt.Print("noop\n")
			case Exit:
				os.Exit(0)
			default:
				fmt.Printf("Unrecognized meta command %s", input)
			}
		} else {
			statement, err := prepareStatement(input)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Printf("statement: %+v\n", statement)
			err = executeStatement(statement)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}

