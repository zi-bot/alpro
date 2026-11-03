// main.go — pembanding Go untuk kasus yang sama.
// Jalankan: go run main.go
//
// Yang berbeda dari Python:
//   - tipe parameter ditulis di belakang nama (subtotal int), tipe nilai
//     balik ditulis setelah tanda kurung
//   - kata kunci return wajib, termasuk pada cabang terakhir
//   - prosedur adalah fungsi tanpa nilai balik sama sekali

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

func hitungDiskon(subtotal int) int {
	if subtotal < 100000 {
		return 0
	} else if subtotal < 250000 {
		return subtotal * 5 / 100
	} else if subtotal < 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

func hitungTotal(subtotal int, diskon int) int {
	dasar := subtotal - diskon
	return dasar + dasar*11/100
}

// Prosedur: tidak ada tipe nilai balik di antara kurung tutup dan kurung kurawal.
func tampilkanTotal(nama string, subtotal int, diskon int, total int) {
	fmt.Printf("%s  subtotal %s  diskon %s  total %s\n",
		nama, rupiah(subtotal), rupiah(diskon), rupiah(total))
}

func main() {
	kasus := [][2]interface{}{
		{"K-01", 50000},
		{"K-02", 99999},
		{"K-03", 100000},
		{"K-04", 100001},
		{"K-05", 249999},
		{"K-06", 250000},
		{"K-07", 499999},
		{"K-08", 500000},
		{"K-09", 500000},
		{"K-10", 0},
	}

	for _, k := range kasus {
		nama := k[0].(string)
		subtotal := k[1].(int)
		diskon := hitungDiskon(subtotal)
		total := hitungTotal(subtotal, diskon)
		tampilkanTotal(nama, subtotal, diskon, total)
	}
}
