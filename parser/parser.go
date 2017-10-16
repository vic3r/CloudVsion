package parser

import (
	"fmt"
	"strings"
)

func Parse_Name(to_parse string) []string {
	to_parse = strings.Split(to_parse, "NOMBRE\n")[1]
	to_parse = strings.Split(to_parse, "\nDOMICILIO")[0]
	return strings.Split(to_parse, "\n")
}

func Parse_Sex(to_parse string) string {
	to_parse = strings.Split(to_parse, "SEXO ")[1]
	to_parse = strings.Split(to_parse, "\n")[0]
	return to_parse
}

func Parse_Zip(to_parse string) string {
	to_parse = strings.Split(to_parse, "DOMICILIO\n")[1]
	to_parse = strings.Split(to_parse, "\nFOLIO")[0]
	to_parse = strings.Split(to_parse, "\n")[1]
	temp := strings.Split(to_parse, " ")
	zip := temp[len(temp)-1]
	return zip
}

func Parse_JSON(data string){

	full_name := Parse_Name(data)
	sex := Parse_Sex(data)
	zip := Parse_Zip(data)

	var name string
	var last_nameP string
	var last_nameF string

	if len(full_name) == 3 {
		name = full_name[2]
		last_nameP = full_name[0]
		last_nameF = full_name[1]
	} else if len(full_name) == 2 {
		name = full_name[1]
		last_nameP = full_name[0]
		last_nameF = " "
	}

	if sex == "H" {
		sex = "Male"
	} else if sex == "M" {
		sex = "Female"
	} else {
		sex = ""
	}

	fmt.Println("Name: ", strings.Title(strings.ToLower(name)))
	fmt.Println("Last_NameP: ", strings.Title(strings.ToLower(last_nameP)))
	fmt.Println("Last_NameF: ", strings.Title(strings.ToLower(last_nameF)))
	fmt.Println("Sex: ", sex)
	fmt.Println("ZIP Code: ", zip)
}





