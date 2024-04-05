// Так как строка — это массив байт, ее можно обойти с помощью цикла for:
package main

import (
    "fmt"
)

func main() {
    s := "hello"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}
/*
	Вывод:

		h
		e
		l
		l
		o
*/

/*
	Таким способом можно обойти только строки, состоящие из ASCII символов. Если строка 
	содержит мультибайтовые символы, вывод будет некорректен:
*/
package main

import (
    "fmt"
)

func main() {
    s := "привет"
    for i := 0; i < len(s); i++ {
        fmt.Println(string(s[i]))
    }

}

/* ========================= */
/*
	Реализуйте функцию shiftASCII(s string, step int) string, которая принимает на вход 
	состоящую из ASCII символов строку s и возвращает новую строку, где каждый символ 
	из входящей строки сдвинут вперед на число step. Например:

	shiftASCII("abc", 0) // "abc"
	shiftASCII("abc1", 1) // "bcd2"
	shiftASCII("bcd2", -1) // "abc1"
	shiftASCII("hi", 10) // "rs"
*/
package solution

func shiftASCII(s string, step int) string {
	arr := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		arr = append(arr, byte(s[i]) + byte(step))
	}

	return string(arr)
}