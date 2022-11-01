package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peserta struct {
	nama, alamat, pekerjaan, alasan string
}

func main() {

	var listPeserta []Peserta = []Peserta{
		{
			nama:      "jamal",
			alamat:    "Jln. Harapan Indah No 3 Blok A1",
			pekerjaan: "progamer",
			alasan:    "Ingin jadi jagoan!",
		},
		{
			nama:      "udin",
			alamat:    "Jln. Salam No 1 Blok C4",
			pekerjaan: "backenc engineer",
			alasan:    "Ingin jadi yang paling hebat",
		},
		{
			nama:      "satrio",
			alamat:    "Jln. Kemanggi No 10 Blok B9",
			pekerjaan: "Front End Engineerr",
			alasan:    "tambah ilmu aja",
		},
		{
			nama:      "Saiful",
			alamat:    "Jln. Planet No 1 Blok Uranus 3",
			pekerjaan: "mahasiswa",
			alasan:    "isi waktu luang",
		},
		{
			nama:      "jamal",
			alamat:    "Jln. Harapan Indah No 3 Blok A1",
			pekerjaan: "progamer",
			alasan:    "Ingin jadi jagoan!",
		},
	}

	idxPeserta, err := strconv.Atoi(os.Args[1])

	if err == nil {
		var jumlahPeserta = len(listPeserta)
		idxPeserta = idxPeserta - 1
		if jumlahPeserta > idxPeserta {
			fmt.Println("Nama: ", listPeserta[idxPeserta].nama)
			fmt.Println("Alamat: ", listPeserta[idxPeserta].alamat)
			fmt.Println("Pekerjaan: ", listPeserta[idxPeserta].pekerjaan)
			fmt.Println("Alasan memilih kelas Golang: ", listPeserta[idxPeserta].alasan)
		} else {
			fmt.Println("Tidak ada absen no: ", idxPeserta+1)
		}
	} else {
		fmt.Println("Salah masukin no absen")
	}
}
