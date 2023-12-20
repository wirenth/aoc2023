package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	calValues := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		digits := make([]string, 0)
		line := scanner.Text()
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
		}
		strValue := fmt.Sprintf("%v%v", digits[0], digits[len(digits)-1])
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			log.Fatal(err)
		}

		calValues += intValue
	}
	fmt.Println(calValues)
}
