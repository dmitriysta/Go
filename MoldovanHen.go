

type MoldovanHen struct {
  eggs int
  description string
  country string
}

func (m *MoldovanHen) getCountOfEggsPerMonth() int {
  return m.eggs
}

func (m *MoldovanHen) getDescription() string {
  return fmt.Println(m.description + "Моя страна - %s."  + "Я несу " + getCountOfEggsPerMonth() + "яиц в месяц", m.country)
}
