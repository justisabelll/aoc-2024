package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)


func splitInput(i string) ([]int, map[int]int) {
	a := strings.Fields(i)
	var leftArr []int
	var rightMap = make(map[int]int)

	for i, e := range a {
		temp := strings.TrimSpace(e)
		res, err := strconv.Atoi(temp)

		if err != nil {
			log.Fatal(err)
		}
		if i%2 == 0 {
			rightMap[res] = 0

		} else {
			leftArr = append(leftArr, res)
		}
	}

	return leftArr, rightMap
	}


func countDuplicates(m map[int]int, arr []int) map[int]int {
	for _, e := range arr {
		_, exists := m[e] 

		if exists {
			m[e] = m[e] + 1
		}
	}

	return m
}

func calSimilarityScore(m map[int]int) int {
	total := 0
	
	for k, v := range m {
		total += k * v 
	}


	return total
}


func main(){

	input := utils.GetInput(true)

	left, rightMap := splitInput(input)
	rightMap = countDuplicates(rightMap, left)

	total := calSimilarityScore(rightMap)

	fmt.Println(total)
}
