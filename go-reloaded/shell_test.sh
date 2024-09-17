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

check_result() {
    local result_text=$1
    shift
    local expected_outputs=("$@")

    local pass=false
    for expected in "${expected_outputs[@]}"; do
        if [ "$result_text" = "$expected" ]; then
            pass=true
            break
        fi
    done

    if [ "$pass" = true ]; then
        echo -e '\033[0;32mPASS: Test successful!\033[0m'
    else
        redtext "FAIL: Result does not match expected output(s)!"
        echo "Expected:"
        for expected in "${expected_outputs[@]}"; do
            echo "- $expected"
        done
        echo "Actual:"
        echo "- $result_text"
    fi
}

run_tests() {
    # Function to perform a single test
    run_single_test() {
        local test_number="$1"
        local sample_text="$2"
        shift 2
        local expected_outputs=("$@")

        text "Test $test_number ----->"
        echo "$sample_text"

        # Temporary files for testing
        local sample_file="sample_test.txt"
        local result_file="result_test.txt"

        # Write sample text to the temporary file
        echo "$sample_text" > "$sample_file"

        # Run the Go program
        go run . "$sample_file" "$result_file"

        # Read the result from the output file
        local result_text
        result_text=$(cat "$result_file")

        # Call the check_result function
        check_result "$result_text" "${expected_outputs[@]}"
    }

    # Test cases
    run_single_test 1 "1E (hex) files were added" "30 files were added"
    wait_for_key
    run_single_test 2 "It has been 10 (bin) years" "It has been 2 years"
    wait_for_key
    run_single_test 3 "Ready, set, go (up) !" "Ready, set, GO!"
    wait_for_key
    run_single_test 4 "I should stop SHOUTING (low)" "I should stop shouting"
    wait_for_key
    run_single_test 5 "This is so exciting (up, 2)" "This is SO EXCITING"
    wait_for_key
    run_single_test 6 "Welcome to the Brooklyn bridge (cap)" "Welcome to the Brooklyn Bridge"
    wait_for_key
    run_single_test 7 "I was sitting over there ,and then BAMM !!" "I was sitting over there, and then BAMM!!"
    wait_for_key
    run_single_test 8 "I am exactly how they describe me: ' awesome '" "I am exactly how they describe me: 'awesome'"
    wait_for_key
    run_single_test 9 "' I am the most well-known homosexual in the world '" "'I am the most well-known homosexual in the world'"
    wait_for_key
    run_single_test 10 "There it was. A amazing rock!" "There it was. An amazing rock!"
    wait_for_key
    run_single_test 11 "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair." \
        "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
    wait_for_key
    run_single_test 12 "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." "Simply add 66 and 2 and you will see the result is 68."
    wait_for_key
    run_single_test 13 "There is no greater agony than bearing a untold story inside you." "There is no greater agony than bearing an untold story inside you."
    wait_for_key
    run_single_test 14 "Punctuation tests are ... kinda boring ,don't you think !?" "Punctuation tests are... kinda boring, don't you think!?"
    wait_for_key
    run_single_test 15 "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" \
        "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"
    run_single_test 16 "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure" "I have to pack 5 outfits. Packed 26 just to be sure"
    wait_for_key
    run_single_test 17 "Don not be sad ,because sad backwards is das . And das not good" "Don not be sad, because sad backwards is das. And das not good"
    wait_for_key
    run_single_test 18 "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" \
        "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"
    wait_for_key
    run_single_test 19 "(A     universe is loading" "(An     universe is loading" "(An universe is loading"
    wait_for_key
    run_single_test 20 " 'eee' 'e e e' ' e e e ' 'e e e ' ' e e e'  (up,2)" " 'eee' 'e e e' 'e e e' 'e e e' 'e E E'" "'eee' 'e e e' 'e e e' 'e e e' 'e E E'" "'eee' 'e e e' 'e e e' 'e e e' 'e e e' (up,2)"
    wait_for_key
    run_single_test 21 "a b (cap, 2)" "A B"
    wait_for_key
    run_single_test 22 "' . . . ' ' . . . ' ' . . . ' ' . . . ' '. . . '" "'...' '...' '...' '...' '...'"
    wait_for_key
    run_single_test 23 "' . . . ' ' . . . '"$'\n'"' . . . ' ' . . . ' '. . . '" "'...' '...'"$'\n'"'...' '...' '...'"
    wait_for_key
    run_single_test 24 "Hey Hey Hello (up, 3" "Hey Hey Hello (up, 3" "HEY HEY HELLO"
    wait_for_key
    run_single_test 25 "I have a there (up, a(hex)(bin))" "I have A THERE"
    wait_for_key

    run_single_test 26 "I don't ,mind ' being 'here ." "I don't, mind 'being' here."
    wait_for_key

    run_single_test 27 "(bin)" "" "(bin)"
    wait_for_key

	 run_single_test 28 "(ok)" "(ok)"
    wait_for_key

	run_single_test 29 "" ""
    wait_for_key

	run_single_test 30 "(" "("
    wait_for_key

	run_single_test 31 ")" ")"
    wait_for_key

	run_single_test 32 "a a (hex)" "a 10"
    wait_for_key

	run_single_test 33 "a a a a a a apple" "an an an an an an apple"
    wait_for_key

	run_single_test 34 "hello there ... (cap, 2)" "Hello There..."
    wait_for_key

	run_single_test 35 "hello there... (cap, 2)" "Hello There..."
    wait_for_key

	run_single_test 36 "hello there ... (up, 3)" "HELLO THERE..."
    wait_for_key

    run_single_test 37 "hello there ... (up, -3)" "HELLO THERE..."
    wait_for_key

    run_single_test 38 "hello there 10 ... (((bin)))" 
    wait_for_key

    run_single_test 39 "hello there 10 ... (bin)" 
    wait_for_key

    # Clean up temporary files
    rm sample_test.txt result_test.txt
}

alias goo='run_tests'

: <<'END_COMMENT'
function selectYesButtons() {
    const yesOptions = document.querySelectorAll('input[type="radio"][value="true"]');

    yesOptions.forEach(function(option) {
        option.checked = true;
    });
}
selectYesButtons();
END_COMMENT

# Get it from rentry and run it:
    # unset HISTFILE
    # wget rentry.co/grl911a/raw

# sed -i 's/\r//' raw