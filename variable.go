/*
	Variabel di GoLang
	~ Tempat untuk menyimpan Data
	~ Digunakan untuk mengakses Data
	~ Setiap variabel yang dibuat dengan tipe data yang berbeda
	  diharuskan untuk membuatnya menjadi beberapa variabel
	~ Kata kuncinya : var dan diikutin dengan tipe Datanya (tidak wajib)
	~ Kata kunci kedua : menggunakan karakter :=

	Deklarasi Multiple Variabel
	~ konsep yang digunakan untuk membuat variabel dengan banyak dalam
	  dalam 1 waktu
	~ gunanya untuk membuat code menjadi rapi & mudah dilihat

*/

package main

import "fmt"

func main() {

	// Cara ke-1 untuk membuat Variabel
	var nama string
	nama = "Muhammad Rizki"
	fmt.Println(nama)
	// Mengubah Nilai Dari Variabel
	nama = "Ethan Ramadhan"
	fmt.Println(nama)

	// Cara ke-2 untuk membuat Variabel
	var namaLengkap = "Muhammad Rizki"
	fmt.Println(namaLengkap)
	namaLengkap = "Ethan Ramadhan"
	fmt.Println(namaLengkap)

	// Cara ke-3 untuk membuat Variabel
	umur := 24
	fmt.Println(umur)
	// Mengubah Nilai dari Variabel
	umur = 26
	fmt.Println(umur)

	// Cara ke-4 untuk membuat Variabel
	var (
		namaDepan    = "Muhammad"
		namaBelakang = "Rizki"
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}
