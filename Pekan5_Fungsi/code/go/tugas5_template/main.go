// tugas5_template.go — Tugas 5 (individu, rubrik A2, 5% nilai akhir, versi Go).
// Jalankan: go run tugas5_template.go
//
// Kerjakan empat TODO di bawah.

package main

import (
	"strconv"
)

type Item struct {
	Harga  int
	Jumlah int
}

type Kasus struct {
	Nama string
	Isi  []Item
}

var kasusUji = []Kasus{
	{"K-01", []Item{{Harga: 50000, Jumlah: 1}}},
	{"K-02", []Item{{Harga: 99999, Jumlah: 1}}},
	{"K-03", []Item{{Harga: 100000, Jumlah: 1}}},
	{"K-04", []Item{{Harga: 100001, Jumlah: 1}}},
	{"K-05", []Item{{Harga: 249999, Jumlah: 1}}},
	{"K-06", []Item{{Harga: 250000, Jumlah: 1}}},
	{"K-07", []Item{{Harga: 499999, Jumlah: 1}}},
	{"K-08", []Item{{Harga: 500000, Jumlah: 1}}},
	{"K-09", []Item{{Harga: 125000, Jumlah: 2}, {Harga: 250000, Jumlah: 1}}},
	{"K-10", []Item{{Harga: 0, Jumlah: 0}}},
}

// TODO 1 — tolak subtotal negatif.
// Kembalikan subtotal apa adanya bila sah; panic dengan pesan
// "subtotal tidak boleh negatif" bila tidak sah.
func validasiSubtotal(subtotal int) int {
	panic("TODO 1: validasiSubtotal belum diisi")
}

// TODO 2 — diskon dalam rupiah, batas bawah inklusif.
// < Rp100.000             -> 0%
// Rp100.000 - Rp249.999   -> 5%
// Rp250.000 - Rp499.999   -> 10%
// >= Rp500.000            -> 15%
func hitungDiskon(subtotal int) int {
	panic("TODO 2: hitungDiskon belum diisi")
}

// TODO 3 — total bayar: nilai setelah diskon ditambah PPN 11%.
func hitungTotal(subtotal int, diskon int) int {
	panic("TODO 3: hitungTotal belum diisi")
}

// TODO 4 — prosedur penyajian. Cetak format:
//
//	K-01  subtotal Rp50.000  diskon Rp0  total Rp55.500
func tampilkanTotal(nama string, subtotal int, diskon int, total int) {
	panic("TODO 4: tampilkanTotal belum diisi")
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

func subtotalDari(isi []Item) int {
	subtotal := 0
	for _, it := range isi {
		subtotal += it.Harga * it.Jumlah
	}
	return subtotal
}

func main() {
	for _, k := range kasusUji {
		subtotal := validasiSubtotal(subtotalDari(k.Isi))
		diskon := hitungDiskon(subtotal)
		total := hitungTotal(subtotal, diskon)
		tampilkanTotal(k.Nama, subtotal, diskon, total)
	}
}
