package main

import "fmt"

type tStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (t *tStruct) Shoot() bool {
	if t.On == false {
		return false
	} else {
		if t.Ammo > 0 {
			t.Ammo--
			return true
		} else {
			return false
		}
	}
}

func (t *tStruct) RideBike() bool {
	switch {
	case t.On == false:
		return false
	case t.Power > 0:
		t.Power--
		return true
	default:
		return false
	}
}

func main() {
	testStruct := tStruct{true, 1, 2} //получаем значения
	fmt.Println(testStruct.Shoot(), testStruct.RideBike())
}
