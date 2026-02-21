from dataclasses import dataclass


@dataclass
class Star:
    symbol: str = ""
    duration: int = 0
    bg_color: str = ""
    fg_color: str = ""


STARS: list[Star] = [Star() for _ in range(5)]

STAR_COUNT: int = 15
SYMBOL: str = "."

DURATION: int = 5

BG_COLOR: str = ""
FG_COLOR: str = ""
