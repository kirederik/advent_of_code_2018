package main

import (
	"fmt"
	"math"
)

func main() {
	gp, x, y := Solve(3, 5719)
	fmt.Printf("x = %+v\n", x)
	fmt.Printf("y = %+v\n", y)
	fmt.Printf("maxPower = %+v\n", gp)
	fmt.Println("-------------")

	var maxPower int
	var declining int
	var prev int
	var b int

	for i := 1; i < 300; i++ {
		var x1, y1 int
		gp, x1, y1 = Solve(i, 5719)
		if gp < 0 && gp < prev {
			declining++
		} else {
			declining = 0
		}
		if gp > maxPower {
			maxPower = gp
			x = x1
			y = y1
			b = i
			fmt.Printf("NEW MAX: For box size %d, found (%d, %d) with power %d\n", i, x1, y1, gp)
		} else {
			fmt.Printf("For box size %d, found (%d, %d) with power %d\n", i, x1, y1, gp)
			prev = gp
		}
		if declining > 10 {
			fmt.Printf("Aborting at %d due to continuosly declining value\n", i)
			break
		}

	}
	fmt.Printf("x = %+v\n", x)
	fmt.Printf("y = %+v\n", y)
	fmt.Printf("maxPower = %+v\n", maxPower)
	fmt.Printf("%d,%d,%d\n", x, y, b)
}

const gridSize = 300

func Solve(boxSize, serialNumber int) (int, int, int) {
	maxPower := -math.MaxInt32
	var x, y int

	for i := 0; i < gridSize-boxSize; i++ {
		for j := 0; j < gridSize-boxSize; j++ {
			gp := calculateGridPower(i, j, serialNumber, boxSize)
			if gp > maxPower {
				maxPower = gp
				x = i
				y = j
			}
		}
	}

	return maxPower, x, y
}

func calculatePower(x, y, serialNumber int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel = (powerLevel % 1000) / 100
	return powerLevel - 5
}

func calculateGridPower(x, y, serialNumber, boxSize int) int {
	var total int
	for i := x; i < x+boxSize; i++ {
		for j := y; j < y+boxSize; j++ {
			pw := calculatePower(i, j, serialNumber)
			total += pw
		}
	}
	return total
}
