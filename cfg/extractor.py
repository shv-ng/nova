import json
from os import path
import sys
import dataclasses
import importlib.util

file_path = path.abspath(sys.argv[1])


def import_from_path(module_name, file_path):
    spec = importlib.util.spec_from_file_location(module_name, file_path)
    module = importlib.util.module_from_spec(spec)
    sys.modules[module_name] = module
    spec.loader.exec_module(module)
    return module


name = path.basename(file_path)

mod = import_from_path(name, file_path)

out = {}

out["star_count"] = mod.STAR_COUNT
out["symbol"] = mod.SYMBOL
out["duration"]=mod.DURATION
out["bg_color"]=mod.BG_COLOR
out["fg_color"]=mod.FG_COLOR

out["stars"] = [dataclasses.asdict(x) for x in mod.STARS]

print(json.dumps(out))
