package main

import (
	"fmt"
	"github.com/ItsWewin/superfactory/xerror"
	"os"
	"strings"
	"time"
)

const (
	CREATE = "CREATE"
	TABLE  = "TABLE"
)

func main() {
	args := os.Args

	err := ValidAndExecute(args)
	if err != nil {
		fmt.Println("execute failed", err.Message)
	}
}

func ValidAndExecute(args []string) *xerror.Error {
	command := args[1]
	object := args[2]
	objectName := args[3]

	fmt.Println(args)

	switch strings.ToUpper(command) {
	case CREATE:
		return dealWith(command, object, objectName)
	default:
		return xerror.NewError(nil, xerror.Code.CParamsError, "unkown command")
	}
}

func dealWith(command, object, objectName string) *xerror.Error {
	c := newCreateCommand(command, object, objectName)
	return c.Execute()
}

type CreateCommand struct {
	Command    string
	Object     string
	ObjectName string
}

func newCreateCommand(command, object, objectName string) *CreateCommand {
	return &CreateCommand{
		Command:    command,
		Object:     object,
		ObjectName: objectName,
	}
}

func (c *CreateCommand) Execute() *xerror.Error {
	switch strings.ToUpper(c.Object) {
	case TABLE:
		return createTable(c.ObjectName)
	default:
		return xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "args is err")
	}

	return nil
}

func createTable(table string) *xerror.Error {
	fileName := fmt.Sprintf("create_table_%s_%d.sql", table, time.Now().Unix())

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return xerror.NewError(nil, xerror.Code.SFileExisted, "file is existed")
	}

	if _, err := os.Create(fileName); err != nil {
		return xerror.NewErrorf(err, xerror.Code.BCreateFileFailed, "create failed failed")
	}

	return nil
}
