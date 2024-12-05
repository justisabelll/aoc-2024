package utils

import (
	"log"
	"os"
)

func GetInput(isBonus bool) string {

	var path string = "input.txt"

	if isBonus {
		path = "../input.txt"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
