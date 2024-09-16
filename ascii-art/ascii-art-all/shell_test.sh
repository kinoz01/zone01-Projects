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


test_ascii() {
    text 'Test 0 <HELLO>'
    go run . HELLO | cat -e
    wait_for_key

    text 'Test 1 <all chars>'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' | cat -e
    wait_for_key

    text 'Test 2 <HELLO\_n\_nHow Are YOU>'
    go run . 'HELLO\n\nHow Are YOU' | cat -e
    wait_for_key

    text 'Test 3 <\_nHEY THERE\_n>'
    go run . '\nHEY THERE\n' | cat -e
    wait_for_key

    text 'Test 4 <"">'
    go run . ""
    wait_for_key

    text 'Test 5 <"\_n\_n\_n">'
    go run . "\n\n\n" | cat -n
    wait_for_key

    text 'Test 6 <"hey hey">'
    go run . "hey hey" | cat -e
    wait_for_key

    text 'Test 7 <some Bash special variables>'
    go run . $$
    go run . $PPID
    go run . $SHELL
    go run . $BASH_VERSION
    go run . $RANDOM
    go run . $SECONDS
    wait_for_key

    text 'Test 8 <some emojis>'
    go run . '<0>..<0>' | cat -e && wait_for_key
    go run . '^j^' | cat -e && wait_for_key
    go run . ':-@' | cat -e && wait_for_key
    go run . 'D<'  | cat -e && wait_for_key
    go run . '@_@' | cat -e && wait_for_key
    go run . ':=8)'| cat -e 
    
}


test_fs() {
    text 'Test 0 <HELLO standard>'
    go run . HELLO standard
    wait_for_key

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
    go run . shadow standard | cat -e
    wait_for_key

    redtext 'Test 7 <hello hey wassup what what>'
    go run . hello hey wassup what what
    wait_for_key

    text 'Test 8 <all chars>'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' thinkertoy 
    wait_for_key

    text 'Test 9 <all chars>'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' shadow | cat -e
    wait_for_key

    text 'Test 10 <all chars>'
    go run . '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' standard | cat -e
    wait_for_key

    redtext 'Test 11 <Hello standard shadow>'
    go run . Hello standard shadow
    wait_for_key

    redtext 'Test 12 <éâ thinkertoy>'
    go run . éâ thinkertoy
    wait_for_key

    text 'Test 13 <\_n\_n\_n thinkertoy>'
    go run . "\n\n\n" thinkertoy | cat -n
    wait_for_key

    text 'Test 14 <"">'
    go run . "" thinkertoy | cat -n

}


test_output() {
    text 'Test 1 <-output=outl.txt Hello>'
    go run . --output=outl.txt Hello 
    cat -e outl.txt && rm outl.txt 
    wait_for_key

    text 'Test 2 <--output=outl.txt \_n\_n\_n>'
    go run . --output=outl.txt "\n\n\n"
    cat -e outl.txt && rm outl.txt
    wait_for_key

    text 'Test 3 <--output=outl.txt Hello\_nHey shadow>'
    go run . --output=outl.txt "Hello\nHey" shadow
    cat -e outl.txt && rm outl.txt
    wait_for_key

    text 'Test 4 <--output=outl.txt --output=hey.txt Hello\_nHey thinkertoy>'
    go run . --output=outl.txt --output=hey.txt "Hello\nHey" thinkertoy
    cat -e outl.txt 
    rm outl.txt && rm hey.txt
    wait_for_key

    text 'Test 5 <--output=outl.txt hello --output=hey.txt Hello thinkertoy>'
    go run . --output=outl.txt hello --output=hey.txt Hello thinkertoy
    cat -e outl.txt && rm outl.txt
    wait_for_key

    text 'Test 6 <--output=outl.txt "">'
    go run . --output=outl.txt ""
    cat -e outl.txt && rm outl.txt
    wait_for_key

    redtext 'Test 7 <--output=.txt hey thinkertoy>'
    go run . --output=.txt "hey" thinkertoy
    cat -e .txt && rm .txt
    wait_for_key

    redtext 'Test 8 <--output=outl hey shadow>'
    go run . --output=outl hey shadow
    cat -e outl && rm outl
    wait_for_key

    text 'Test 9 <--output=txt>'
    go run . --output=txt
    wait_for_key

    text 'Test 10 <--output=>'
    go run . --output=
    wait_for_key

    text 'Test 11 <--output>'
    go run . --output=
    wait_for_key

    text 'Test 12 <--o>'
    go run . --o
    wait_for_key

    text 'Test 13 <-->'
    go run . --
    wait_for_key

    redtext 'Test 14 <--output=outl.txt "éâ">'
    go run . --output=outl.txt "éâ"
    wait_for_key

    text 'Test 15 <all chars>'
    go run . --output=outl.txt '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' thinkertoy
    cat -e outl.txt && rm outl.txt
    wait_for_key

    text 'Test 16 <--output=outl.txt \_n>'
    go run . --output=outl.txt "\n"
    cat -e outl.txt && rm outl.txt
    wait_for_key

    redtext 'Test 17 <--output=outl.txt --output=outl.txt>'
    go run . --output=outl.txt --output=outl.txt
    cat -e outl.txt && rm outl.txt
    wait_for_key

    redtext 'Test 18 <--output=outl.txt --output>'
    go run . --output=outl.txt --output
    cat -e outl.txt && rm outl.txt
    wait_for_key

    redtext 'Test 19 <hello --output=outl.txt hey>'
    go run . hello --output=outl.txt hey
    cat -e outl.txt && rm outl.txt
    wait_for_key

    redtext 'Test 20 <hello --output=outl.go>'
    go run . hello --output=outl.go
    cat -e outl.go && rm outl.go

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

    text 'Test 5 <justify>'
    go run . --align=justify "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 6 <right>'
    go run . --align=right "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 7 <center>'
    go run . --align=center "Hello World\nHey There How\nare You"
    wait_for_key

    text 'Test 8 <--align=center hello | cat -n>'
    go run . --align=center hello | cat -e
    wait_for_key

    text 'Test 9 <--align=right hey>'
    go run . --align=right hey
    wait_for_key

    text 'Test 10 (three "new lines")'
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
    go run . --align=center '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~' thinkertoy | cat -e 
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

    text 'Test 34 <--output=hello.txt --align=center Hey>'
    go run . --output=hello.txt --align=center Hey
    cat hello.txt && rm hello.txt
    wait_for_key

    text 'Test 35 <--align=center --output=hello.txt Hey>'
    go run . --output=hello.txt --align=center Hey
    cat hello.txt && rm hello.txt
    wait_for_key

    redtext 'Test 36 <--align= center Hey>'
    go run . --align= center Hey

}


