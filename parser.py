import numpy as np
import pandas as pd
import json
import sys

print("hola")
to_parse = sys.argv[1]
full_name = to_parse.split("NOMBRE\n")[1].split("\nDOMICILIO")[0].split("\n")
sex = to_parse.split("SEXO ")[1].split("\n")[0]
zip = to_parse.split("DOMICILIO\n")[1].split("\nFOLIO")[0].split("\n")[1].split(" ")[-1]

if len(full_name) == 3:
    name = full_name[-1].title()
    last_nameP = full_name[0].title()
    last_nameF = full_name[1].title()
elif len(name) == 2:
    name = full_name[-1].title()
    last_nameP = full_name[0].title()

if sex ==  "H":
    sex = "Male"
elif sex == "M":
    sex = "Female"
else:
    sex = ""

print(json.dumps({'name': name,
                  'last_nameP': last_nameP,
                  'last_nameF': last_nameF,
                  'sex': sex,
                  'zip': zip }, indent=4, separators=(',', ': ')))
