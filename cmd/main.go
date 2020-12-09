package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	sort.Ints(result)
	return result, scanner.Err()
}

// Check to see if any two of these numbers sum together to make 2020
func FindSum(targetSum int, intSlice []int) (int, int) {
	for i := range intSlice {
		justRight := 2020 - intSlice[i]
		for j := len(intSlice) - 1; j > 0; j-- {
			if intSlice[j] == justRight {
				return intSlice[i], intSlice[j]
			}
		}
	}
	return 0, 0
}

func main() {
	log.Println(os.Environ())
	log.Println(os.Getwd())
	file, err := os.Open("../inputs/day1/input.txt")
	if err != nil {
		log.Fatal("no open file", err)
	}
	defer file.Close()
	ints, err := ReadInts(file)
	if err != nil {
		log.Fatal("err reading the file into ints", err)
	}
	one, two := FindSum(int(2020), ints)
	log.Println(one, two)
	log.Println(one * two)
	// for i := range ints {
	//	log.Println(ints[i])
	// }
}
