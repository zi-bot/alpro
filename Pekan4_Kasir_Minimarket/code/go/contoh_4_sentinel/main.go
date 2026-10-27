// Contoh 4 - perulangan while dengan penanda berhenti
// Pembanding untuk 01d_contoh_sentinel.py. Tidak wajib dikerjakan.
//
// Jalankan: go run ./contoh_4_sentinel

package main

import "fmt"

func main() {
	// Pasangan (harga, jumlah), diakhiri harga -1 sebagai penanda berhenti.
	masukan := []int{15_000, 2, 5_000, 3, 8_000, 1, -1}

	subtotal := 0
	i := 0

	for masukan[i] >= 0 {
		harga := masukan[i]
		jumlah := masukan[i+1]
		subtotal += harga * jumlah
		i += 2
	}

	fmt.Printf("Subtotal : Rp%d\n", subtotal)
}
