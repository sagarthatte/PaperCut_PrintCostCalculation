""" Code to read print orders from sample.csv file and output printing costs
    Input: CSV Data for Print jobs from .csv file
    Output: Printing cost for individual print jobs and total cost of printing
"""

#import module to open and read CSV files.
import csv

#create dictionary for print costs. Can be expanded to include other cost types.
a4PrintCosts = {'oneSidedBWCost': 0.15, 'oneSidedColorCost': 0.25, 'twoSidedBWCost': 0.10, 'twoSidedColorCost': 0.20}

csvHasHeader = True #can switch this value if CSV File does not contain header
totalPrintCost = 0.0 #variable to hold total cost for all print job.
# function to read file object and return raw data
def csv_reader (file):
    return csv.reader(file)

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


if __name__ == "__main__":

    csvFilePath = "sample.csv" #assumes that csv file is in same folder as python file

    with open(csvFilePath, "r") as fileObject:
        reader = csv_reader(fileObject)

        for row in reader:
            if csvHasHeader:
                csvHasHeader = False
            else:
                currentJobCost = calculatePrintJobCost(row, a4PrintCosts)
                print('Print Cost for current job is: $' + currentJobCost)
                totalPrintCost += float(currentJobCost)

        print('Total Printing Cost is: $' + "{:.2f}".format(totalPrintCost))