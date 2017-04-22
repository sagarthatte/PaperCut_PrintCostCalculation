""" Code to read print orders from sample.csv file and output printing costs
    Input: CSV Data for Print jobs from .csv file
    Output: Printing cost for individual print jobs and total cost of printing
"""

#import module to open and read CSV files.
import csv

# function to read file object and return raw data
def csv_reader (file):
    return csv.reader(file)


if __name__ == "__main__":
    if __name__ == '__main__':
        csvFilePath = "sample.csv" #assumes that csv file is in same folder as python file

        with open(csvFilePath, "r") as fileObject:
            reader = csv_reader(fileObject)
#test for rows in csv file being returned as lists
        for row in reader:
            print(row)


