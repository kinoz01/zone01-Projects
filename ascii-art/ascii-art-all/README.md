## ASCII Art Program in Go
### Overview

<div align="center">
  <img src="doc/example.gif" alt="Alt Text" width="1000">
</div>


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

```bash
"small", "phoenix", "o2", "starwar", "stop", "varsity", "standard", "shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles", "jacky", "catwalk", "coins", "fire", "jazmine", "matrix", "blocks", "univers", "impossible", "georgi"
```

You can also use your own fonts. Just put them along side the executable file in a folder and name it "banners". Make sure they match the standard templates (only 8-lines font)!

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

### Color Output

This option allows you to manipulate text colors using the command line interface. You can specify a color and choose which letters to colorize within a given string. It offers flexibility in choosing between coloring a single letter or a set of letters. Note that while color substring support space characters, unfortunately it doesn't support newlines syntax implemented in the input (eg, `\n`) since they are considered as normal characters.

**Usage**

```console
go run . --color=<color> <substring to be colored> string [banner]
```

- You can directly color text without using substring: 

```console
go run . --color=<color> <text to be colored> [banner]
```

- You can also use as much color flags as you like:

```console
go run . --color=<color1> <substring1> --color=<color2> <substring2> ... --color=<colorN> <substringN> string [banner]
```

**Supported colors:**

- string colors
- RGB (ex: `rgb(255, 0, 0)`)
- `HSL` (ex: `hsl(0, 100%, 50%)`)
- Hexadecimal colors (ex: `#00ff00`)

**Note 2:**
Be careful when using characters like single quotes and others, you will need to escape them so the program can work correctly!

**Note 3:**
You can't use a substring that has a flag pattern as a color argument.


**Extreme Example**

```md
go run . --color=green '!' --color=yellow '"' --color=blue '#' --color=magenta '$' --color=cyan '%' --color=white '&' --color=sky "'" --color=orange '(' --color=forest ')' --color=lavender '*' --color=rose '+' --color=lemon , --color=turquoise '-' --color=cherry '.' --color=emerald '/' --color=red 0 --color=green 1 --color=yellow 2 --color=blue 3 --color=magenta 4 --color=cyan 5 --color=white 6 --color=sky 7 --color=orange 8 --color=forest 9 --color=ocean ':' --color=lavender ';' --color=rose '<' --color=lemon = --color=turquoise '>' --color=cherry '?' --color=emerald '@' --color=red A --color=green B --color=yellow C --color=blue D --color=magenta E --color=cyan F --color=white G --color=sky H --color=orange I --color=forest J --color=ocean K --color=lavender L --color=rose M --color=lemon N --color=turquoise O --color=cherry P --color=emerald Q --color=red R --color=green S --color=yellow T --color=blue U --color=magenta V --color=cyan W --color=white X --color=sky Y --color=orange Z --color=forest '[' --color=ocean '\' --color=lavender ']' --color=rose '^' --color=lemon _ --color=turquoise '`' --color=cherry a --color=emerald b --color=red c --color=green d --color=yellow e --color=blue f --color=magenta g --color=cyan h --color=white i --color=sky j --color=orange k --color=forest l --color=ocean m --color=lavender n --color=rose o --color=lemon p --color=turquoise q --color=cherry r --color=emerald s --color=red t --color=green u --color=yellow v --color=blue w --color=magenta x --color=cyan y --color=white z --color=sky '{' --color=orange '|' --color=forest '}' --color=ocean '~' '!"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
```
