package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testCase struct {
	Input   []string
	Output  int
	Output2 int
}

func TestPartI(t *testing.T) {
	testCases := []testCase{
		{
			Input: []string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			},
			Output:  240,
			Output2: 4455,
		},
	}

	var errors []string
	for i, v := range testCases {
		f, err := createFileWithContent(v.Input)
		if err != nil {
			t.Error("Could not create file for input", err)
		}

		r, r2 := PartI(f.Name())
		if r != v.Output {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output, i, r))
		}

		if r2 != v.Output2 {
			errors = append(errors, fmt.Sprintf("Expecting %d for test case %d. Got %d", v.Output2, i, r2))
		}

		os.Remove(f.Name()) // clean up
	}
	if len(errors) > 0 {
		t.Error("\n" + strings.Join(errors, "\n"))
	}
}

func createFileWithContent(s []string) (*os.File, error) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.WriteString(strings.Join(s, "\n")); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}
