package main

import (
	"bufio"
	"fmt"
	"io"
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
	regex_string := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(regex_string)
	for scanner.Scan() {

		lines := scanner.Text()

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
	fmt.Printf("Sum is for part 1: %d \n", sum)
	defer file.Close()
	//part 2
	_, err = file.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println(err)
	}

	scanner2 := bufio.NewScanner(file)
	sum2 := 0
	//var results []string
	for scanner2.Scan() {
		line := scanner2.Text()
		//new_instructions := regexp.MustCompile(`\bdo\b|\bdon't\b`)
		raw_string := `\b\w*do\w*'?\w*\b`
		new_instructions := regexp.MustCompile(raw_string)
		//parts := new_instructions.Split(line, -1)
		matches := new_instructions.FindAllString(line, -1)
		match_index := new_instructions.FindAllStringIndex(line, -1)
		fmt.Println(match_index)
		fmt.Println(matches)
		for i := 0; i < len(matches); i++ {
			if i == 0 {
				chunk := line[:match_index[i][0]]
				fmt.Printf("first chnk %s \n", chunk)
				total_chunk := calculate_product(chunk)
				sum2 += total_chunk
			}
			if matches[i] != "don't" {
				if i+1 >= len(matches) {
					fmt.Printf("start end spots are %d \n", match_index[i])
					chunk := line[match_index[i][0]:]
					fmt.Printf("Chunk data is: %s \n", chunk)
					total_chunk := calculate_product(chunk)
					sum2 += total_chunk
				} else {
					fmt.Printf("start end spots are %d \n", match_index[i])
					chunk := line[match_index[i][0]:match_index[i+1][0]]
					fmt.Printf("Chunk data is: %s \n", chunk)
					total_chunk := calculate_product(chunk)
					sum2 += total_chunk
				}
			}
		}

	}
	fmt.Printf("Final for part 2 is %d \n", sum)
	fmt.Printf("Final 2 for part 2 is: %d \n", sum2)
}

func calculate_product(chunk string) int {
	regex_string := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(regex_string)
	mul_items := re.FindAllStringSubmatch(chunk, -1)
	total_chunnk := 0
	for _, item := range mul_items {
		value1, err := strconv.Atoi(item[1])
		if err != nil {
			println(err)
		}
		value2, err := strconv.Atoi(item[2])
		if err != nil {
			println(err)
		}
		total_chunnk += value1 * value2
		//fmt.Printf("The product is %d \n", value1*value2)
	}
	return total_chunnk

}
