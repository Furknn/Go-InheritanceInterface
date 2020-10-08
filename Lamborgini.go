package main

import "strconv"

type Lamborgini struct {
	//Kalıtım alma
	Araba
	OzelUretim
}

func (_ Lamborgini) Calis() bool {
	return true
}

func (_ Lamborgini) Dur() bool {
	return true
}

func (x *Lamborgini) Information() string {
	ret := "\t" + x.Marka + "\n" + x.Model + "\n" + x.Renk + "\n" + strconv.Itoa(x.Hiz) + "\n" + strconv.Itoa(x.Fiyat)
	add := "Yes"
	if x.Ozel {
		ret += "\n" + add
	}
	return ret
}
