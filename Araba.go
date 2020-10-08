package main

type Araba struct {
	Marka string
	Model string
	Renk  string
	Hiz   int
	Fiyat int
}

type OzelUretim struct {
	Ozel bool
}

type Arabaface interface {
	Calis()
	Dur()
	Bilgi()
}
