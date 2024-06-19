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
   #local width=$(( ($(tput cols) + ${#text}) / 2))
   #printf "%*s\n" $width "$(echo -e "${color}${text}${RESET}")"
   printf "$(echo -e "${color}${text}${RESET}")\n"
}

# Function to center-align text with color
redtext() {
   local text="$1"
   local color='\033[0;31m'
   #local width=$(( ($(tput cols) + ${#text}) / 2))
   #printf "%*s\n" $width "$(echo -e "${color}${text}${RESET}")"
   printf "$(echo -e "${color}${text}${RESET}")\n"
}

wait_for_key() {
    while true; do
        read -rsn 3 key </dev/tty
        if [[ "$key" == $'\e[C' ]] || [[ "$key" == $'\e[D' ]] || [[ "$key" == $'\e[B' ]]; then
            break
        fi
    done
}

# Testing Color and output.
run_go_commands() {
    text "Test 1"
    go run . hello | cat -n
    wait_for_key

    text "Test 2"
    go run . hey
    wait_for_key

    text "Test 3"
    go run . --color=red ll lllHello
    wait_for_key

    text "Test 4"
    go run . "\n\n\n" | cat -n
    wait_for_key

    text "Test 5"
    go run . "--color=orange" "GuYs" "HeY GuYs?"
    wait_for_key

    text "Test 6"
    go run . "--color=blue" "B" 'RGB()'
    wait_for_key

    text "Test 7"
    go run . '--color=yellow' '(%&) ??'
    wait_for_key

    text "Test 8"
    go run . '--color=green' '1 + 1 = 2'
    wait_for_key

    text "Test 9"
    go run . --color=blue shadow shadow shadow
    wait_for_key

    text "Test 10"
    go run . '--color=red' 'hello world'
    wait_for_key

    text "Test 11"
    go run . '--color=blue' hey --output=h.txt "hey Hello"
    cat h.txt
    wait_for_key

    text "Test 12"
    go run . '--color=lemon' "Hello\nWorld"
    wait_for_key

    text "Test 13"
    go run . --align=center '--color=rgb(100, 210, 40)' "Hello\nWorld"
    wait_for_key

    text "Test 14"
    go run . --color=green '!' --color=yellow '"' --color=blue '#' --color=magenta '$' --color=cyan '%' --color=white '&' --color=sky "'" --color=orange '(' --color=forest ')' --color=lavender '*' --color=rose '+' --color=lemon , --color=turquoise '-' --color=cherry '.' --color=emerald '/' --color=red 0 --color=green 1 --color=yellow 2 --color=blue 3 --color=magenta 4 --color=cyan 5 --color=white 6 --color=sky 7 --color=orange 8 --color=forest 9 --color=ocean ':' --color=lavender ';' --color=rose '<' --color=lemon = --color=turquoise '>' --color=cherry '?' --color=emerald '@' --color=red A --color=green B --color=yellow C --color=blue D --color=magenta E --color=cyan F --color=white G --color=sky H --color=orange I --color=forest J --color=ocean K --color=lavender L --color=rose M --color=lemon N --color=turquoise O --color=cherry P --color=emerald Q --color=red R --color=green S --color=yellow T --color=blue U --color=magenta V --color=cyan W --color=white X --color=sky Y --color=orange Z --color=forest '[' --color=ocean '\' --color=lavender ']' --color=rose '^' --color=lemon _ --color=turquoise '`' --color=cherry a --color=emerald b --color=red c --color=green d --color=yellow e --color=blue f --color=magenta g --color=cyan h --color=white i --color=sky j --color=orange k --color=forest l --color=ocean m --color=lavender n --color=rose o --color=lemon p --color=turquoise q --color=cherry r --color=emerald s --color=red t --color=green u --color=yellow v --color=blue w --color=magenta x --color=cyan y --color=white z --color=sky '{' --color=orange '|' --color=forest '}' --color=ocean '~' '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
    wait_for_key

    text 'Test 15: <--color=red he --color=green ey hey>'
    go run . --color=red he --color=green ey hey
}

test_errors() {
    text "Test 1"
    go run . '--color=rgb(2503, 210, 40)' "Hello World"
    wait_for_key

    text "Test 2"
    go run . '--color=hsl(-1, 50%, 40%)' "Hello World"
    wait_for_key

    text 'Test 3: <hey '--color=red' green "Hello World">'
    go run . hey '--color=red' green "Hello World"
    wait_for_key

    text 'Test 4: <--color=red green green "Hello World">'
    go run . '--color=red' green green "Hello World"
    wait_for_key
}

test_reverse() {
    text "Test 1"
    go run . --output=test1.txt "       hey     hello  "
    go run . --reverse=test1.txt | cat -e
    wait_for_key

    text "Test 2"
    go run . --output=test2.txt '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
    go run . --reverse=test2.txt
    wait_for_key

    text "Test 3"
    go run . --output=test2.txt " hey \n   What? " 
    go run . --reverse=test2.txt | cat -e
    wait_for_key
}

test_align() {
    text 'Test 1 <--align=justify "Hello World">'
    go run . --align=justify "Hello World"
    wait_for_key

    text 'Test 2 <--align=right "Hello World">'
    go run . --align=right "Hello World"
    wait_for_key

    text 'Test 3 <--align=center "Hello World">'
    go run . --align=center "Hello World"
    wait_for_key

    text 'Test 4 <--align=justify "Hello World">'
    go run . --align=justify "Hello World"
    wait_for_key

    text 'Test 5'
    go run . --align=justify "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 6'
    go run . --align=right "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 7'
    go run . --align=center "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 8 <--align=center hello | cat -n>'
    go run . --align=center hello | cat -e
    wait_for_key

    text 'Test 9 <--align=right hey>'
    go run . --align=right hey
    wait_for_key

    text 'Test 10'
    go run . "\n\n\n" | cat -n
    wait_for_key

    text 'Test 11 "" | cat -n'
    go run . "" | cat -n
    wait_for_key

    text 'Test 12'
    go run . --align=left "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 13 <--align=justify "     Hello              Hey     There    ">'
    go run . --align=justify "     Hello              Hey     There    "
    wait_for_key

    text 'Test 14 <--align=justify " t            g        g    ">'
    go run . --align=justify " t            g        g    "
    wait_for_key

    text 'Test 15'
    go run . --align=justify " t            g        g   \n  hey    g g "
    wait_for_key

    text 'Test 16 <Testing two align flags>'
    go run . --align=justify --align=left "Hey There"
    wait_for_key

    text 'Test 17---> 20 <Testing long Input>'
    go run . --align=justify "This is a very very very long text\nHello World"
    wait_for_key

    text 'Test 18'
    go run . --align=center "This is a very very very long text\nHello World"
    wait_for_key

    text 'Test 19'
    go run . --align=right "This is a very very very long text\nHello World"
    wait_for_key

    text 'Test 20'
    go run . --align=left "This is a very very very long text\nHello World"
    wait_for_key

    text 'Test 21 <--align=center "Hey there" shadow>'
    go run . --align=center "Hey there" shadow
    wait_for_key

    text 'Test 22 <--align=right "Hey there" thinkertoy>'
    go run . --align=right "Hey there" thinkertoy
    wait_for_key

    text 'Test 23 <Testing All Asciis>'
    go run . --align=center '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' thinkertoy
    wait_for_key

    redtext 'Test 24 <hey --align=center Hello>'
    go run . hey --align=center Hello
    wait_for_key

    text 'Test 25 <--align=center shadow shadow>'
    go run . --align=center shadow shadow
    wait_for_key

    redtext 'Test 25 <--align=center shadow hey>'
    go run . --align=center shadow hey
    wait_for_key

    redtext 'Test 26 <--align=center hello hey whatssup>'
    go run . --align=center hello hey whatssup
    wait_for_key

    text 'Test 27 <-->'
    go run . --
    wait_for_key

    text 'Test 28 <--align>'
    go run . --align
    wait_for_key

    text 'Test 29 <--align=>'
    go run . --align=
    wait_for_key

    text 'Test 30 <--align=what>'
    go run . --align=what
    wait_for_key

    text 'Test 31 <--align=center>'
    go run . --align=center 
    wait_for_key

    redtext 'Test 32 <--align=what hey>'
    go run . --align=what hey
    wait_for_key

    text 'Test 33 <--alig>'
    go run . --alig
    wait_for_key
}

test_fs() {
    text 'Test 1 <shadow shadow>'
    go run . shadow shadow
    wait_for_key

    text 'Test 2 <hey standard.txt>'
    go run . hey standard.txt
    wait_for_key

    text 'Test 3 <hey shadow.txt>'
    go run . hey shadow.txt
    wait_for_key

    text 'Test 4 <standard.txt>'
    go run . standard.txt
    wait_for_key

    text 'Test 5 <standard shadow>'
    go run . standard shadow
    wait_for_key

    text 'Test 6 <shadow standard>'
    go run . shadow standard
    wait_for_key

    redtext 'Test 7 <hello hey wassup what what>'
    go run . hello hey wassup what what
    sleep 0.5

    text 'Test 8'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' thinkertoy
    wait_for_key

    text 'Test 9'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' shadow
    wait_for_key

    text 'Test 10'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' standard
    wait_for_key
}


# Define an alias that calls the functions
alias rgc='run_go_commands'
alias reverse='test_reverse'
alias align='test_align'
alias error='test_errors'
alias fonts='test_fs'


: <<'END_COMMENT'
function clickYesButtons() {
    const YesButtons = document.querySelectorAll('.exerciseButton');

    // Click each button
    YesButtons.forEach(button => {
        button.click();
    });
}

// Execute the function
clickYesButtons();
END_COMMENT
