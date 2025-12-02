package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var solutionCounter int = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func moveDial(position int, clicks int, direction int) int {
	newPosition := position
	loop0counter := 0
	for i := 0; i < clicks; i++ {
		newPosition += direction
		//fmt.Println("clicks to ", newPosition)
		if newPosition == 0 {
			// We point at zero
			solutionCounter++
			loop0counter++
		}
		if newPosition < 0 {
			newPosition = 99
		} else if newPosition > 99 {
			newPosition = 0
			// We point at zero in this case too
			solutionCounter++
			loop0counter++
		}
	}
	fmt.Println("to point at", newPosition)
	if loop0counter > 0 {
		fmt.Println("during this rotation, it points at `0` N times, N being", loop0counter)
	}
	return newPosition
}

func main() {
	lines := 0
	dialPos := 50
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	fmt.Println("The dial is starting at ", dialPos)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		fmt.Println("The dial is rotating", line)
		rightWay := strings.Split(line, "R")
		if len(rightWay) > 1 {
			clicks, err := strconv.Atoi(rightWay[1])
			check(err)
			dialPos = moveDial(dialPos, clicks, +1)
		} else {
			leftWay := strings.Split(line, "L")
			clicks, err := strconv.Atoi(leftWay[1])
			check(err)
			dialPos = moveDial(dialPos, clicks, -1)
		}
	}

	fmt.Println("The dial has been at 0 for ", solutionCounter)
}
