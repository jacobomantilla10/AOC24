package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func prob1(testFile string) int {
	arr1, arr2 := getItems(testFile)
	sort(arr1, 0, len(arr1)-1)
	sort(arr2, 0, len(arr2)-1)

	res := 0
	for i := 0; i < len(arr1); i++ {
		res += int(math.Abs(float64(arr1[i]) - float64(arr2[i])))
	}
	return res
}

func prob2(testFile string) int {
	arr1, arr2 := getItems(testFile)
	sort(arr2, 0, len(arr2)-1)

	// loop through all of the elements in arr1 finding the amount of times they occur in arr2
	res := 0
	for _, num := range arr1 {
		appearances := occurrences(num, arr2)
		res += num * appearances
	}
	return res
}

func getItems(testFile string) ([]int, []int) {
	arr1 := []int{}
	arr2 := []int{}
	file, err := os.Open(fmt.Sprintf("./%s", testFile))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Fields(scanner.Text())
		item1, _ := strconv.Atoi(items[0])
		item2, _ := strconv.Atoi(items[1])
		arr1 = append(arr1, item1)
		arr2 = append(arr2, item2)
	}
	return arr1, arr2
}

// Re
func occurrences(target int, arr []int) int {
	firstIdx := search(arr, 0, len(arr)-1, target)

	if firstIdx == -1 {
		return 0
	}

	count := 1
	for i := firstIdx + 1; i < len(arr) && arr[i] == target; i++ {
		count++
	}
	for j := firstIdx - 1; j >= 0 && arr[j] == target; j-- {
		count++
	}

	return count
}
