package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/random9s/holiday/holiday"
)

func main() {
	year := 2022
	argLen := len(os.Args)
	if argLen > 1 {
		parsedYear, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		year = parsedYear
	}

	fmt.Printf("Generating %d holidays based on New York timezone\n", year)
	fmt.Println("---------------------------------------------------------------------------")
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}

	holidays := holiday.Generate(ny, year)
	for h, date := range holidays {
		fmt.Printf("%s falls on %d-%d-%d\n", h, date.Year(), date.Month(), date.Day())
	}

	fmt.Println()
	fmt.Println()

	fmt.Printf("Generating %d holidays based on Toronto timezone\n", year)
	fmt.Println("---------------------------------------------------------------------------")
	toronto, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Fatal(err)
	}

	holidays = holiday.Generate(toronto, year)
	for h, date := range holidays {
		fmt.Printf("%s falls on %d-%d-%d\n", h, date.Year(), date.Month(), date.Day())
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Automated Testing 2038 holidays based on New York timezone")
	fmt.Println("---------------------------------------------------------------------------")
	holidays = holiday.Generate(ny, 2038)
	for h, date := range holidays {
		var expectedDate time.Time

		switch h {
		case holiday.GregorianEaster:
			expectedDate = time.Date(2038, time.April, 26, 0, 0, 0, 0, time.Now().Location())
		case holiday.LaborDay:
			expectedDate = time.Date(2038, time.September, 6, 0, 0, 0, 0, time.Now().Location())
		case holiday.ThanksgivingDay:
			expectedDate = time.Date(2038, time.November, 25, 0, 0, 0, 0, time.Now().Location())
		default:
			continue
		}

		fmt.Printf("%s falls on %d-%d-%d and the generated date is %d-%d-%d\n", h, date.Year(), date.Month(), date.Day(), expectedDate.Year(), expectedDate.Month(), expectedDate.Day())
	}
}
