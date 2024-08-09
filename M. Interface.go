package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type kotak struct {
	panjang float64
}
func (k kotak) luas() float64  {
	return k.panjang * k.panjang
}
func (k kotak) keliling() float64  {
	return  k.panjang * 4
}


type bulat struct {
	diameter float64
}
func (b bulat) luas()  float64{
	return math.Pi * (b.diameter/2) * (b.diameter/2)
}
func (b bulat) keliling() float64  {
	return math.Pi + b.diameter
}

func main() {
	var bangunDatar hitung

	bangunDatar = kotak{10}

	fmt.Println("-----------")
	fmt.Println("Luas : ",bangunDatar.luas())
	fmt.Println("keliling : ",bangunDatar.keliling())
	fmt.Println("-----------")
}