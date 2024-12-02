package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %d \n", err)
	}

	scanner := bufio.NewScanner(data)
	report_count := 0
	for scanner.Scan() {
		line_items := strings.Fields(scanner.Text())
		values := make([]int, 0)
		for _, item := range line_items {
			int_item, err := strconv.Atoi(item)
			if err != nil {
				fmt.Printf("Error: %d \n", err)
				continue
			}
			values = append(values, int_item)
		}
		var increaseing bool
		valid_report := true
		if values[1] > values[0] {
			increaseing = true
		} else {
			increaseing = false
		}
		for i := 1; i < len(values); i++ {
			prev := values[i-1]
			current := values[i]
			difference := math.Abs(float64(prev - current))
			if difference > 3 {
				fmt.Printf("The report %d faild due to the increase/decrease being greater than 3. For the values of %d and %d \n", values, prev, current)
				valid_report = false
			}
			if prev == current {
				fmt.Printf("The report %d is not valid due to %d and %d not being an increase/decrease \n", values, prev, current)
				valid_report = false
			}
			if increaseing && prev > current {
				fmt.Printf("The report %d failed to increase when a increaseing report. Went from %d to %d \n", values, prev, current)
				valid_report = false
			}
			if !increaseing && prev < current {
				fmt.Printf("The report %d Failed to decrease when a decreasing report. Went from %d to %d \n", values, prev, current)
				valid_report = false
			}
			//fmt.Printf("Previous value: %d , current value: %d \n", values[i-1], values[i])
		}
		if valid_report {
			fmt.Printf("The report %d is valid! \n", values)
			report_count += 1
		}
	}
	fmt.Printf("Total valid reports: %d \n", report_count)
}
