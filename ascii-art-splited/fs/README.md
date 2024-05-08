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

- **standard**
- **shadow**
- **thinkertoy**
- **zigzag**
- **arob**

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
