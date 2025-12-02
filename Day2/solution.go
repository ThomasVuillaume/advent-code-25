package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sol int = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addToSol(num int) {
	sol += num
}

func allSameStrings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func getSlices(s string) {
	l := len(s)
	//res := [][]string{}
	mod := l % 2

	if mod == 0 {
		// fmt.Println("MOD0 - String sliced :", s)
		for i := 1; i <= l/2; i++ {
			// fmt.Println("Packet size", i)
			res := []string{}
			if l%i == 0 {
				for j := 0; (i + j) <= l; {
					// fmt.Println("Indexes", j, j+i)
					// fmt.Println("slice", s[j:j+i])
					res = append(res, s[j:j+i])
					j += i
					// fmt.Println("res", res)
				}
				if allSameStrings(res) {
					fmt.Println("Invalid ID", s)
					n, err := strconv.Atoi(s)
					check(err)
					addToSol(n)
					res = nil
					break
				}
				//fmt.Println("Result", res)
			}
		}
	} else {
		// fmt.Println("MOD NOT 0", s)
		for i := 1; i <= (l+1)/2; i++ {
			// fmt.Println("Packet size", i)
			res := []string{}
			if l%i == 0 {
				for j := 0; (i + j) <= l; {
					// fmt.Println("Indexes", j, j+i)
					// fmt.Println("slice", s[j:j+i])
					res = append(res, s[j:j+i])
					j += i
					// fmt.Println("res", res)
				}
				if allSameStrings(res) {
					fmt.Println("Invalid ID", s)
					n, err := strconv.Atoi(s)
					check(err)
					addToSol(n)
					res = nil
					break
				}
				//fmt.Println("Result", res)
			}
		}
	}
}

func isMirorNum(num int) {
	stringValue := strconv.Itoa(num)
	getSlices(stringValue)
}

func analyseRange(stringRange string) {
	minMax := strings.Split(stringRange, "-")
	min, err := strconv.Atoi(minMax[0])
	check(err)
	max, err := strconv.Atoi(minMax[1])
	check(err)
	for i := min; i <= max; i++ {
		//fmt.Println("Analysing", i)
		isMirorNum(i)
	}
}

func main() {
	// Vars
	lines := 0

	// File read
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var stringRanges []string
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		stringRanges = strings.Split(line, ",")
	}

	for _, stringRange := range stringRanges {
		//fmt.Println("Analysing range", stringRange)
		analyseRange(stringRange)
	}

	fmt.Println("Solution", sol)

}
