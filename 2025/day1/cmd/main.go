package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var dialEnd int
	var doorPassword []int
	var password int

	file, err := os.Open("../data/rotations.txt")

	if err != nil {
		fmt.Printf("My Error: %v", err)
	}
	
	scanner := bufio.NewScanner(file)

	for dialStart := 50; scanner.Scan(); dialStart = dialEnd {
		var puzzleInput string = scanner.Text()
		var rotation string = puzzleInput[:1]
		var clicks string = puzzleInput[1:]

		numberOfClicks, err := strconv.Atoi(clicks)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("The dial clicks %d times to the %s.\n", numberOfClicks, rotation)

		if rotation == "R" {
			dialEnd = (dialStart + numberOfClicks) % 100
			if dialEnd == 0 {
				doorPassword = append(doorPassword, dialEnd)
			}
		} else {
			dialEnd = (dialStart - numberOfClicks) % 100
			if dialEnd < 0 {
				dialEnd += 100
			}
			if dialEnd == 0 {
				doorPassword = append(doorPassword, dialEnd)
			}
		}
	}

	password = len(doorPassword)
	fmt.Printf("The password is %d\n", password)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}