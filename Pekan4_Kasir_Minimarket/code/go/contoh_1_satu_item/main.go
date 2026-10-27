// Contoh 1 - satu barang, satu keputusan (if / else)
// Pembanding untuk 01a_contoh_satu_item.py. Tidak wajib dikerjakan.
//
// Jalankan: go run ./contoh_1_satu_item

package main

import "fmt"

func main() {
	// Tipe ditulis eksplisit dengan var. Go juga bisa memakai :=
	// bila tipenya dapat disimpulkan dari nilai.
	var harga int = 120_000
	var jumlah int = 1
	var subtotal int = harga * jumlah

	var diskon float64
	if subtotal > 100_000 {
		diskon = 0.10 * float64(subtotal)
	} else {
		diskon = 0
	}

	fmt.Printf("Subtotal : Rp%d\n", subtotal)
	fmt.Printf("Diskon   : Rp%.0f\n", diskon)
}
