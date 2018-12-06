package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer1, answer2 := PartIandII(os.Args[1])
	fmt.Println("Part I: ", answer1)
	fmt.Println("Part II: ", answer2)
}

type Coordinate struct {
	x, y int
}

type Point struct {
	minDistance int
	closest     *Coordinate
}

func readAll(filepath string) ([]Coordinate, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []Coordinate

	for {
		var x, y int
		_, err := fmt.Fscanf(f, "%d, %d\n", &x, &y)

		if err == io.EOF {
			return data, nil
		}
		if err != nil {
			return nil, err
		}

		data = append(data, Coordinate{x: x, y: y})
	}
}

func PartIandII(filepath string) (int, int) {
	coordinates, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	return calculateDistances(coordinates)
}

func calculateDistances(cs []Coordinate) (int, int) {
	var xMin, yMin, xMax, yMax int
	grid := map[Coordinate]int{}
	xMin = math.MaxInt32
	yMin = math.MaxInt32
	xMax = -1
	yMax = -1
	for _, c := range cs {
		if c.x < xMin {
			xMin = c.x
		}
		if c.x > xMax {
			xMax = c.x
		}
		if c.y < yMin {
			yMin = c.y
		}
		if c.y > yMax {
			yMax = c.y
		}
	}

	var totalArea int

	for j := yMin; j <= yMax; j++ {
		for i := xMin; i <= xMax; i++ {
			var k = Coordinate{x: i, y: j}
			var ci int
			var minDistance = math.MaxInt32
			var totalDistance int
			for i := range cs {
				localDistance := distance(cs[i], k)
				totalDistance += localDistance
				if localDistance == minDistance {
					ci = -1
				}
				if localDistance < minDistance {
					minDistance = localDistance
					ci = i
				}
			}
			if totalDistance < 10000 {
				totalArea++
			}

			grid[k] = ci
		}
	}

	counts := make([]int, len(cs))
	max := -1

	infinites := map[int]bool{}
	for y := yMin; y <= yMax; y++ {
		infinites[grid[Coordinate{x: xMin, y: y}]] = true
		infinites[grid[Coordinate{x: xMax, y: y}]] = true
	}
	for x := xMin; x <= xMax; x++ {
		infinites[grid[Coordinate{x: x, y: yMin}]] = true
		infinites[grid[Coordinate{x: x, y: yMax}]] = true
	}

	for _, el := range grid {
		if infinites[el] {
			continue
		}
		counts[el]++
		if counts[el] > max {
			max = counts[el]
		}
	}

	return max, totalArea
}

func distance(a, b Coordinate) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}
