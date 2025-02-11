package fmt_util

import (
	"fmt"
)

func SprintfError(err error) string {
	format := "\tType: %T\n" +
		"\tGo value: %#v\n" +
		"\tValue: %v\n" +
		"\tMessage: %s"
	return fmt.Sprintf(format, err, err, err, err.Error())
}
