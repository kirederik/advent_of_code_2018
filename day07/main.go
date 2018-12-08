package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer1 := PartI(os.Args[1])
	fmt.Println("Part I: ", answer1)
	PartII(os.Args[1])
}

type Step struct {
	value    byte
	dependOn []byte
}

func readAll(filepath string) (map[byte]*Step, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	deps := map[byte]*Step{}

	for {
		var a, b byte
		_, err := fmt.Fscanf(f, "Step %c must be finished before step %c can begin.\n", &a, &b)

		if err != nil {
			break
		}

		if deps[a] == nil {
			deps[a] = &Step{value: a}
		}
		if deps[b] == nil {
			deps[b] = &Step{value: b}
		}

		deps[b].dependOn = append(deps[b].dependOn, a)
	}

	return deps, nil
}

func PartI(filepath string) string {
	deps, _ := readAll(filepath)

	var order string

	for len(deps) > 0 {
		var s *Step
		for _, node := range deps {
			if len(node.dependOn) == 0 && (s == nil || s.value > node.value) {
				s = node
			}
		}
		delete(deps, s.value)
		order += string(s.value)

		for _, node := range deps {
			for i := 0; i < len(node.dependOn); i++ {
				if node.dependOn[i] == s.value {
					l := len(node.dependOn)
					node.dependOn[l-1], node.dependOn[i] = node.dependOn[i], node.dependOn[l-1]
					node.dependOn = node.dependOn[:l-1]
				}
			}
		}
	}

	return order
}

type Work struct {
	value byte
	time  int
}

func PartII(filepath string) {
	order := PartI(filepath)
	deps, _ := readAll(filepath)

	var totalTime int

	completed := map[byte]bool{}
	inProgress := map[byte]*Work{}

	for len(completed) < len(order) {
		for k, v := range inProgress {
			if v.time == totalTime {
				completed[k] = true
				delete(inProgress, k)
				for i := range deps {
					for j, dd := range deps[i].dependOn {
						if dd == k {
							deps[i].dependOn = append(deps[i].dependOn[:j], deps[i].dependOn[j+1:]...)
						}
					}
				}
			}
		}

		var potential []byte
		for _, s := range deps {
			if inProgress[s.value] != nil || completed[s.value] {
				continue
			}
			if len(s.dependOn) == 0 {
				potential = append(potential, s.value)
			}
		}

		if len(inProgress) < 5 && len(potential) > 0 {
			for i := 0; i < 5 && i < len(potential); i++ {
				s := deps[potential[i]]
				inProgress[s.value] = &Work{value: s.value, time: totalTime + 61 + int(s.value-'A')}
			}
		}

		totalTime++
	}

	fmt.Printf("totalTime = %+v\n", totalTime-1)
}
