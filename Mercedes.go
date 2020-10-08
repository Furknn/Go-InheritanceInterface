package main

import "strconv"

type Mercedes struct {
	//Kalıtım alma
	Araba
	OzelUretim
}

func (_ Mercedes) Calis() bool {
	return true
}

func (_ Mercedes) Dur() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Marka + "\n" + x.Model + "\n" + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.Itoa(x.Fiyat)
	add := "Yes"
	if x.Ozel {
		ret += "\n" + add
	}
	return ret
}
