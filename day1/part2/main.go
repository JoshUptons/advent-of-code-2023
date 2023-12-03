package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var (
	digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
)

/*
will check the first chunk of a string heap for a digit
this does not check for digits from actual integers, that
will be handled by the isDigit function
*/
func hasDigit(p string) (bool, int) {
	// would result in index out of range error
	if len(p) < 3 {
		return false, 0
	}
	for r := range p {
		if isDigit(rune(r)) {
			return false, 0
		}
	}
	for key := range digits {
		length := len(key)
		if len(p) < length {
			continue
		}
		if p[:length] == key {
			return true, digits[key]
		}
	}
	return false, 0
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isDigit(r rune) bool {
	if ok := unicode.IsDigit(r); ok {
		return true
	}
	return false
}

/*
I am going to need to read digits still, but also read the line for numbers
in string form.
*/
func stringToInts(s string) []int {
	var ints []int
	for i := 0; i < len(s); i++ {
		if isDigit(rune(s[i])) {
			i, err := strconv.Atoi(string(s[i]))
			if err != nil {
				fmt.Println(err)
			}
			ints = append(ints, i)
		}
		var chunk string
		if i+5 > len(s) {
			chunk = s[i:]
		} else {
			chunk = s[i : i+5]
		}
		if ok, val := hasDigit(chunk); ok {
			ints = append(ints, val)
		}
	}
	return ints
}

func main() {
	var total int
	file, err := os.Open("../input.txt")
	handleErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// parse line into bytes
	for scanner.Scan() {
		line := scanner.Text()
		ints := stringToInts(line)
		n, err := strconv.Atoi(fmt.Sprintf("%d%d", ints[0], ints[len(ints)-1]))
		handleErr(err)
		total += n
	}

	if scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)

}
