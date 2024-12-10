package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getSafeCount(fileName string) int {
	numSafe := 0
	// Read the input line by line
	file, err := os.Open(fmt.Sprintf("./%s", fileName))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		report := parseLine(line)
		reportValue := 1
		isIncreasing := true
		if report[0]-report[1] > 0 {
			isIncreasing = false
		}
		for i := 1; i < len(report); i++ {
			if isUnsafe(report[i-1], report[i], isIncreasing) {
				reportValue = 0
				break
			}
		}
		numSafe += reportValue
	}
	return numSafe
}

func getSafeCountDampener(fileName string) int {
	// keep counter of failures starting at 1 for each level
	// run the same code as for problem1
	// if two numbers don't meet the safety criteria and failures > 0, run the current number and the one after the next, else return unsafe
	// if those two do, decrement failures to 0
	// if those two don't, return unsafe
	numSafe := 0
	// Read the input line by line
	file, err := os.Open(fmt.Sprintf("./%s", fileName))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dampener := 1
		line := strings.Fields(scanner.Text())
		report := parseLine(line)
		reportValue := 1
		isIncreasing := true
		if report[0]-report[1] > 0 {
			isIncreasing = false
		}
		for i := 1; i < len(report); i++ {
			isUnsafePair := isUnsafe(report[i-1], report[i], isIncreasing)
			if isUnsafePair && dampener == 1 {
				if i < len(report)-1 && isUnsafe(report[i-1], report[i+1], isIncreasing) {
					reportValue = 0
					break
				} else {
					dampener--
				}
			} else if isUnsafePair && dampener == 0 {
				reportValue = 0
				break
			}
		}
		numSafe += reportValue
	}
	return numSafe
}

func isUnsafe(first, second int, isIncreasing bool) bool {
	diff := first - second
	absDiff := int(math.Abs(float64(diff)))
	return (absDiff < 1 || absDiff > 3) || (diff > 0 && isIncreasing) || (diff < 0 && !isIncreasing)
}

func parseLine(line []string) []int {
	res := []int{}
	for _, item := range line {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}
