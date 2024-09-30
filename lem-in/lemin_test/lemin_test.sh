#!/bin/bash

# Check if a file argument is provided
if [ $# -ne 1 ]; then
  echo "Usage: $0 <input_file>"
  exit 1
fi

# Store the input file argument
input_file=$1

# Run leminTest and lemin, saving their outputs to respective files
./leminTest "$input_file" > leminTest.txt
./lemin "$input_file" > lemin.txt

# Extract content after two consecutive newlines for lemin.txt
awk 'BEGIN {found=0} 
     { 
       if ($0 == "") {found=1}  # Detect a single blank line
       if (found && $0 != "") print  # Print lines after the blank line, but skip any additional blank lines
     }' lemin.txt > lemin_i.txt

# Count number of lines and words for both files
leminTest_lines=$(wc -l < leminTest.txt)
leminTest_words=$(wc -w < leminTest.txt)
lemin_lines=$(wc -l < lemin_i.txt)
lemin_words=$(wc -w < lemin_i.txt)

# Print the results
echo "leminTest: $leminTest_lines lines, $leminTest_words words"
echo "lemin    : $lemin_lines lines, $lemin_words words"


if [ "$leminTest_lines" -eq "$lemin_lines" ] && [ "$leminTest_words" -eq "$lemin_words" ]; then
  echo -e "\e[32mOK\e[0m"  
else
  echo -e "\e[31mKO\e[0m"  
fi
