// latihan_template.go — Praktik (versi Go).
// Jalankan: go run latihan_template.go
//
// Tugas Anda: isi empat TODO di bawah sehingga keluaran program ini PERSIS
// sama dengan keluaran 02a_sebelum_refaktor.go.
//
// Aturan:
//   - Jangan mengubah main() dan jangan mengubah teks yang dicetak.
//   - Setiap fungsi hanya mengerjakan satu hal.
//   - Fungsi yang menghitung tidak boleh mencetak; yang mencetak tidak boleh
//     menghitung sendiri.

package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	Harga  int
	Jumlah int
}

type Transaksi struct {
	Nama string
	Isi  []Item
}

var keranjang = []Transaksi{
	{
		Nama: "A",
		Isi: []Item{
			{Harga: 15000, Jumlah: 2},
			{Harga: 5000, Jumlah: 3},
			{Harga: 8000, Jumlah: 1},
		},
	},
	{
		Nama: "B",
		Isi: []Item{
			{Harga: 125000, Jumlah: 2},
			{Harga: 250000, Jumlah: 1},
		},
	},
	{
		Nama: "C",
		Isi: []Item{
			{Harga: 200000, Jumlah: 1},
			{Harga: 100000, Jumlah: 1},
		},
	},
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

// TODO 1 — jumlahkan harga x jumlah untuk seluruh baris keranjang.
func hitungSubtotal(isi []Item) int {
	panic("TODO 1: hitungSubtotal belum diisi")
}

// TODO 2 — kembalikan diskon dalam rupiah.
// Aturan (batas bawah inklusif):
//
//	< Rp100.000             -> 0%
//	Rp100.000 - Rp249.999   -> 5%
//	Rp250.000 - Rp499.999   -> 10%
//	>= Rp500.000            -> 15%
func hitungDiskon(subtotal int) int {
	panic("TODO 2: hitungDiskon belum diisi")
}

// TODO 3 — PPN 11% dari nilai setelah diskon.
func hitungPpn(dasar int) int {
	panic("TODO 3: hitungPpn belum diisi")
}

// TODO 4 — prosedur: cetak satu nota, tidak mengembalikan nilai.
func cetakNota(nama string, isi []Item) {
	panic("TODO 4: cetakNota belum diisi")
}

func main() {
	totalTerbesar := 0
	namaTerbesar := ""

	for _, k := range keranjang {
		cetakNota(k.Nama, k.Isi)
		subtotal := hitungSubtotal(k.Isi)
		diskon := hitungDiskon(subtotal)
		total := subtotal - diskon
		total += hitungPpn(total)

		if total > totalTerbesar {
			totalTerbesar = total
			namaTerbesar = k.Nama
		}
	}

	fmt.Printf("Transaksi terbesar: %s dengan total %s\n", namaTerbesar, rupiah(totalTerbesar))
}
