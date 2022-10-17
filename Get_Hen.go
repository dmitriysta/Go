package main



type Country struct {
  Russia string
  Ukraine string
  Moldova string
  Belarusia string
}



func (c *Country) getHen() {
 c.Russia = "Russia"
 c.Ukraine = "Ukraine"
 c.Moldova = "Moldova"
 c.Belarusia = "Belarusia"
 var nameCountry string //здесь мы принимаем значение переменной извне
  if nameCountry == c.Russia {
    return RussianHen.getDescription()
  }

  if nameCountry == c.Ukraine {
    return UkrainianHen.getDescription()
  }

  if nameCountry == c.Moldova {
    return MoldovanHen.getDescription()
  }

  if nameCountry == c.Belarusia {
    return BelarusianHen.getDescription()
  } else {
    return nil
  }
}
