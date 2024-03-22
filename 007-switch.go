/*
	В Go присутствует единственная альтернатива if — конструкция switch. 
	Для этой конструкции используется стандартный синтаксис, но логика работы отличается 
	от С-подобных языков. Когда срабатывает условие какого-либо case, программа выполняет 
	блок и выходит из конструкции switch без необходимости писать break:
*/
x := 10

switch x {
    default: // default всегда выполняется последним независимо от расположения в конструкции
        fmt.Println("default case")
    case 10:
        fmt.Println("case 10")
}
// output: case 10

/*
	Однако при необходимости можно реализовать логику С-подобных языков и «провалиться» 
	в следующий case:
*/
x := 10

switch { // выражение отсутствует. Для компилятора выглядит как: switch true
    default:
        fmt.Println("default case")
    case x == 10:
        fmt.Println("equal 10 case")
        fallthrough
    case x <= 10:
        fmt.Println("less or equal 10 case")
}

/* 
Output:
	equal 10 case
	less or equal 10 case
*/

/* ========================= */
/*
	Реализуйте функцию ModifySpaces(s, mode string) string, 
	которая изменяет строку s в зависимости от переданного режима mode
*/
package solution

import (
	"strings"
)

func ModifySpaces(s, mode string) string {
	var result string

	switch mode {
		default:
			result = strings.ReplaceAll(s, " ", "*")
		case "dash":
			result = strings.ReplaceAll(s, " ", "-")
		case "underscore":
			result = strings.ReplaceAll(s, " ", "_")
	}

	return result
}
