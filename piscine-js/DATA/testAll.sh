#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RESET='\033[0m'

# Function to center-align text with color
text() {
   local text="$1"
   local color="\033[38;5;208m"
   printf "$(echo -e "${color}${text}${RESET}")\n"
}

# Function to center-align text with color
redtext() {
   local text="$1"
   local color='\033[0;31m'
   printf "$(echo -e "${color}${text}${RESET}")\n"
}

wait_for_key() {

    while true; do
        read -rsn 3 key < /dev/tty
        if [[ "$key" == $'\e[C' ]] || [[ "$key" == $'\e[D' ]] || [[ "$key" == $'\e[B' ]]; then
            break
        fi
    done
}

test_DATA() {
    text 'change'
    node change_test.js
    wait_for_key

    text 'circular'
    node circular_test.js
    wait_for_key

    
}

alias jss=test_DATA