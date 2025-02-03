package fmt

import (
	go_fmt "fmt"
)

func SprintfError(err error) string {
	format := "\tType: %T\n" +
		"\tGo value: %#v\n" +
		"\tValue: %v\n" +
		"\tMessage: %s"
	return go_fmt.Sprintf(format, err, err, err, err.Error())
}
