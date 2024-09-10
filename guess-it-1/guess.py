import math
import sys

def calculate_mean(numbers):
    return sum(numbers) / len(numbers)

def calculate_standard_deviation(numbers, mean):
    return math.sqrt(sum((x - mean) ** 2 for x in numbers) / len(numbers))

def calculate_range(mean, stddev, factor=1.5):
    lower_bound = mean - (factor * stddev)
    upper_bound = mean + (factor * stddev)
    return int(lower_bound), int(upper_bound)

def main():
    numbers = []
    factor = 1.5  # Adjust for range size
    for line in sys.stdin:
        current_number = int(line.strip())
        if numbers:
            mean = calculate_mean(numbers)
            stddev = calculate_standard_deviation(numbers, mean)
            lower_bound, upper_bound = calculate_range(mean, stddev, factor)
            print(f"{lower_bound} {upper_bound}")
        print(current_number)
        numbers.append(current_number)

main()
