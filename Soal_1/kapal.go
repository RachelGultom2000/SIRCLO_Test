package kapal

import "fmt"

type kapal struct {
	nama    string
	jenis   string
	panjang int
	lebar   int
}

func (k kapal) getName() string {
	return k.nama
}

func (k kapal) GetKapalJenis() string {
	return fmt.Sprintf("%s itu berjenis %s dan memiliki panjang %d serta lebar %d", k.getName(), k.jenis, k.panjang, k.lebar)
}

func NewKapal(nama string, jenis string, panjang int, lebar int) kapal {
	return kapal{nama: nama, jenis: jenis, panjang: panjang, lebar: lebar}
}
