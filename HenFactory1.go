
package main

import "fmt"

type Hen interface {
  getCountOfEggsPerMonth() int
  getDescription() string

}

// catalog/RussianHen.go

type RussianHen struct {
  eggs int
  description string
  country string
}

func (R *RussianHen) getCountOfEggsPerMonth() int {
  return R.eggs
}

func (R *RussianHen) getDescription() string {
  return fmt.Println(R.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", R.country)
}

// catalog/UkrainianHen.go

type UkrainianHen struct {
  eggs int
  description string
  country string
}

func (U *UkrainianHen) getCountOfEggsPerMonth() int {
  return U.eggs
}

func (U *UkrainianHen) getDescription() string {
  return fmt.Println(U.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", U.country)
}

// catalog/MoldovanHen.go

type MoldovanHen struct {
  eggs int
  description string
  country string
}

func (M *MoldovanHen) getCountOfEggsPerMonth() int {
  return M.eggs
}

func (M *MoldovanHen) getDescription() string {
  return fmt.Println(M.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", M.country)
}

// catalog/BelarusianHen.go

type BelarusianHen struct {
  eggs int
  description string
  country string
}

func (B *BelarusianHen) getCountOfEggsPerMonth() int {
  return B.eggs
}

func (B *BelarusianHen) getDescription() string {
  return fmt.Println(B.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", B.country)
}

// catalog/getHen.go

type Country struct {
  nameCountry string
  nameHen string
}

func main() {
  Russia := {
    nameCountry: "Russia",
    nameHen: "RussianHen",
  }
  Ukraine := {
    nameCountry: "Ukraine",
    nameHen: "UkrainianHen",
  }
Moldova := {
  nameCountry: "Moldova",
  nameHen: "MoldovanHen",
}
Belarusia := {
  nameCountry: "Belarusia",
  nameHen: "BelarusianHen",
}



func (c *Country) getHen() {
  fmt.Println(c.nameHen)
}
