package main

import "strconv"

type Ferrari struct {
	//Kalıtım alma
	Araba
	OzelUretim
}

func (_ Ferrari) Calis() bool {
	return true
}

func (_ Ferrari) Dur() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Marka + "\n" + x.Model + "\n" + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.Itoa(x.Fiyat)
	add := "Yes"
	if x.Ozel {
		ret += "\n" + add
	}
	return ret
}
