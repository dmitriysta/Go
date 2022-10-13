

type UkrainianHen struct {
  eggs int
  description string
  country string
}

func (u *UkrainianHen) getCountOfEggsPerMonth() int {
  return u.eggs
}

func (u *UkrainianHen) getDescription() string {
  return fmt.Println(u.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", u.country)
}
