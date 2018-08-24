package helper

import "fmt"

func MessageFormat(msg ...interface{}) string {
	if len(msg) == 1 {
		return msg[0].(string)
	}
	if len(msg) > 1 {
		return fmt.Sprintf(msg[0].(string), msg[1:]...)
	}
	return ""
}

func MessageFormatEqual(actual, expected interface{}, msg ...interface{}) string {
	var message string
	if msg == nil || len(msg) == 0 {
		message = fmt.Sprintf("%v != %v", actual, expected)
	} else {
		message = MessageFormat(msg)
	}
	return message
}
