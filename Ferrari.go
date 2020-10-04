package main

import "strconv"

type Araba struct {
	Marka string
	Model string
	Renk  string
	Hız   int
	Fiyat int
}

type OzelUretim struct {
	Ozel bool
}

type Ferrari struct {
	//Kalıtım alma
	Araba
	OzelUretim
}

type Arabaface interface {
	Calıs()
	Dur()
	Bilgi()
}

func (_ Ferrari) Calis() bool {

	return true
}
func (_ Ferrari) Dur() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Marka + "\n" + x.Model + "\n" + x.Renk + "\n" + strconv.Itoa(x.Hız) + "\n" + strconv.Itoa(x.Fiyat)
	add := "Yes"
	if x.Ozel {
		ret += "\n" + add
	}
	return ret
}
