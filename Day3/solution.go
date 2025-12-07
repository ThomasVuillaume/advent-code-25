package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var sol int = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMaxJoltage(bank string) {
	fmt.Println("Bank analysed", bank)
	l := len(bank)
	decimalSlice := strings.Split(bank[:l-1], "")
	//fmt.Println("Decimal Slice analysed", decimalSlice)

	decimal := slices.Max(decimalSlice)
	indexOfDecimal := slices.IndexFunc(decimalSlice, func(d string) bool { return d == decimal })
	// fmt.Println("Joltage Decimal / Index = ", decimal, indexOfDecimal)

	// fmt.Println("Unit Slice Search Starts At", indexOfDecimal+1)
	unitSlice := strings.Split(bank[indexOfDecimal+1:], "")
	unit := slices.Max(unitSlice)
	//fmt.Println("Joltage", decimal, unit)

	joltage, err := strconv.Atoi(decimal + unit)
	check(err)
	// fmt.Println("Joltage", joltage)
	sol += joltage
}

func main() {
	// Vars

	// File read
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		banks := scanner.Text()
		getMaxJoltage(banks)
	}

	fmt.Println("Solution", sol)

}
