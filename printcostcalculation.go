//declaring this as the main package
package main

//importing necessary packages
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

/*unction to calculate printing costs
 *	Input: Takes the map (for printing cost data) and printData (individual records) to calculate cost
 *  Output: Calculates and returns the cost for the specific print job.
 */
func calculatePrintJobCost(costDict map[string]float64, printData []string) float64 {
	var currentPrintJobCost float64
	totalPages, err := strconv.ParseFloat(strings.Replace(printData[0], " ", "", -1), 32)
	if err != nil {
		log.Fatal(err)
	}
	colorPages, err := strconv.ParseFloat(strings.Replace(printData[1], " ", "", -1), 32)
	if err != nil {
		log.Fatal(err)
	}
	bwPages := (totalPages - colorPages)
	isTwoSided, err := strconv.ParseBool(strings.Replace(printData[2], " ", "", -1))

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

//Main Function
func main() {
	a4PrintCosts := make(map[string]float64) //Map stores printing cost for each condition.
	a4PrintCosts["oneSidedBWCost"] = 0.15
	a4PrintCosts["oneSidedColorCost"] = 0.25
	a4PrintCosts["twoSidedBWCost"] = 0.10
	a4PrintCosts["twoSidedColorCost"] = 0.20

	csvHasHeader := true                      //this variable indicates that csvFile has a header
	currJobNum := 1                           //intialize current job number
	var totalPrintJobCost float64             //initializing total printing cost
	dat, err := ioutil.ReadFile("sample.csv") //reading csv file
	check(err)                                //checking for errors

	r := csv.NewReader(strings.NewReader(string(dat)))

	//looping through csv file for individual records
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if csvHasHeader == true { //skips the first line, if csv file has header
			csvHasHeader = false
		} else { //calls cost calculation function for each record.
			currentPrintJobCost := calculatePrintJobCost(a4PrintCosts, record)
			fmt.Printf("Cost for print job %d is: $%.2f \n", currJobNum, currentPrintJobCost)
			currJobNum += 1
			totalPrintJobCost += currentPrintJobCost //adding costs to the total cost
		}
	}
	fmt.Printf("Total Printing Cost is: $%.2f", totalPrintJobCost) //outputting total cost to console/terminal
}
