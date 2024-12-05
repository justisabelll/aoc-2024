package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)


func splitInput(i string) ([]int, []int) {
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

func main() {
	input := utils.GetInput(false)

	left, right := splitInput(input)

	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i, e := range left {
		total += int(math.Abs(float64(e - right[i])))
	}

	fmt.Println(total)
}
