package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var total int
	file, err := os.Open("../input.txt")
	handleErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// parse line into bytes
	for scanner.Scan() {
		var digits []int
		lineBytes := scanner.Bytes()
		// parse bytes into characters
		for _, b := range lineBytes {
			if digit, err := strconv.Atoi(string(b)); err == nil {
				// is a digit
				digits = append(digits, digit)
			}
		}
		// combine the numbers to add them to the total
		num, err := strconv.Atoi(
			fmt.Sprintf(
				"%d%d",
				digits[0],
				digits[len(digits)-1],
			),
		)
		handleErr(err)
		// finally add them to the total
		total += num
	}

	if scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)

}
