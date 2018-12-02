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

	answer, err := CheckSum(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Part I: ", answer)

	answer2, err := FindCommonLetters(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Part II: ", answer2)
}

func readAll(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	for {
		var s string
		_, err := fmt.Fscanln(f, &s)
		if err == io.EOF {
			return lines, nil
		}
		if err != nil {
			return nil, err
		}

		lines = append(lines, s)
	}
}

func calcCheckSum(l []map[rune]int) int {
	var twoCount, threeCount int

	for _, m := range l {
		var foundTwo, foundThree bool
		for _, v := range m {
			if v == 2 && !foundTwo {
				twoCount++
				foundTwo = true
			}
			if v == 3 && !foundThree {
				threeCount++
				foundThree = true
			}
		}
	}

	return twoCount * threeCount
}

func CheckSum(filepath string) (int, error) {
	data, err := readAll(filepath)
	if err != nil {
		return 0, err
	}

	var l []map[rune]int
	for _, str := range data {
		letters := make(map[rune]int)
		for _, c := range []rune(str) {
			letters[c]++
		}

		l = append(l, letters)
	}

	return calcCheckSum(l), nil
}

func diffOne(s, s2 string) (bool, int) {
	if len(s) != len(s2) {
		return false, -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] != s2[i] {
			return s[i+1:] == s2[i+1:], i
		}
	}
	return false, -1
}

func FindCommonLetters(filepath string) (string, error) {
	data, err := readAll(filepath)
	if err != nil {
		return "", err
	}

	var d bool
	var diffAt int
	var r string

	for i, s := range data {
		if r != "" {
			break
		}
		for j, s2 := range data {
			if i != j {
				d, diffAt = diffOne(s, s2)
				if d {
					r = s
					break
				}
			}
		}
	}

	return r[:diffAt] + r[diffAt+1:], nil
}
