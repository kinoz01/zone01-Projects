## ASCII Art Program in Go
### Overview

This program can generates ASCII art representations of input strings. The program will take a string as input and output its graphical representation using ASCII characters. The ASCII representations will be based on predefined templates stored in banner files.

### Functionality

- Input: The program will accept strings containing numbers, letters, spaces, special characters, and newline characters `('\n')`.
- Output: It will generate ASCII art representations of the input strings according to predefined graphical templates stored in banner files. Each character will be represented as a series of ASCII characters with a height of 8 lines, separated by newline characters `('\n')`.

### Example

```bash
go run . "Hello World"
```

**Result:**
```md
 _    _          _   _                __          __                 _       _  
| |  | |        | | | |               \ \        / /                | |     | | 
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | 
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | 
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | 
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| 
                                                                                
                                                                                
```            

### Fonts Usage

You can change the displaying font by running: 

```bash
go run . "text" font
```

Available fonts:

PLease checkout the `banners` folder.

### Example Using fonts

```console
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

## Options
```console
Usage: go run . [OPTION] [STRING] [BANNER]
```

### Writing Output to File

You can write the output to a file using the `--output=<fileName.txt>` flag. For example:

```console
go run . --output=output.txt "Hello World" thinkertoy
```

### Align Output

This option allows you to adjust the alignment of Ascii Art output according to your preferences. The alignment is based on your actual terminal window size. Ensure the terminal window is large enough to display the text appropriately.

Supported alignment types:

- **center**: Aligns the text at the center.
- **left**: Aligns the text to the left.
- **right**: Aligns the text to the right.
- **justify**: Justifies the text.

**Usage Examples:**

Assume the bars in the display below are the terminal borders:

```console
|$ go run . --align=center "hello" standard                                                                                 |
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=left "Hello There" standard                                                                             |
| _    _           _    _                 _______   _                                                                       |
|| |  | |         | |  | |               |__   __| | |                                                                      |
|| |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___                                           |
||  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \                                          |
|| |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/                                          |
||_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|                                          |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=right "hello" shadow                                                                                    |
|                                                                                                                           |
|                                                                                          _|                _| _|          |
|                                                                                          _|_|_|     _|_|   _| _|   _|_|   |
|                                                                                          _|    _| _|_|_|_| _| _| _|    _| |
|                                                                                          _|    _| _|       _| _| _|    _| |
|                                                                                          _|    _|   _|_|_| _| _|   _|_|   |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=justify "how are you" shadow                                                                            |
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|$                                                                                                                          |
```
