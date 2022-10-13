

func (c *Country) getHen() {
  nameCountry := null

  if nameCountry == c.Russia {
    return RussianHen
  }

  if nameCountry == c.Ukraine {
    return UkrainianHen
  }

  if nameCountry == c.Moldova {
    return MoldovanHen
  }

  if nameCountry == c.Belarusia {
    return BelarusianHen
  }

  else {
    return nameCountry
  }
