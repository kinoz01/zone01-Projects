#!/bin/bash

# Define a function that runs multiple commands
run_go_commands() {
    echo "Test 1:"
    go run . hello | cat -n
    sleep 0.3

    echo "Test 2:"
    go run . hey
    sleep 0.3

    echo "Test 3:"
    go run . --color=red ll lllHello
    sleep 0.3

    echo "Test 4:"    
    go run . "\n\n\n" | cat -n
    sleep 0.3

    echo "Test 5:"
    go run . "--color=orange" "GuYs" "HeY GuYs?"
    sleep 0.3

    echo "Test 6:"
    go run . "--color=blue" "B" 'RGB()'
    sleep 0.3

    echo "Test 7:"
    go run . '--color=yellow' '(%&) ??'
    sleep 0.3

    echo "Test 8:"
    go run . '--color=green' '1 + 1 = 2'
    sleep 0.3

    echo "Test 9:"
    go run . --color=blue shadow shadow shadow
    sleep 0.3

    echo "Test 10:"
    go run . '--color=red' 'hello world'
    sleep 0.3

    echo "Test 11:"
    go run . '--color=blue' hey --output=h.txt "hey Hello"
    cat h.txt
    sleep 0.3

    echo "Test 12:"
    go run . '--color=lemon' "Hello\nWorld"
}


# Define an alias that calls the function
alias rgc='run_go_commands'
