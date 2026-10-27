// Latihan terbimbing - kasir minimarket (versi Go)
// Pembanding untuk latihan_template.py. Tidak wajib dikerjakan.
//
// Jalankan program: go run ./latihan
// Jalankan uji    : go test ./latihan
//
// Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
// Kunci ada di solusi/latihan/main.go.

package main

import "fmt"

type Hasil struct {
	subtotal int
	diskon   int
	pajak    int
	total    int
	ditolak  int
}

func hitung(keranjang [][2]int) Hasil {
	subtotal := 0
	ditolak := 0

	for _, item := range keranjang {
		harga := item[0]
		jumlah := item[1]

		// TODO 1 - Validasi baris
		// Bila jumlah <= 0 atau harga < 0:
		//   tambahkan 1 ke ditolak, lalu continue.

		subtotal += harga * jumlah
	}

	// TODO 2 - Diskon bertingkat
	// Ganti 0 dengan persen diskon: 0, 5, 10, atau 15.
	// Batas: < 100.000, < 250.000, < 500.000, sisanya.
	persen := 0

	// Pembulatan setengah ke atas: tambah 50 lalu bagi 100.
	diskon := (subtotal*persen + 50) / 100

	// TODO 3 - Pajak
	// Ganti 0 dengan ((subtotal - diskon) * 11 + 50) / 100
	pajak := 0

	// TODO 4 - Total
	// Ganti 0 dengan subtotal - diskon + pajak
	total := 0

	return Hasil{
		subtotal: subtotal,
		diskon:   diskon,
		pajak:    pajak,
		total:    total,
		ditolak:  ditolak,
	}
}

// Format ribuan: 58830 menjadi 58.830
func rupiah(nilai int) string {
	if nilai < 0 {
		return "-" + rupiah(-nilai)
	}
	teks := fmt.Sprintf("%d", nilai)
	hasil := ""

	for i := 0; i < len(teks); i++ {
		if i > 0 && (len(teks)-i)%3 == 0 {
			hasil += "."
		}
		hasil += string(teks[i])
	}

	return hasil
}

func main() {
	daftar := [][][2]int{
		{{15_000, 2}, {5_000, 3}, {8_000, 1}},
		{{60_000, 2}},
		{{70_000, 2}, {80_000, 2}},
		{{150_000, 4}},
		{},
		{{50_000, 2}, {10_000, 0}, {30_000, -1}},
	}

	for nomor, keranjang := range daftar {
		h := hitung(keranjang)

		fmt.Printf("Keranjang %d\n", nomor+1)
		fmt.Printf("  Subtotal : Rp%s\n", rupiah(h.subtotal))
		fmt.Printf("  Diskon   : Rp%s\n", rupiah(h.diskon))
		fmt.Printf("  Pajak    : Rp%s\n", rupiah(h.pajak))
		fmt.Printf("  Total    : Rp%s\n", rupiah(h.total))
		fmt.Printf("  Ditolak  : %d baris\n", h.ditolak)
	}

	fmt.Println()
	fmt.Println("Jalankan go test ./latihan untuk memeriksa.")
}
