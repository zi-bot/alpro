// 01a — Satu fungsi, satu nilai balik (versi Go).
// Jalankan: go run 01a_satu_fungsi.go
//
// Perhatikan tiga hal pada hitungSubtotal:
//   1. func    — fungsi didefinisikan dengan kata kunci func.
//   2. parameter harga dan jumlah bertipe int (tipe ditulis di belakang).
//   3. nilai balik bertipe int; kata kunci return wajib ditulis.

package main

import (
	"fmt"
	"strconv"
)

func rupiah(nilai int) string {
	teks := strconv.Itoa(nilai)
	if nilai < 0 {
		teks = teks[1:]
	}
	keluar := ""
	for i, c := range teks {
		if i > 0 && (len(teks)-i)%3 == 0 {
			keluar += "."
		}
		keluar += string(c)
	}
	if nilai < 0 {
		return "Rp-" + keluar
	}
	return "Rp" + keluar
}

func hitungSubtotal(harga int, jumlah int) int {
	return harga * jumlah
}

func main() {
	harga := 15000
	jumlah := 2
	subtotal := hitungSubtotal(harga, jumlah)

	fmt.Println("harga    :", rupiah(harga))
	fmt.Println("jumlah   :", jumlah)
	fmt.Println("subtotal :", rupiah(subtotal))
}
