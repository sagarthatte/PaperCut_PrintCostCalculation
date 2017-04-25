//declaring this to be the main package
package main

//importing several necessary packages
import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//function to calculate printing costs
func calculatePrintJobCost(costDict map[string]float64, printData []string) float64 {
	var currentPrintJobCost float64
	totalPages, err := strconv.ParseFloat(printData[0], 32)
	if err != nil {
		log.Fatal(err)
	}
	colorPages, err := strconv.ParseFloat(printData[1], 32)
	if err != nil {
		log.Fatal(err)
	}
	bwPages := (totalPages - colorPages)
	isTwoSided, err := strconv.ParseBool(printData[2])

	if isTwoSided == true {
		currentPrintJobCost := (bwPages * costDict["twoSidedBWCost"])
		currentPrintJobCost += (colorPages * costDict["twoSidedColorCost"])
		return currentPrintJobCost
	} else if isTwoSided == false {
		currentPrintJobCost := (bwPages * costDict["oneSidedBWCost"])
		currentPrintJobCost += (colorPages * costDict["oneSidedColorCost"])
		return currentPrintJobCost
	}
	return currentPrintJobCost
}

func main() {
	a4PrintCosts := make(map[string]float64)
	a4PrintCosts["oneSidedBWCost"] = 0.15
	a4PrintCosts["oneSidedColorCost"] = 0.25
	a4PrintCosts["twoSidedBWCost"] = 0.10
	a4PrintCosts["twoSidedColorCost"] = 0.20

	csvHasHeader := true
	totalPrintJobCost := 0.0
	dat, err := ioutil.ReadFile("sample.csv")
	check(err)

	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//logic to calculate and output print job costs
		if csvHasHeader == true {
			csvHasHeader = false
		} else {
			currentPrintJobCost := calculatePrintJobCost(a4PrintCosts, record)
			fmt.Printf("Cost for current print job is: $%.2f \n", currentPrintJobCost)
			totalPrintJobCost += currentPrintJobCost
		}
	}
	fmt.Printf("Total Print Cost is: $%.2f", totalPrintJobCost)
}