test_color() {
    text 'Test 1 <--color=blue hello>'
    go run . --color=blue hello 
    wait_for_key

    text 'Test 2 <--color=blue o hello>'
    go run . --color=blue o hello
    wait_for_key

    text 'Test 3 <--color=red ll lllHello>'
    go run . --color=red ll lllHello
    wait_for_key

    text 'Test 4 <--color=yellow "HEY  what" "yay HEy  whatssup">'
    go run . --color=yellow "HEY  what" "yay HEY  whatssup HEY"
    wait_for_key

    text 'Test 5 <--color=orange "GuYs" "HeY\_nGuYs?>'
    go run . "--color=orange" "GuYs" "HeY\nGuYs?"
    wait_for_key

    text 'Test 6 <--color=blue "B" RGB()>'
    go run . "--color=blue" "B" 'RGB()'
    wait_for_key

    text 'Test 7 <--color=yellow "  HEY  what" "y  HEY  whatsup HEY">'
    go run . --color=yellow "  HEY  what" "y  HEY  whatsup HEY"
    wait_for_key

    text 'Test 8 <--color=red HeyHey HeyHeyHeyYoYouHey>'
    go run . --color=red HeyHey HeyHeyHeyYoYouHey
    wait_for_key

    text 'Test 9 <--color=blue shadow shadow shadow>'
    go run . --color=blue shadow shadow shadow
    wait_for_key

    redtext 'Test 10 <invalid color>'
    go run . --color=string 'hello world'
    wait_for_key

    text 'Test 11 <--color=blue hey --output=h.txt "hey Hello" | cat h.txt>'
    go run . --color=blue hey --output=h.txt "hey Hello"
    cat h.txt
    rm h.txt
    wait_for_key

    text 'Test 12 <--color=yellow "Hello\_nWorld">'
    go run . --color=yellow "Hello\nWorld"
    wait_for_key

    text 'Test 13 <testing rgb color>'
    go run . '--color=rgb(100, 210, 40)' "Hello\nWorld"
    wait_for_key

    text 'Test 14 <multiple color flags>'
    go run . --color=green '!' --color=yellow '"' --color=blue '#' --color=magenta '$' --color=cyan '%' --color=white '&' --color=sky "'" --color=orange '(' --color=forest ')' --color=lavender '*' --color=rose '+' --color=lemon , --color=turquoise '-' --color=cherry '.' --color=emerald '/' --color=red 0 --color=green 1 --color=yellow 2 --color=blue 3 --color=magenta 4 --color=cyan 5 --color=white 6 --color=sky 7 --color=orange 8 --color=forest 9 --color=ocean ':' --color=lavender ';' --color=rose '<' --color=lemon = --color=turquoise '>' --color=cherry '?' --color=emerald '@' --color=red A --color=green B --color=yellow C --color=blue D --color=magenta E --color=cyan F --color=white G --color=sky H --color=orange I --color=forest J --color=ocean K --color=lavender L --color=rose M --color=lemon N --color=turquoise O --color=cherry P --color=emerald Q --color=red R --color=green S --color=yellow T --color=blue U --color=magenta V --color=cyan W --color=white X --color=sky Y --color=orange Z --color=forest '[' --color=ocean '\' --color=lavender ']' --color=rose '^' --color=lemon _ --color=turquoise '`' --color=cherry a --color=emerald b --color=red c --color=green d --color=yellow e --color=blue f --color=magenta g --color=cyan h --color=white i --color=sky j --color=orange k --color=forest l --color=ocean m --color=lavender n --color=rose o --color=lemon p --color=turquoise q --color=cherry r --color=emerald s --color=red t --color=green u --color=yellow v --color=blue w --color=magenta x --color=cyan y --color=white z --color=sky '{' --color=orange '|' --color=forest '}' --color=ocean '~' '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
    wait_for_key

    text 'Test 15: <--color=red he --color=green ey hey>'
    go run . --color=red he --color=green ey hey
    wait_for_key

    redtext 'Test 16 <"--color=rgb(2503, 210, 40)" "Hello World">'
    go run . '--color=rgb(2503, 210, 40)' "Hello World"
    wait_for_key

    redtext 'Test 17 <--color=hsl(-1, 50, 40) "Hello World">'
    go run . '--color=hsl(-1, 50%, 40%)' "Hello World"
    wait_for_key

    redtext 'Test 18: <hey "--color=red" green "Hello World">'
    go run . hey '--color=red' green "Hello World"
    wait_for_key

    redtext 'Test 19: <--color=red green green "Hello World">'
    go run . '--color=red' green green "Hello World"
    wait_for_key

    redtext 'Test 20: <--color= green "Hello World">'
    go run . '--color=red' green green "Hello World"
    wait_for_key

    text 'Test 21: <--color=red "o\nWo" "Hello\_nWorld">'
    go run . --color=red "o\nWo" "Hello\nWorld"
    wait_for_key

    text 'Test 22 <--color=red --align=center Hey>'
    go run . --color=red --align=center Hey
    wait_for_key

    text 'Test 23 <--color=red H --align=center Hey>'
    go run . --color=red H --align=center Hey

}


test_reverse() {
    text 'Test 1 <--output=testR.txt "       hey     hello  ">'
    go run . --output=testR.txt "       hey     hello  "
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 2 <all chars with new lines>' 
    go run . --output=testR.txt '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 3 <--output=testR.txt " hey \_n   What? ">'
    go run . --output=testR.txt " hey \n   What? " 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 4 <--output=testR.txt "hello how are you!">'
    go run . --output=testR.txt "hello how are you!" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 5 <only spaces>'
    go run . --output=testR.txt "  " 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 6 <only one \_n>'
    go run . --output=testR.txt "\n" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 7 <Three \_n>'
    go run . --output=testR.txt "\n\n\n" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 8 <nothing>'
    go run . --output=testR.txt "" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 9 <hey\_n\_n\_n\_nlol>'
    go run . --output=testR.txt "hey\n\n\n\nlol" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 10 <\_n\_n\_nhey\_n\_n\_n\_n>'
    go run . --output=testR.txt "\n\n\nhey\n\n\n\n" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 11 <\_n\_nhey\_n\_n\_n\_nlol\_n\_n>'
    go run . --output=testR.txt "\n\nhey\n\n\n\nlol\n\n" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    redtext 'Test 12 <--reverse=testR.txt shadow>'
    go run . --output=testR.txt "\n\nhey\n\n\n\nlol\n\n" 
    go run . --reverse=testR.txt shadow | cat -e
    wait_for_key

    text 'Test 13 <\_n\_nhey    \_n\_n\_n\_n  the   what  \_n\_n>'
    go run . --output=testR.txt "\n\nhey   \n\n\n\n  the   what  \n\n" 
    go run . --reverse=testR.txt | cat -e
    wait_for_key

    text 'Test 14 <  \_n\_n  hey    \_n\_n\_n\_n  the   what  \_n\_n>'
    go run . --output=testR.txt "  \n\n  hey   \n\n\n\n  the   what  \n\n" 
    go run . --reverse=testR.txt | cat -e
    rm testR.txt
}

# Get it from dpaste and run it:
    # unset HISTFILE
    # wget rentry.co/aat911a/raw


alias ascii='test_ascii'
alias fst='test_fs'
alias out='test_output'
alias align='test_align'
alias color='test_color'
alias reverse='test_reverse'

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

# sed -i 's/\r//' raw
