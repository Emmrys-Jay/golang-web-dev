package main

import (
	"log"
	"os"
	"text/template"
)

type calHotels struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func main() {
	var tpl *template.Template

	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

	hotels := []calHotels{
		calHotels{
			"Moonstone Landing",
			"6.8 mi from Heart Castle",
			"Cambria",
			"10111",
			"Southern",
		},

		calHotels{
			"Hotel Corque",
			"5 minutes walk from Old Mission Santa Ines",
			"Solvang",
			"10112",
			"Central",
		},

		calHotels{
			"Hotel Mission De Oro",
			"Hotel in Santa Nella",
			"Santa Nella",
			"10113",
			"Northern",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
