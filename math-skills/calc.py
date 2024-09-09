import math
import sys

def readData(fileName):
    try:
        with open(fileName, 'r') as file:
            data = [int(line.strip()) for line in file if line.strip().isdigit()]

        # Check if the data list is empty
        if not data:
            print("Error: didn't found any numbers in you file.")
            sys.exit(1)  # Exit the program with a non-zero status to indicate error

        return data

    except FileNotFoundError:
        print(f"Error: The file '{fileName}' was not found.")
        sys.exit(1) 

def average(data):
    return sum(data) / len(data)

def median(data):
    sortedData = sorted(data)
    mid = len(sortedData) // 2
    if len(sortedData) % 2 == 0:
        return (sortedData[mid - 1] + sortedData[mid]) / 2
    else:
        return sortedData[mid]

def variance(data, avg):
    return sum((x - avg) ** 2 for x in data) / len(data)

def standard_deviation(variance):
    return math.sqrt(variance)

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 calc.py <file name>")
        sys.exit(1)

    fileName= sys.argv[1]
    data = readData(fileName)
    
    avg = average(data)
    med = median(data)
    var = variance(data, avg)
    std_dev = standard_deviation(var)

    # Printing the rounded results
    print(f"Average: {round(avg)}")
    print(f"Median: {math.ceil(med)}")
    print(f"Variance: {round(var)}")
    print(f"Standard Deviation: {round(std_dev)}")

# Calling the main function.
main()
