# ASCII Font Converter

## Introduction

This is a simple program designed to convert ASCII art fonts from the [patorjk/figlet.js](https://github.com/patorjk/figlet.js/blob/main/fonts) repository into text files. These text files can then be utiilzed in the "ascii art" project in golang.

## Requirements

Go programming language installed on your system
Access to the internet to download the font files from the repository

## Installation
Clone or download this repository to your local machine.  
Navigate to the directory containing the main.go file.  
Run the following command to compile the program:  

    ```go
    go build main.go
    ```

## Usage
After compiling, execute the program by running:

```md
  ./main <font_file_name> <output_file_name>
```

or without building:

```md
  go run . <font_file_name> <output_file_name>
```

Replace `<font_file_name>` with the name of the font file you wish to convert (e.g., `standard.flf`), and `<output_file_name>` with the desired name of the output text file (e.g., `output`).  
The program will read the font file, convert it into a text file, and save it with the specified name in the current directory.

### Example

Suppose you want to convert the font file `standard.flf` into a text file named `standard.txt`. You would execute the following command:

``` bash
./main standard.flf standard
```

## Credits

- This program is based on the Golang programming language.
- ASCII fonts are sourced from the [patorjk/figlet.js](https://github.com/patorjk/figlet.js/blob/main/fonts) repository.
- [fonts](http://www.jave.de/figlet/fonts/overview.html)
- [Asciiart EU](https://www.asciiart.eu/text-to-ascii-art)

Feel free to modify and expand upon this README file as needed!
