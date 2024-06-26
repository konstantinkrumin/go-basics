/*
	Для работы со строками в Go существует стандартный пакет strings, который содержит 
	основные функции. С некоторыми мы уже встречались в первом модуле 
	(например strings.ReplaceAll). Теперь рассмотрим список самых часто встречающихся 
	функций:
*/
import "strings"

// проверяет наличие подстроки в строке
strings.Contains("hello", "h") // true

// разбивает строку по Юникод символам или по переданному разделителю
strings.Split("hello", "") // ["h", "e", "l", "l", "o"]

// склеивает строки из слайса с разделителем
strings.Join([]string{"hello","world!"}, " ") // "hello world!"

// обрезает начальные и конечные символы строки, содержащиеся во втором аргументе
strings.Trim(" hey !", " ") // "hey !"

/*
	Очень важная часть пакета strings — это Builder. Когда необходимо собрать большую 
	строку по каким-то правилам, использование конкатенации — не лучшее решение, 
	потому что каждая операция создает новую строку, что сильно влияет на 
	производительность при большом количестве операций. Такая задача решается 
	с помощью билдера:
*/
import "strings"

sb := &strings.Builder{}

sb.WriteString("hello")
sb.WriteString(" ")
sb.WriteString("world")

sb.String() // "hello world"

/* ========================= */
/*
	В пакете unicode есть функция unicode.Is(unicode.Latin, rune), которая проверяет, 
	что руна является латинским символом или нет.

	Реализуйте функцию latinLetters(s string) string, которая возвращает только латинские 
	символы из строки s. Например:
*/
package solution

import (
	"strings"
	"unicode"
)

func latinLetters(s string) string {
	arr := []rune(s)
	sb := &strings.Builder{}

	for _, char := range arr {
		if (unicode.Is(unicode.Latin, char)) {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}