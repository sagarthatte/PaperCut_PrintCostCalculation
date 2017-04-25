""" Code to read print orders from sample.csv file and output printing costs
    Input: CSV Data for Print jobs from .csv file
    Output: Printing cost for individual print jobs and total cost of printing
"""

#import module to open and read CSV files.
import csv

#create dictionary for print costs. Can be expanded to include other cost types.
a4PrintCosts = {'oneSidedBWCost': 0.15, 'oneSidedColorCost': 0.25, 'twoSidedBWCost': 0.10, 'twoSidedColorCost': 0.20}
currJobNum = 1
csvHasHeader = True #variable indicating that CSV file has a header. Can be switched for csv files without headers.
totalPrintCost = 0.0 #variable to hold total cost for printing

def csv_reader (file):	# function to read file object and return raw data
    return csv.reader(file)

#function for calculation of printing costs.
def calculatePrintJobCost(printData, costDict):
    bwPages = int(printData[0]) - int(printData[1])
    colorPages = int(printData[1])
    isTwoSided = printData[2].replace(' ', '')

    if isTwoSided == 'true':
        printJobCost = (bwPages * costDict.get('twoSidedBWCost')) + (colorPages * costDict.get('twoSidedColorCost'))
        return "{:.2f}".format(printJobCost)
    elif isTwoSided == 'false':
        printJobCost = (bwPages * costDict.get('oneSidedBWCost')) + (colorPages * costDict.get('oneSidedColorCost'))
        return "{:.2f}".format(printJobCost)

#main
if __name__ == "__main__":

    csvFilePath = "sample.csv" #assumes that csv file is in same folder as python file

    with open(csvFilePath, "r") as fileObject:
        reader = csv_reader(fileObject)

        for row in reader:	#looping through csv data for individual records
            if csvHasHeader:	#skips first line for csv files with headers
                csvHasHeader = False
            else:
                currentJobCost = calculatePrintJobCost(row, a4PrintCosts)
                print('Cost for print job ' + str(currJobNum) + ' is: $' + currentJobCost)
                currJobNum += 1
                totalPrintCost += float(currentJobCost)	#Outputs individual printing costs to console

        print('Total Printing Cost is: $' + "{:.2f}".format(totalPrintCost))	#Outputs total printing cost to console