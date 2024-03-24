// Композиция
/*
	Если нам необходимо получить все возможности структуры Car и дополнить их в классе 
	Пожарная машина (FireEngine), то мы можем использовать композицию 
	(сделать FireEngine членом Car):
*/
type Car struct {
	// … содержимое
  }
  
type FireEngine struct {
	basis Car
	// … дополнение
}

// Встраивание
/*
	Допустим структура Car имеет метод Drive. Мы должны скопировать точное поведение 
	метода Drive в структуре FireEngine.
*/
// Для этого мы можем применить делегирование:
type Car struct {
	// … содержимое
  }
  
func (c *Car) Drive() { … }

type FireEngine struct {
	basis Car
	// … дополнение
}

func (fe *FireEngine) Drive() { fe.basis.Drive() }

/*
	Однако оно ведёт к дублированию кода. Поэтому имеет механизм Встраивание, 
	что позволяет значительно сократить код:
*/
type Car struct {
	// … содержимое
  }
  
func (c *Car) Drive() { … }

type FireEngine struct {
	Car
	// … дополнение
}

// Интерфейсы
/*
	Допустим, что наше приложение расширяется и в ней появляется всё больше видов 
	специализированных машин: Полицейская Машина (PoliceCar), Машина Скорой Помощи 
	(AmbulanceCar), Поливомоечная машина (WateringCar).

	Все они должны иметь метод Drive, однако реализует его каждая по-разному. Например, 
	PoliceCar едет со звуком сирены, а WateringCar во время поездки поливает дорогу водой.
	То есть, мы должны определить "поведение", которое должно присутствовать в каждой из 
	этих структур, но реализовано оно может быть по-разному.

	В таком случае на сцену и выходят интерфейсы (interfaces). Интерфейсы определяют, 
	что тип делает, а не кем он является.
*/
type IDriveable interface {
	Drive()
  }
  
type Car struct {
// … 
}

type PoliceCar struct {
	// … 
}

func (c Car) Drive() {
	fmt.Println("Просто еду по дороге")
}

func (pc PoliceCar) Drive() {
	fmt.Println("Еду по дороге с мигалкой. Виу-виу!")
}

func main() {
cars := []IDriveable{&Car{}, &PoliceCar{}}
for _, vehicle := range cars {
	vehicle.Drive()
	// => Просто еду по дороге
	// =>  Еду по дороге с мигалкой. Виу-виу!
}
}

/* ========================= */
/*
	Реализуйте интерфейс IVoiceable для структур Cat, Cow и Dog так, чтобы при вызове 
	метода Voice экземпляр структуры Cat возвращал строку "Мяу", экземпляр Cow строку 
	"Мууу", а экземпляр Dog сообщение Гав:
*/
package solution

type Voicer interface {
    Voice() string
}

type Cat struct {
    // … 
}

type Cow struct {
    // … 
}

type Dog struct {
	// … 
}

func (c Cat) Voice() string {
	return "Мяу"
}

func (cw Cow) Voice() string {
	return "Мууу"
}

func (d Dog) Voice() string {
	return "Гав"
}
