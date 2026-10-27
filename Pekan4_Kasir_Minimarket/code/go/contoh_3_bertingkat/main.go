// Contoh 3 - diskon bertingkat + pajak (if / else if / else)
// Pembanding untuk 01c_contoh_bertingkat.py. Tidak wajib dikerjakan.
//
// Jalankan: go run ./contoh_3_bertingkat

package main

import "fmt"

func main() {
	keranjang := [][2]int{{60_000, 2}, {70_000, 2}}

	subtotal := 0
	for _, item := range keranjang {
		subtotal += item[0] * item[1]
	}

	var diskon float64
	if subtotal < 100_000 {
		diskon = 0
	} else if subtotal < 250_000 {
		diskon = 0.05 * float64(subtotal)
	} else {
		diskon = 0.10 * float64(subtotal)
	}

	pajak := 0.11 * (float64(subtotal) - diskon)
	total := (float64(subtotal) - diskon) + pajak

	fmt.Printf("Subtotal : Rp%d\n", subtotal)
	fmt.Printf("Diskon   : Rp%.0f\n", diskon)
	fmt.Printf("Pajak    : Rp%.0f\n", pajak)
	fmt.Printf("Total    : Rp%.0f\n", total)
}
