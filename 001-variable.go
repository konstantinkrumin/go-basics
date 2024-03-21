// Существует два способа объявить переменную в Go. Длинная запись с ключевым словом var:
var num int = 11
//И короткая запись:
num := 22

// двоеточие используется только при инициализации
num := 22
num = 33

/* ========================= */

package main

import "fmt"

func main() {
	firstName := "John"
	lastName := "Smith"

	fmt.Println(firstName, lastName)
}