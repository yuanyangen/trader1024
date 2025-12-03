package logs

import "fmt"

func Info(format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a...))
}
