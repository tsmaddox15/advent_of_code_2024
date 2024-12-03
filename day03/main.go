package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("corrupted_memory.txt")
	if err != nil {
		fmt.Printf("Error: %d", err)
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		regex_string := `mul\((\d+),(\d+)\)`
		lines := scanner.Text()
		re := regexp.MustCompile(regex_string)
		matches := re.FindAllStringSubmatch(lines, -1)
		for _, operation := range matches {
			first, err := strconv.Atoi(operation[1])
			if err != nil {
				fmt.Printf("Error: %d", err)
			}

			secound, err := strconv.Atoi(operation[2])
			if err != nil {
				fmt.Printf("Error: %d", err)
			}

			result := first * secound
			sum += result
			fmt.Println(result)
		}
	}
	fmt.Printf("Sum is: %d \n", sum)
}
