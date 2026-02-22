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
