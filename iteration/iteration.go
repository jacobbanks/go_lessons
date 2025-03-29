package iteration

import "strings"


func Repeat(char string, repeatTimes int) string {

	var repeated strings.Builder
	for i := 0; i < repeatTimes; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
