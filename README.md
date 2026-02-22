# Nova ✨

A lightweight Go terminal animation that render ANSI-powered starfield 

![demo](demo.gif)

## Features

- ANSI true-color (24 bit) star rending
- Alternative buffer
- Custom star symbol, duration and color
- Python based configaration

## Requirement

- If you want true-color you need to have a terminal emulator that support true-color
- `stty` must be in your machine, for windows wait for next release.
- `python` command won't say `command not found`, if you have `python3` alias it with `python`

## Installation

- If go already install `go install github.com/shv-ng/nova`

OR

- Download binary from [releases](https://github.com/shv-ng/nova/releases) 

## Usage

### help

```bash
$ nova -h
Usage of nova:
  -color string
    	color of all stars (default "200,0,100")
  -config string
    	load python based config file
  -count int
    	number of stars to display (default 100)
  -duration int
    	time (in s) to fade away a star (default 4)
  -symbol string
    	what to display - the star (default ".")

# inline style
$ nova -count=20 -duration=3 -symbol='*' -color=0,255,255

# use config file
$ nova -config=path/to/config.py
```

to close, hit universal killer `Ctrl-c`

> [!Note]
> the `-config` will override other command

> [!Warning]
> there is no great error handling/edge cases done yet, use with caution

### Config file

Use this example as template, and edit as per your need. It'll evaluated at runtime, so you are allow to do crazy stuffs

`config.py`

```python
from dataclasses import dataclass


@dataclass
class Star:
    symbol: str = ""
    duration: int = 0
    fg_color: str = "255,255,255"


STARS: list[Star] = [
    Star(symbol="·", duration=10, fg_color="100,100,100"),  
    Star(symbol="*", duration=15, fg_color="200,200,200"),  
    
    Star(symbol="✦", duration=25, fg_color="135,206,235"),  
    Star(symbol="✧", duration=20, fg_color="255,255,190"),  
    
    Star(symbol="❂", duration=40, fg_color="255,100,100"),  
    Star(symbol="✵", duration=35, fg_color="255,165,0"),    

    Star(symbol="⚡", duration=50, fg_color="191,64,191"),   
    Star(symbol="❈", duration=60, fg_color="0,255,255"),    
] 

STAR_COUNT: int = 100
SYMBOL: str = "."

DURATION: int = 5

# must be in format of rr,bb,gg i.e. 0,0,0 or 255,250,24, will add hex color support later
FG_COLOR: str = "200,0,100"
BG_COLOR: str = ""  # it wont work
```

### Color format

As of now, I didn't wrote parser and great edge case handling for colors, it'll assume you wrote correct color code.

```bash
# some color examples
    "255,255,255"   # White
    "255,0,0"       # Red
    "0,255,255"     # Cyan
```

Will add hex color support and alias for color very soon

## Future Improvements

- fix every bug/edge cases i know
- windows support
- hex color and color support
- resize window
- star fading animation
- moving star animation

## LICENSE

Yeah, It's [MIT](LICENSE)
