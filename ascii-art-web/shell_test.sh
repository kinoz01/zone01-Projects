#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
BLUE='\036m'
RESET='\033[0m'

text() {
   local txt="$1"
   local color="\033[38;5;208m"
   printf "$(echo -e "${color}${txt}${RESET}")\n"
}

redtext() {
   local txt="$1"
   local color='\033[0;31m'
   printf "$(echo -e "${color}${txt}${RESET}")\n"
}

wait_for_key() {

    while true; do
        read -rsn 3 key < /dev/tty
        if [[ "$key" == $'\e[C' ]] || [[ "$key" == $'\e[D' ]] || [[ "$key" == $'\e[B' ]]; then
            break
        fi
    done
}

colorize_pattern() {

    while IFS= read -r line; do
        if [[ $line =~ HTTP/1\.1\ [0-9]+\ .+ ]]; then
            echo -e "\e[36m$line\e[0m"
        elif [[ $line =~ HTTP/2\ [0-9]+\ .+ ]]; then
            echo -e "\e[36m$line\e[0m"
        else
            echo "$line"
        fi
    done
}


test_web() {

    text 'Test 1 GET / (200 OK)'
    curl -X GET http://127.0.0.1:8088/ -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 2 POST / (Method Not Allowed)'
    curl -X POST http://127.0.0.1:8088/ -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 3 DELETE / (Method Not Allowed)'
    curl -X DELETE http://127.0.0.1:8088/ -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 4 GET /random (404)'
    curl -X GET http://127.0.0.1:8088/random -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 5 GET /ascii-art (Method not Allowed)"
    curl -X GET http://127.0.0.1:8088/ascii-art -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 6 POST /ascii-art (Empty text)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=&banner=standard" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 7 POST /ascii-art (Empty banner)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=HEY&banner=" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 8 POST /ascii-art (Both Empty)"
    curl http://127.0.0.1:8088/ascii-art -d "text=&banner=" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 9 POST /ascii-art (valid post)"
    curl http://127.0.0.1:8088/ascii-art -d "text=HELLO&banner=shadow" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 10 PUT /ascii-art (Method not Allowed)"
    curl -X PUT http://127.0.0.1:8088/ascii-art -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 11 POST /ascii-art (Large Payload)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=$(wget rentry.co/xxll/raw)&banner=standard" -D - -o /dev/null | colorize_pattern
    rm raw.1 && wait_for_key

    text "Test 11.2 POST /ascii-art (Large Payload from random file)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=$(head -c 10000 < /dev/urandom | base64)&banner=standard" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 12 POST /ascii-art (Invalid Characters)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=éâ&banner=standard" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 13 POST /ascii-art (Special Characters)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d 'text=@#$^&*()_+{}:<>?&banner=standard' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 14 POST /ascii-art (Timeout)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=HEY&banner=standard" --max-time 1 -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 15 POST /ascii-art (All characters Standard)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d 'text=text=!\"#$&()\n*+,-./012345\n6789:<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_\`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~&banner=standard' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 16 POST /ascii-art (All characters Shadow)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d 'text=!\"#$&()\n*+,-./012345\n6789:<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_\`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~&banner=shadow' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 17 POST /ascii-art (All characters Thinkertoy)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d 'text=!\"#$&()\n*+,-./012345\n6789:<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_\`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~&banner=thinkertoy' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text "Test 18 POST /ascii-art (XSS in Font)"
    curl -X POST http://127.0.0.1:8088/ascii-art -d 'text=HEY&banner=<script>alert("XSS")</script>' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 19 POST / (post in / with data)'
    curl -X POST http://127.0.0.1:8088/ -d 'text=HEY&banner=standard' -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 20 POST /ascii-art (special characters)'
    curl -X POST http://127.0.0.1:8088/ascii-art -d "text=;%&banner=standard" -D - -o /dev/null | colorize_pattern
    wait_for_key

    text 'Test 21 Benchmarking localhost vs raw ip'
    text localhost: 
    curl -o /dev/null -s -w 'Time: %{time_total}\n' http://localhost:8088
    text 127.0.0.1: 
    curl -o /dev/null -s -w 'Time: %{time_total}\n' http://127.0.0.1:8088

}

# Form Values are to be found in (inspect/network/request)
# or (inspect/network/payload)

# Get it from rentry and run it:
    # unset HISTFILE
    # wget rentry.co/aaw911a/raw

# Define an alias that calls the functions
alias web='test_web'

: <<'END_COMMENT'
function selectTrueRadioButtons() {
    // Select all radio buttons with value "true"
    const trueOptions = document.querySelectorAll('input[type="radio"][value="true"]');

    // Loop through the selected radio buttons and set their checked property to true
    trueOptions.forEach(function(option) {
        option.checked = true;
    });
}
selectTrueRadioButtons();
END_COMMENT

# sed -i -E 's/(curl.+-d.+null)/\1 -d "submit=show"/' raw
# sed -i -E 's/banner=/wwwww=/' raw
# sed -i -E 's/text=/wwwww=/' raw
# sed -i 's/\r//' raw
