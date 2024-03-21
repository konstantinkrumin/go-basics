/*
	Логический тип в Go представлен привычными значениями true и false c операторами:
	— && (и)
	— == (равно)
	— || (или)
	— ! (не)
*/
true && false // false
false || true // true

// Из-за строгой типизации в Go можно сравнивать только одинаковые типы данных:
true == false // false
false == false // true

/*
	Когда необходимо проверить на истинность разные значения, нелогические типы нужно 
	привести к логическому:
*/
flag := true
text := "hello"

// вариант не сработает, потому что нельзя конвертировать строку в bool
flag && bool(text) // cannot convert text (type string) to type bool

// правильный вариант: если строка не пустая, то в ней есть текст
flag && text != "" // true

/* ========================= */
package solution

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}