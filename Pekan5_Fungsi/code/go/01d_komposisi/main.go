// 01d — Kasus utuh: lima fungsi, satu alur utama (versi Go).
// Jalankan: go run 01d_komposisi.go
//
// Pembagian tanggung jawab:
//   hitungSubtotal(keranjang)      -> int   menjumlahkan isi keranjang
//   hitungDiskon(subtotal)         -> int   memutuskan tingkat diskon
//   hitungPpn(dasar)               -> int   pajak 11% dari nilai setelah diskon
//   hitungTotal(subtotal, diskon)  -> int   memanggil hitungPpn, lalu menjumlah
//   tampilkanHasil(...)            -> void  prosedur: hanya mencetak

package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	Harga  int
	Jumlah int
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

func hitungSubtotal(keranjang []Item) int {
	subtotal := 0
	for _, it := range keranjang {
		subtotal += it.Harga * it.Jumlah
	}
	return subtotal
}

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

func hitungPpn(dasar int) int {
	return dasar * 11 / 100
}

func hitungTotal(subtotal int, diskon int) int {
	dasar := subtotal - diskon
	return dasar + hitungPpn(dasar)
}

func tampilkanHasil(nama string, keranjang []Item) {
	subtotal := hitungSubtotal(keranjang)
	diskon := hitungDiskon(subtotal)
	total := hitungTotal(subtotal, diskon)

	fmt.Printf("NOTA %s\n", nama)
	fmt.Println("----------------------------------")
	for _, it := range keranjang {
		fmt.Printf("  %11s x %d = %11s\n",
			rupiah(it.Harga),
			it.Jumlah,
			rupiah(it.Harga*it.Jumlah),
		)
	}
	fmt.Println("----------------------------------")
	fmt.Printf("  %-20s%14s\n", "Subtotal", rupiah(subtotal))
	fmt.Printf("  %-20s%14s\n", "Diskon", rupiah(diskon))
	fmt.Printf("  %-20s%14s\n", "PPN 11%", rupiah(hitungPpn(subtotal-diskon)))
	fmt.Printf("  %-20s%14s\n", "TOTAL", rupiah(total))
	fmt.Println()
}

func main() {
	keranjangA := []Item{
		{Harga: 15000, Jumlah: 2},
		{Harga: 5000, Jumlah: 3},
		{Harga: 8000, Jumlah: 1},
	}
	keranjangB := []Item{
		{Harga: 125000, Jumlah: 2},
		{Harga: 250000, Jumlah: 1},
	}

	tampilkanHasil("KERANJANG A", keranjangA)
	tampilkanHasil("KERANJANG B", keranjangB)
}
