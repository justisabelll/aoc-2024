package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)


func cleanInput(i string) ([]int, []int) {
	a := strings.Fields(i)
	var leftArr []int
	var rightArr []int

	for i, e := range a {
		temp := strings.TrimSpace(e)
		res, err := strconv.Atoi(temp)

		if err != nil {
			log.Fatal(err)
		}
		if i%2 == 0 {
			rightArr = append(rightArr, res)
		} else {
			leftArr = append(leftArr, res)
		}
	}
	return leftArr, rightArr
}

func getInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)

}

func main() {

	input := getInput()

	left, right := cleanInput(input)

	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i, e := range left {
		total += int(math.Abs(float64(e - right[i])))
	}

	fmt.Println(total)
}
