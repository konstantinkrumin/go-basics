// Функции в Go объявляются через ключевое слово func:
func multiply(x int, y int) int {
    return x * y
}

/* 
	Из одной функции можно возвращать несколько значений. 
	Чаще всего это используется для возвращения ошибок:
*/
package math

import "errors"

func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("cannot divide on zero")
    }

    return x / y, nil
}

/* ========================= */
// Функция конвертации числа в строку
package solution

import "strconv"


func IntToString(num int) string {
	return strconv.Itoa(num)
}
