package main

import (
	"fmt"
	"math"
)

// Структура для хранения информации о грузе
type Cargo struct {
	Weight      float64 // вес груза
	Floor       int     // этаж
	HasElevator bool    // наличие работающего лифта
}

// Метод для расчёта базовой стоимости в зависимости от массы
func (c *Cargo) BaseCost() float64 {
	switch {
	case c.Weight <= 50:
		return 300
	case c.Weight <= 100:
		return 1000
	case c.Weight <= 300:
		return 2000
	default:
		return 0 // В случае, если груз весит больше 300 кг (можно добавить отдельную логику для этого)
	}
}

// Метод для расчёта дополнительной стоимости за подъём вручную
func (c *Cargo) ManualLiftCost() float64 {
	if c.HasElevator {
		return 0 // Если есть лифт, доплата не требуется
	}

	// Рассчитаем стоимость ручного подъёма
	// 300 руб за этаж за каждые 100 кг груза
	additionalCostPerFloor := 300.0
	weightMultiplier := math.Ceil(c.Weight / 100.0)

	// Оплата за подъём на этажи (учитываем, что с 1 этажа не платят, т.е. Floor - 1)
	floorsToClimb := float64(c.Floor - 1)
	return additionalCostPerFloor * weightMultiplier * floorsToClimb
}

// Метод для расчёта общей стоимости
func (c *Cargo) TotalCost() float64 {
	return c.BaseCost() + c.ManualLiftCost()
}

func main() {
	// Пример ввода
	var weight float64 = 3000
	var floor int = 8
	var hasElevator bool = true

	fmt.Println("Введите вес груза (в кг):")
	fmt.Scan(&weight)

	fmt.Println("Введите этаж:")
	fmt.Scan(&floor)

	fmt.Println("Есть ли лифт? (1 - да, 0 - нет):")
	var elevatorInput int
	fmt.Scan(&elevatorInput)
	hasElevator = elevatorInput == 1

	// Создание объекта Cargo
	cargo := Cargo{
		Weight:      weight,
		Floor:       floor,
		HasElevator: hasElevator,
	}

	// Расчёт и вывод итоговой стоимости
	totalCost := cargo.TotalCost()
	fmt.Printf("Общая стоимость подъёма груза: %.2f руб.\n", totalCost)
}
