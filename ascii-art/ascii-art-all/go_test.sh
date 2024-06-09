#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RESET='\033[0m'

# Function to center-align text with color
text() {
   local text="$1"
   local color='\033[0;31m'
   #local width=$(( ($(tput cols) + ${#text}) / 2))
   #printf "%*s\n" $width "$(echo -e "${color}${text}${RESET}")"
   printf "$(echo -e "${color}${text}${RESET}")\n"
}

# Testing Color and output.
run_go_commands() {
    text "Test 1"
    go run . hello | cat -n
    sleep 0.3

    text "Test 2"
    go run . hey
    sleep 0.3

    text "Test 3"
    go run . --color=red ll lllHello
    sleep 0.3

    text "Test 4"
    go run . "\n\n\n" | cat -n
    sleep 0.3

    text "Test 5"
    go run . "--color=orange" "GuYs" "HeY GuYs?"
    sleep 0.3

    text "Test 6"
    go run . "--color=blue" "B" 'RGB()'
    sleep 0.3

    text "Test 7"
    go run . '--color=yellow' '(%&) ??'
    sleep 0.3

    text "Test 8"
    go run . '--color=green' '1 + 1 = 2'
    sleep 0.3

    text "Test 9"
    go run . --color=blue shadow shadow shadow
    sleep 0.3

    text "Test 10"
    go run . '--color=red' 'hello world'
    sleep 0.3

    text "Test 11"
    go run . '--color=blue' hey --output=h.txt "hey Hello"
    cat h.txt
    sleep 0.3

    text "Test 12"
    go run . '--color=lemon' "Hello\nWorld"
    sleep 0.3

    text "Test 13"
    go run . --align=center '--color=rgb(100, 210, 40)' "Hello\nWorld"
    sleep 0.3
}

test_errors() {
    text "Test 1"
    go run . '--color=rgb(2503, 210, 40)' "Hello World"
    sleep 0.3

    text "Test 2"
    go run . '--color=hsl(-1, 50%, 40%)' "Hello World"
    sleep 0.3
}

test_reverse() {
    text "Test 1"
    go run . --output=test1.txt "       hey     hello  "
    go run . --reverse=test1.txt | cat -e
    sleep 0.3

    text "Test 2"
    go run . --output=test2.txt '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
    go run . --reverse=test2.txt
    sleep 0.3

    text "Test 3"
    go run . --output=test2.txt " hey \n   What? " 
    go run . --reverse=test2.txt | cat -e
    sleep 0.3
}

test_align() {
    text "Test 1"
    go run . --align=justify "Hello World"
    sleep 0.3

    text "Test 2"
    go run . --align=right "Hello World"
    sleep 0.3

    text "Test 3"
    go run . --align=center "Hello World"
    sleep 0.3
}

# Define an alias that calls the functions
alias rgc='run_go_commands'
alias reverse='test_reverse'
alias align='test_align'
alias error='test_errors'
