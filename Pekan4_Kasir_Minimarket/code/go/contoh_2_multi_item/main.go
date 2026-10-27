// Contoh 2 - banyak barang, perulangan + penjumlahan
// Pembanding untuk 01b_contoh_multi_item.py. Tidak wajib dikerjakan.
//
// Jalankan: go run ./contoh_2_multi_item

package main

import "fmt"

func main() {
	// [][2]int adalah daftar pasangan (harga, jumlah).
	keranjang := [][2]int{{15_000, 2}, {5_000, 3}, {8_000, 1}}

	subtotal := 0
	for _, item := range keranjang {
		subtotal += item[0] * item[1]
	}

	fmt.Printf("Subtotal : Rp%d\n", subtotal)
}
