package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделить строку на слайс строк.
	parts := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов,
	// вид активности и продолжительность.
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный формат данных")
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, err
	}
	// Преобразовать третий элемент слайса в time.Duration.
	// В пакете time есть метод для парсинга строки в time.Duration. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}
	// Если всё прошло без ошибок, верните количество шагов, вид активности, продолжительность и nil (для ошибки).
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
	// Для вычисления дистанции:
	// рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	// Соответствующая константа уже определена в пакете.
	stepLength := height * stepLengthCoefficient
	// умножьте пройденное количество шагов на длину шага.
	distance := float64(steps) * stepLength
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
	return distance / mInKm
	// Обратите внимание, что целочисленную переменную steps необходимо будет привести к другому числовому типу.
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	// Вычислить дистанцию с помощью distance().
	dist := distance(steps, height)
	// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах. Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Получить значения из строки данных с помощью функции parseTraining(),
	// обработать возможные ошибки и вывести их в лог с помощью log.Println(err).
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	// Проверить, какой вид тренировки был передан в строке, которую парсили (лучше использовать switch).
	// Для каждого из видов тренировки рассчитать дистанцию, среднюю скорость и калории.
	// Для каждого вида тренировки сформировать и вернуть строку, образец которой был представлен выше.
	// Если был передан неизвестный тип тренировки, вернуть ошибку с текстом неизвестный тип тренировки.
	switch activity {
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activity, duration.Hours(), dist, speed, calories), nil
	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activity, duration.Hours(), dist, speed, calories), nil
	default:
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверить входные параметры на корректность.
	// Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверные параметры ввода")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	// Рассчитать и вернуть количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверить входные параметры на корректность.
	// Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверные параметры ввода")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	// Рассчитать количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	calories := (weight * speed * duration.Minutes()) / minInH
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient.
	// Соответствующая константа объявлена в пакете. Вернуть полученное значение.
	return calories * walkingCaloriesCoefficient, nil
}
