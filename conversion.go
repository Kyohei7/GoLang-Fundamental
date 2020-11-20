/*
	Konversi Tipe Data
	~ untuk melakukan perubahan tipe data
	~ contohnya mengubah tipe data dari int32 menjadi int64
*/

package main

import "fmt"

func main() {

	// Tipe Data Awal yaitu int32
	var integer32 int32 = 30000

	// Melakukan Perubahan menjadi int64
	var integer64 int64 = int64(integer32)
	fmt.Println(integer64)

	// Melakukan Perubahan menjadi int16
	var integer16 int16 = int16(integer32)
	fmt.Println(integer16)

	// Konversi kedua -> Merubah byte menjadi string
	var nama = "Muhammad Rizki"
	var e byte = nama[0]
	var eString = string(e)

	fmt.Println(nama)    // Output : Muhammad Rizki
	fmt.Println(eString) // Output : M

}
