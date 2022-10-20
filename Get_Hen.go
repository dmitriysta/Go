package main

import "fmt"

type Country struct {
	name string
}

func (c *Country) getHen() {
	Russia := Country{"Russia"}
	Ukraine := Country{"Ukraine"}
	Moldova := Country{"Moldova"}
	Belarusia := Country{"Belarusia"}
	nameCountry := Country{"Принимаем значение страны извне"}
	switch {
	case nameCountry == Russia:
		RussianHen.getDescription()
	case nameCountry == Ukraine:
		UkrainianHen.getDescription()
	case nameCountry == Moldova:
		MoldovanHen.getDescription()
	case nameCountry == Belarusia:
		BelarusianHen.getDescription()
	default:
		fmt.Println(nameCountry)
	}
}
