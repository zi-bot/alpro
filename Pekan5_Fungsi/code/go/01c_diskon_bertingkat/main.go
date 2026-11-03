// 01c — Diskon bertingkat: satu fungsi, empat tingkat (versi Go).
// Jalankan: go run 01c_diskon_bertingkat.go
//
// Aturan (batas bawah inklusif, batas atas milik tingkat berikutnya):
//   < Rp100.000                 -> 0%
//   Rp100.000 - Rp249.999       -> 5%
//   Rp250.000 - Rp499.999       -> 10%
//   >= Rp500.000                -> 15%

package main

import (
	"fmt"
	"strconv"
)

func hitungDiskon(subtotal int) int {
	if subtotal < 0 {
		panic("subtotal tidak boleh negatif")
	}
	if subtotal < 100000 {
		return 0
	} else if subtotal < 250000 {
		return subtotal * 5 / 100
	} else if subtotal < 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

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

func persen(subtotal int) string {
	if subtotal < 100000 {
		return "0%"
	} else if subtotal < 250000 {
		return "5%"
	} else if subtotal < 500000 {
		return "10%"
	}
	return "15%"
}

func main() {
	uji := []int{0, 53000, 99999, 100000, 249999, 250000, 499999, 500000, 1250000}

	fmt.Println("SUBTOTAL      TINGKAT   DISKON")
	fmt.Println("----------------------------------")
	for _, subtotal := range uji {
		fmt.Printf("%13s  %6s  %10s\n",
			rupiah(subtotal),
			persen(subtotal),
			rupiah(hitungDiskon(subtotal)),
		)
	}

	fmt.Println()
	fmt.Println("Dua batas yang paling mudah salah:")
	fmt.Printf("   99.999 -> %s  (di bawah batas, tetap 0%%)\n", rupiah(hitungDiskon(99999)))
	fmt.Printf("  100.000 -> %s  (tepat di batas, sudah 5%%)\n", rupiah(hitungDiskon(100000)))
	fmt.Println()
	fmt.Println("Pecahan rupiah dibuang ke bawah, bukan dibulatkan:")
	fmt.Printf("  249.999 x 5%% = 12.499,95 -> %s\n", rupiah(hitungDiskon(249999)))
}
