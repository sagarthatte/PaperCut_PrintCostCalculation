""" Code to read print orders from sample.csv file and output printing costs
    Input: CSV Data for Print jobs from .csv file
    Output: Printing cost for individual print jobs and total cost of printing
"""

#import module to open and read CSV files.
import csv

#cost variables for different types of prints
oneSidedBWCost = 0.15
oneSidedColorCost = 0.25
twoSidedBWCost = 0.10
twoSidedColorCost = 0.20

csvHasHeader = True #can switch this value if CSV File does not contain header
totalPrintCost = 0.0 #variable to hold total cost for all print job.
# function to read file object and return raw data
def csv_reader (file):
    return csv.reader(file)

if __name__ == "__main__":

    csvFilePath = "sample.csv" #assumes that csv file is in same folder as python file

    with open(csvFilePath, "r") as fileObject:
        reader = csv_reader(fileObject)

        for row in reader:
            if csvHasHeader:
                csvHasHeader = False
            else:
                bwPages = int(row[0]) - int(row[1])
                colorPages = int(row[1])
                isTwoSided = row[2].replace(' ', '')

                if isTwoSided == 'true':
                    printJobCost = (bwPages * twoSidedBWCost) + (colorPages * twoSidedColorCost)
                    print('Cost for this print job is: $' + "{:.2f}".format(printJobCost))
                    totalPrintCost += printJobCost
                elif isTwoSided == 'false':
                    printJobCost = (bwPages * oneSidedBWCost) + (colorPages * oneSidedColorCost)
                    print('Cost for this print job is: $' + "{:.2f}".format(printJobCost))
                    totalPrintCost += printJobCost

        print('Total Printing Cost is: $' + "{:.2f}".format(totalPrintCost))

        #"{:.9f}".format(numvar)
