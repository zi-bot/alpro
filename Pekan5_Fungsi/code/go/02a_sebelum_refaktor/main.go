// 02a — Program sebelum dipecah: semuanya di satu alur (versi Go).
// Jalankan: go run 02a_sebelum_refaktor.go
//
// Program ini BENAR. Jangan diubah dulu — catat keluarannya, karena keluaran
// itulah bukti bahwa hasil refaktor Anda nanti tidak merusak apa pun.

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

func main() {
	keranjang := []Transaksi{
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

	totalTerbesar := 0
	namaTerbesar := ""

	for _, k := range keranjang {
		subtotal := 0
		for _, it := range k.Isi {
			subtotal += it.Harga * it.Jumlah
		}

		var diskon int
		if subtotal < 100000 {
			diskon = 0
		} else if subtotal < 250000 {
			diskon = subtotal * 5 / 100
		} else if subtotal < 500000 {
			diskon = subtotal * 10 / 100
		} else {
			diskon = subtotal * 15 / 100
		}

		dasar := subtotal - diskon
		ppn := dasar * 11 / 100
		total := dasar + ppn

		fmt.Printf("NOTA KERANJANG %s\n", k.Nama)
		fmt.Println("----------------------------------")
		for _, it := range k.Isi {
			fmt.Printf("  %11s x %d = %11s\n",
				rupiah(it.Harga),
				it.Jumlah,
				rupiah(it.Harga*it.Jumlah),
			)
		}
		fmt.Println("----------------------------------")
		fmt.Printf("  %-20s%14s\n", "Subtotal", rupiah(subtotal))
		fmt.Printf("  %-20s%14s\n", "Diskon", rupiah(diskon))
		fmt.Printf("  %-20s%14s\n", "PPN 11%", rupiah(ppn))
		fmt.Printf("  %-20s%14s\n", "TOTAL", rupiah(total))
		fmt.Println()

		if total > totalTerbesar {
			totalTerbesar = total
			namaTerbesar = k.Nama
		}
	}

	fmt.Printf("Transaksi terbesar: %s dengan total %s\n", namaTerbesar, rupiah(totalTerbesar))
}
