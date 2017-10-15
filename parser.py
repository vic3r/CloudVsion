import numpy as np
import pandas as pd
import json

to_parse = "INSTITUTO FEDES LELECTORAL\nREGISTRO FEDERAL DE ELECTORES\nCREDENCIAL PARA VO AR\nUNIDOS\nNOMBRE\nVAZQUEZ\nESPINOZA\nJOSE LUIS\nDOMICILIO\nAVINDEPENDENCIAL 9 M 25\nCOL MODERNA 85330\nEMPALME ,SON.\nFOLIO 1226042401988 ANO DE REGISTR\nCLAVE DE ELECTOR VZESLS94011226H80\nCURP VAEL940112HSRZSS08\nESTADO 26 MUNICIPIO 060\nEDA\nSEXO s\n2012 :00\nLOCALIDAD 0001 SECCION 1011\nEMISIÓN 2012 VIGENCIA HASTA 2022\nFIRMA\n"
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
