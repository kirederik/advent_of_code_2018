package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer1, answer2 := PartI(os.Args[1])
	fmt.Println("Part I: ", answer1)
	fmt.Println("Part I: ", answer2)
}

type InputData struct {
	ID       int
	Datetime time.Time
	Action   string
}

func toInt(b string) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return n
}

func readAll(filepath string) ([]InputData, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var inputData []InputData
	for _, line := range strings.Split(string(data), "\n") {
		var input InputData
		var y, m, d, h, min int
		var action string

		y = toInt(line[1:5])
		m = toInt(line[6:8])
		d = toInt(line[9:11])
		h = toInt(line[12:14])
		min = toInt(line[15:17])
		action = line[19:]
		input.Datetime = time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
		input.Action = action

		inputData = append(inputData, input)
	}

	return inputData, nil
}

func PartI(filepath string) (int, int) {
	data, err := readAll(filepath)
	if err != nil {
		panic(err)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Datetime.Before(data[j].Datetime)
	})

	m := make(map[string]map[int]int)
	var lastGuardID string
	var sleepStarted int
	for _, v := range data {
		if strings.Contains(v.Action, "begins shift") {
			lastGuardID = strings.TrimLeft(strings.Split(v.Action, " ")[1], "#")
			continue
		}

		if strings.Contains(v.Action, "falls asleep") {
			sleepStarted = v.Datetime.Minute()
			continue
		}

		awakeAt := v.Datetime.Minute()
		if m[lastGuardID] == nil {
			m[lastGuardID] = make(map[int]int)
		}
		for i := sleepStarted; i < awakeAt; i++ {
			m[lastGuardID][i]++
		}
	}

	var maxSleeper string
	var maxMinutesSleeped int
	var topMinSleeped int

	var frequentSleeper string
	var frequentMinute int
	var mostSleepedMinute int

	for k, v := range m {
		var sumMinutesSleeped int
		var localMaxMinutesSleeped int
		var localMinuteSleepedMost int

		for min, timesSleeped := range v {
			sumMinutesSleeped += timesSleeped
			if timesSleeped > localMaxMinutesSleeped {
				localMaxMinutesSleeped = timesSleeped
				localMinuteSleepedMost = min
			}

			if timesSleeped > mostSleepedMinute {
				mostSleepedMinute = timesSleeped
				frequentMinute = min
				frequentSleeper = k
			}
		}

		if sumMinutesSleeped > maxMinutesSleeped {
			maxMinutesSleeped = sumMinutesSleeped
			maxSleeper = k
			topMinSleeped = localMinuteSleepedMost
		}
	}

	return toInt(maxSleeper) * topMinSleeped, toInt(frequentSleeper) * frequentMinute
}
