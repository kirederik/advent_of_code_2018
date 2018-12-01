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

	answer, err := SumSequence(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Part I: ", answer)

	answer, err = FindLoop(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Part II: ", answer)
}

func SumSequence(filepath string) (int64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	var sum int64
	for {
		var i int64
		_, err := fmt.Fscanln(f, &i)
		if err != nil {
			if err == io.EOF {
				return sum, nil
			}
			return 0, err
		}
		sum += i
	}
}

func FindLoop(filepath string) (int64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	var sum int64

	m := map[int64]int{0: 1}
	for {
		var i int64
		_, err := fmt.Fscanln(f, &i)

		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			f.Seek(0, 0)
			continue
		}

		sum += i
		m[sum]++
		if m[sum] == 2 {
			return sum, nil
		}
	}
}
