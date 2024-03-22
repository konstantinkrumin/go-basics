// Строки в Go объявляются с типом string:
var s string = "hello"

// сокращенная запись
s := "hey"

/*
	Практически всегда для строк используются двойные кавычки. 
	Однако они не подходят, когда нужно написать несколько строк. 
	Для этого используются обратные кавычки:
*/
q := `
    SELECT *
    FROM person
    WHERE age > 18
`

/* ========================= */
package solution

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	str := strings.Trim(name, " ")
	str = strings.ToLower(str)
	str = strings.Title(str)

	return fmt.Sprintf("Привет, %s!", str)
}
