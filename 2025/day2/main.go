package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	
	var match bool = false
	// var matchedIds []string
	
	file, err := os.Open("./data/input.csv")

	if err != nil {
		log.Fatalf("Error trying to open the file: %v", err)
	}

	defer file.Close()
	
	csvReader := csv.NewReader(file)
	
	idRanges, err := csvReader.Read()

		if err != nil {
			log.Fatalf("Coudn't read id ranges. Error: %v", err)
		}

	for _, idRange := range idRanges{
		fmt.Printf("Id range: %s\n", idRange)
		rangeBoundaries := strings.Split(idRange, "-")
		
		rangeStart, err := strconv.Atoi(rangeBoundaries[0])
		if err != nil {
			fmt.Printf("Error converting the rangeStart %d\n", rangeStart)
		}

		rangeEnd, err := strconv.Atoi(rangeBoundaries[1])
		if err != nil {
			fmt.Printf("Error converting the rangeEnd %d\n", rangeEnd)
		}

		
		for rangeCount := rangeStart; rangeCount <= rangeEnd; rangeCount++ {
			fmt.Printf("rangeCount is now %d\n", rangeCount)
		}
		

		// if there is a match, set match := true
		if match {
			// matchedIds := append(matchedIds, idRange)
		}
	}

	// take all strings in matchedIds, convert to int
	// add up all the invalid ids for answer.

	
}
