/*
	Сортировка массива — распространненая задача в программировании. Во всех языках 
	существуют готовые решения для этой задачи, и Go — не исключение. Стандартный пакет 
	sort предоставляет функции для сортировки:
*/
nums := []int{2,1,6,5,3,4}

sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

fmt.Println(nums) // [1 2 3 4 5 6]

/*
	Рассмотрим функцию Slice(x interface{}, less func(i, j int) bool). В описании функции 
	присутствует неизвестный тип данных interface{}. Понятие интерфейса будет рассмотренно 
	в следующих модулях курса. Следует запомнить, что пустой интерфейс interface{} в Go 
	означает тип данных, под который подходит любой другой тип. Например:
*/
func Print(arg interface{}) {
    fmt.Println(arg)
}

func main() {
    Print("hello!")
    Print(123)
    Print([]int{1,5,10})
}

/*
	То есть в функцию Slice(x interface{}, less func(i, j int) bool) передается слайс 
	любого типа данных, как первый аргумент. Вторым аргументом передается функция, 
	которая берет элементы по индексу и определяет должен ли элемент по индексу i 
	находиться перед элементом по индексу j.

	"Под капотом" в функции sort.Slice используется быстрая сортировка. В пакете также 
	присутствует сортировка вставками sort.SliceStable:
*/
nums := []int{2,1,6,5,3,4}

sort.SliceStable(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

fmt.Println(nums) // [1 2 3 4 5 6]

/* ========================= */
/*
	Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64, которая возвращает 
	отсортированный слайс, состоящий из уникальных идентификаторов userIDs. Обработка 
	должна происходить in-place, то есть без выделения доп. памяти.
*/
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}