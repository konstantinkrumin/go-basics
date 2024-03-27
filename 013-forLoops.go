/*
	Из-за стремления создателей Go к простоте и однозначности обход коллекций 
	осуществляется только через цикл for. Условное выражение пишется без скобок 
	и тело цикла всегда должно быть внутри { }:
*/
nums := make([]int, 0, 10)

// начиная с 0; пока i меньше 10; инкрементим i после каждого шага
for i := 0; i < 10; i++ {
    nums = append(nums, i)
}

fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

// При необходимости for используется как цикл while в других языках:
i := 0
nums := make([]int, 0, 10)

for i < 10 {
    nums = append(nums, i)
    i++
}

fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

// Если опустить условное выражение, можно уйти в бесконечный цикл:
i := 0

nums := make([]int, 0, 10)

for {
    if i == 10 {
        break
    }

    nums = append(nums, i)
    i++
}

fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

/*
	Если условное выражение возвращает false или был вызван break в теле, то происходит 
	остановка цикла. Перейти на следующую итерацию можно с помощью ключевого слова continue:
*/
nums := make([]int, 0, 10)

// добавляем только четные числа в слайс
for i := 0; i < 10; i++ {
    if i % 2 != 0 {
        continue
    }

    nums = append(nums, i)
}

fmt.Println(nums) // [0 2 4 6 8]

/*
	Для обхода коллекции в Go есть "синтаксический сахар" range. Эта конструкция обходит 
	слайс, возвращая пару - индекс и элемент на каждом шаге:
*/
names := []string{"John", "Harold", "Vince"}

// i — это индекс, name — это значение на текущем шаге цикла
for i, name := range names {
    fmt.Println("Hello ", name, " at index ", i)
}

// => Hello  John  at index  0
// => Hello  Harold  at index  1
// => Hello  Vince  at index  2

// Если пропустить вторую переменную, то получим только индексы:
for i := range names {
    fmt.Println("index = ", i)
}

// => index =  0
// => index =  1
// => index =  2

// Можно пропустить первую переменную, это можно сделать с помощью _:
for _, name := range names {
    fmt.Println("Hello ", name)
}

// => Hello  John
// => Hello  Harold
// => Hello  Vince

/* ========================= */
/*
	Реализуйте функцию func Map(strs []string, mapFunc func(s string) string) []string, 
	которая преобразует каждый элемент слайса strs с помощью функции mapFunc и возвращает 
	новый слайс. Учтите, что исходный слайс, который передается как strs, не должен 
	измениться в процессе выполнения.
*/
package solution

func Map(strs []string, mapFunc func(s string) string) []string {
	newStrs := make([]string, 0, len(strs))

	for i := 0; i < len(strs); i++ {
		newStrs = append(newStrs, mapFunc(strs[i]))
	}

	return newStrs
}

// OR

package solution

func Map(strs []string, mapFunc func(s string) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}

	return mapped
}