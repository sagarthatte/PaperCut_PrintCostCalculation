//declaring this to be the main package
package main

//importing several necessary packages
import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

//defining printing costs as constant values
const (
	oneSidedBWCost    float64 = 0.15
	oneSidedColorCost float64 = 0.25
	twoSidedBWCost    float64 = 0.10
	twoSidedColorCost float64 = 0.20
	//csvHasHeader      bool    = true
)

//error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.55, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
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
			totalPages, err := strconv.ParseFloat(record[0], 32)
			if err != nil {
				log.Fatal(err)
			}
			colorPages, err := strconv.ParseFloat(record[1], 32)
			if err != nil {
				log.Fatal(err)
			}
			bwPages := (totalPages - colorPages)
			isTwoSided, err := strconv.ParseBool(record[2])

			if isTwoSided == true {
				currentPrintJobCost := (bwPages * twoSidedBWCost)
				currentPrintJobCost += (colorPages * twoSidedColorCost)
				fmt.Printf("Cost for current print job is: $%.2f \n", currentPrintJobCost)
				totalPrintJobCost += currentPrintJobCost
			} else if isTwoSided == false {
				currentPrintJobCost := (bwPages * oneSidedBWCost)
				currentPrintJobCost += (colorPages * oneSidedColorCost)
				fmt.Printf("Cost for current print job is: $%.2f \n", currentPrintJobCost)
				totalPrintJobCost += currentPrintJobCost
			}
		}
	}
	fmt.Printf("Total Print Cost is: $%.2f", totalPrintJobCost)
}
