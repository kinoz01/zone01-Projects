# go-reloaded

This program will accept two command-line arguments: the name of the input file containing text that needs modifications, and the output file where the modified text will be stored. Here are the specific text modifications the program should perform:

- **Number Using Flags**:
  - `(hex)`: Converts the preceding hexadecimal number to its decimal form.
  - `(bin)`: Converts the preceding binary number to its decimal form.

- **Text Case Modifications Using Flags**:
  - `(up)`: Converts the preceding word to uppercase.
  - `(low)`: Converts the preceding word to lowercase.
  - `(cap)`: Capitalizes the preceding word.
  - These modifiers can also specify a number of preceding words to modify, e.g., `(up, 2)`.

- **Punctuation Corrections**:
  - Standardize space around `.`, `,`, `!`, `?`, `:` and `;`.
  - Correct placement for `'`, including support for phrases within quotation marks.

- **Grammar Corrections**:
  - Automatic correction from `a` to `an` before words starting with a vowel or an `h`.

## Usage

Here is a command example to show how the program should work:

```console
$ go run . sample.txt result.txt
```
This command reads the input from sample.txt, applies the required text transformations, and writes the modified text to result.txt.

## Flag Parsing

The program is designed to recognize and process flags that indicate text modifications. Flags should be placed immediately following a word and can be formatted with or without spaces before or after the parentheses. Here are the valid formats for flag placement:

```md
(word)(flag)(word)
(word) (flag)(word)
(word)(flag) (word)
(word) (flag) (word)
```
All flags are parsed and removed after processing, except some cases when you put a flag inside a flag.  
Note: the program can't handle text with accent for now, avoid using characters like ê, â or é, thank you for your understanding. 
