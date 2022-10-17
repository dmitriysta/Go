package RussianHen

type RussianHen struct {
  eggs int
  description string
  country string
}

func (r *RussianHen) getCountOfEggsPerMonth() int {
  return r.eggs
}

func (r *RussianHen) getDescription() string {
  return fmt.Println(r.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", r.country)
}
