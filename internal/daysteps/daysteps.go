package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Mandelshtam21/fitness-tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделить строку на слайс строк.
	parts := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Неверный формат данных")
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	// Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше 0")
	}
	// Преобразовать второй элемент слайса в time.Duration.
	// В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("Продолжительность должна быть больше 0")
	}
	// Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage().
	// В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	//Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.
	if steps <= 0 || duration <= 0 {
		return ""
	}
	// Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага.
	// Константа stepLength (длина шага) уже определена в коде.
	distance := float64(steps) * stepLength
	// Перевести дистанцию в километры, разделив её на число метров в километре (константа mInKm, определена в пакете).
	distanceKm := distance / mInKm
	// Вычислить количество калорий, потраченных на прогулке.
	// Функция для вычисления калорий WalkingSpentCalories() будет определена в пакете spentcalories, которую вы тоже реализуете.
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//Сформировать строку, которую будете возвращать, пример которой был представлен выше.
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
