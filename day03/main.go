package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer := FindOverlaps(os.Args[1])
	fmt.Println("Part I: ", answer)

	answer2 := FindNonOverlapping(os.Args[1])
	fmt.Println("Part II: ", answer2)
}

type InputData struct {
	ID                    string
	LeftOffset, TopOffset int
	Width, Height         int
}

func readAll(filepath string) ([]InputData, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var inputData []InputData
	for {
		var input InputData
		_, err := fmt.Fscanf(
			f,
			"%s @ %d,%d: %dx%d",
			&input.ID,
			&input.LeftOffset,
			&input.TopOffset,
			&input.Width,
			&input.Height,
		)

		if err == io.EOF {
			return inputData, nil
		}
		if err != nil {
			return nil, err
		}

		inputData = append(inputData, input)
	}
}

func offsetArea(i InputData) int {
	return i.LeftOffset * i.TopOffset
}

func FindOverlaps(filepath string) int {
	claims, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	m := make(map[string]int)

	for _, c := range claims {
		for i := c.LeftOffset; i < c.LeftOffset+c.Width; i++ {
			for j := c.TopOffset; j < c.TopOffset+c.Height; j++ {
				m[fmt.Sprintf("%d,%d", i, j)]++
			}
		}
	}

	var overlapping int
	for _, v := range m {
		if v > 1 {
			overlapping++
		}
	}
	return overlapping
}

func FindNonOverlapping(filepath string) string {
	claims, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	for i, c1 := range claims {
		var foundOverlapping bool

		for j, c2 := range claims {
			if i == j {
				continue
			}
			if overlaps(c1, c2) {
				foundOverlapping = true
				break
			}
		}
		if foundOverlapping {
			continue
		}

		return c1.ID
	}
	return ""
}

func overlaps(i, j InputData) bool {
	left := i.LeftOffset+i.Width <= j.LeftOffset
	top := i.TopOffset+i.Height <= j.TopOffset
	right := i.LeftOffset >= j.LeftOffset+j.Width
	bottom := i.TopOffset >= j.TopOffset+j.Height

	return !(left || top || right || bottom)
}
