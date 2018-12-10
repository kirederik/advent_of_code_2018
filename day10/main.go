package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Light struct {
	position Point
	speedX   int
	speedY   int
}

type Point struct {
	x, y int
}

type Grid struct {
	lights [][]string
}

func main() {
	PartI("input")
}

func readAll(filepath string) []Light {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil
	}

	lights := []Light{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range strings.Split(string(f), "\n") {
		nums := re.FindAllStringSubmatch(string(line), -1)

		var x, y, r, d int
		x, _ = strconv.Atoi(nums[0][0])
		y, _ = strconv.Atoi(nums[1][0])
		r, _ = strconv.Atoi(nums[2][0])
		d, _ = strconv.Atoi(nums[3][0])

		lights = append(lights, Light{position: Point{x: x, y: y}, speedX: r, speedY: d})
	}

	return lights
}

const boxSize = 100

func PartI(filepath string) {
	lights := readAll(filepath)
	if lights == nil {
		panic("empty lights")
	}

	for i := 0; i < 100000; i++ {
		var minX, minY, maxX, maxY int
		minX, minY = lights[0].position.x, lights[0].position.y
		for _, l := range lights {
			if l.position.x < minX {
				minX = l.position.x
			}

			if l.position.y < minY {
				minY = l.position.y
			}

			if l.position.x > maxX {
				maxX = l.position.x
			}

			if l.position.y > maxY {
				maxY = l.position.y
			}
		}

		if minX+boxSize >= maxX && minY+boxSize >= maxY {
			fmt.Printf("i = %+v\n", i)
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if isOnAt(x, y, lights) {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Println("")
			}
		}

		for i := range lights {
			lights[i].position.x += lights[i].speedX
			lights[i].position.y += lights[i].speedY
		}
	}
}

func isOnAt(x, y int, lights []Light) bool {
	for _, l := range lights {
		if l.position.x == x && l.position.y == y {
			return true
		}
	}
	return false
}
