package main

type BelarusianHen struct {
  eggs int
  description string
  country string
}

func (b *BelarusianHen) getCountOfEggsPerMonth() int {
  return B.eggs
}

func (b *BelarusianHen) getDescription() string {
  return fmt.Println(b.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", b.country)
}
