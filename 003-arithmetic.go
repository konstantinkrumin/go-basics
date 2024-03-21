// Любые операции осуществляются только над числами одинакового типа:
x := 5.05
y := 10

x + y  // invalid operation: x + y (mismatched types float64 and int)

/* 
	Чтобы осуществить сложение из прошлого примера, нам нужно конвертировать значения 
	к одному типу 
*/
x := 5.05
y := 10

x + float64(y)  // 15.05

/* ========================= */
// Функция нахождения минимального числа
package solution

import "math"

func MinInt(x, y int) int {
	min := math.Min(float64(x), float64(y))
	return int(min)
}
