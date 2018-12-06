package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer1 := PartI(os.Args[1])
	fmt.Println("Part I: ", answer1)

	answer2 := PartII(os.Args[1])
	fmt.Println("Part II: ", answer2)
}

func readAll(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PartI(filepath string) int {
	data, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	return removeCollisions(data, 0, len(data))
}

func PartII(filepath string) int {
	data, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	m := make(map[string]int)

	min := removeCollisions([]byte(string(data)), 0, len(data))
	for _, v := range data {
		lowerV := strings.ToLower(string(v))
		if _, found := m[lowerV]; found {
			continue
		}

		m[lowerV]++

		modifiedStr := strings.Replace(string(data), lowerV, "", -1)
		modifiedStr = strings.Replace(modifiedStr, strings.ToUpper(lowerV), "", -1)

		r := removeCollisions([]byte(modifiedStr), 0, len(modifiedStr))
		if r < min {
			min = r
		}

	}
	return min
}

func removeCollisions(str []byte, start, end int) int {
	if start < 0 {
		start = 1
	}
	for i := start + 1; i < end; i++ {
		transformation := 32
		if str[i] < 'a' {
			transformation = -32
		}

		if str[i-1]+byte(transformation) == str[i] {
			newData := []byte(str[:i-1])
			newData = append(newData, str[i+1:]...)
			return removeCollisions(newData, i-2, len(newData))
		}
	}

	return len(str)

}
