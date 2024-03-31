// Как и слайс, мапу можно обойти с помощью конструкции for range:
idToName := map[int64]string{1: "Alex", 2: "Dan", 3: "George"}

// первый аргумент — ключ, второй — значение
for id, name := range idToName {
    fmt.Println("id: ", id, "name: ", name)
}
/* Вывод:
	id:  1 name:  Alex
	id:  2 name:  Dan
	id:  3 name:  George
*/

// Стоит учитывать, что порядок ключей в мапе рандомизирован:
numExistence := make(map[int]bool, 0)

// записали по порядку числа от 0 до 9
for i := 0; i < 10; i++ {
    numExistence[i] = true
}

// обходим мапу и выводим ключи
for num := range numExistence {
    fmt.Println(num)
}
/*
	Вывод:
		8
		1
		2
		3
		6
		7
		9
		0
		4
		5
*/

/* ========================= */
/*
	Реализуйте функцию MostPopularWord(words []string) string, которая возвращает самое 
	часто встречаемое слово в слайсе. Если таких слов несколько, то возвращается первое 
	из них.
*/
package solution

func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int)

	for _, word := range words {
		count, exists := wordsCount[word]
		if (exists) {
			wordsCount[word] = count + 1
		} else {
			wordsCount[word] = 1
		}
	}

	maxCount := 0
	var mostPopular string = ""

	for word, count := range wordsCount {
		if (count > maxCount) {
			maxCount = count
			mostPopular = word
		}
	}

	return mostPopular
}