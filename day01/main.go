package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//os.Stdout.Write(data)
	row := make([]string, 0)
	row1 := make([]string, 0)
	scanner := bufio.NewScanner((data))
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		row = append(row, line[0])
		row1 = append(row1, line[1])

	}

	int_row := make([]int, len(row))

	for i, str := range row {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("hahahaha")
		}
		int_row[i] = num
	}

	int_row1 := make([]int, len(row1))

	for i, str := range row1 {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("hahahaha")
		}
		int_row1[i] = num
	}

	sort.Ints(int_row)
	fmt.Println(int_row)

	sort.Ints(int_row1)
	fmt.Println(int_row1)
	difference_array := make([]int, 0)
	if len(int_row) == len(int_row1) {
		for i := 0; i < len(int_row); i++ {
			var difference int
			if int_row[i] > int_row1[i] {
				difference = int_row[i] - int_row1[i]
				fmt.Printf("item1 : %d, item 2 %d difference: %d\n", int_row[i], int_row1[i], difference)
			} else if int_row1[i] > int_row[i] {
				difference = int_row1[i] - int_row[i]
				fmt.Printf("item1 : %d, item 2 %d difference: %d\n", int_row[i], int_row1[i], difference)
			}
			difference_array = append(difference_array, difference)
		}
	}
	sum := 0
	for _, num := range difference_array {
		sum += num
	}

	//Part 2
	similarity_sum := 0
	for _, row_value := range int_row {
		similarity_count := 0
		for _, row_value1 := range int_row1 {
			if row_value == row_value1 {
				similarity_count += 1
				fmt.Printf("item1 : %d, item 2 %d are the same\n", row_value, row_value1)
			}

		}
		similarity_sum += row_value * similarity_count
	}

	fmt.Printf("Total distance is %d\n", sum)
	fmt.Printf("The total similarity score is: %d\n", similarity_sum)
}
