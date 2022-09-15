// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
package main

import (
	"fmt"
	"math"
	"sort"
)

func curGroup(a float64) float64 {
	b := a - math.Mod(a, 10)
	if a > 0 {
		a = a - b
	} else {
		a = a + b
	}
	return a
}

func main() {
	temperature := []float64{-25.4, -27.0, 1, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Slice(temperature, func(i, j int) bool {
		return temperature[i] < temperature[j]
	})

	newTemperature := map[int][]float64{}
	innerSlice := []float64{}
	curFirst := temperature[0]

	curGroup := int(curFirst - math.Mod(curFirst, 10))

	innerSlice = append(innerSlice, float64(curFirst))

	fmt.Println(temperature)

	for i := 1; i < len(temperature); i++ {
		if math.Abs(temperature[i]-curFirst) <= 10 {
			innerSlice = append(innerSlice, temperature[i])
		} else {
			newTemperature[curGroup] = innerSlice
			curFirst = temperature[i]
			curGroup = int(curFirst - math.Mod(curFirst, 10))
			innerSlice = []float64{temperature[i]}

		}
		if i == len(temperature)-1 {
			newTemperature[curGroup] = innerSlice
		}
	}

	fmt.Println(newTemperature)
}
