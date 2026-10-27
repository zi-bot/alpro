// Tugas 4 - checkout dan aturan voucher (versi Go)
// Pembanding untuk tugas4_template.py. Tidak wajib dikerjakan.
//
// Jalankan program: go run ./tugas4
// Jalankan uji    : go test ./tugas4
//
// Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
// Kunci ada di solusi/tugas4/main.go.

package main

import "fmt"

type Kasus struct {
	nama      string
	keranjang [][2]int
}

type Hasil struct {
	subtotal int
	voucher  int
	ongkir   int
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

	// TODO 2 - Voucher
	// Ganti 0 dengan voucher 10% bila subtotal > 100.000.
	// Pembulatan setengah ke atas: (subtotal*10 + 50) / 100
	voucher := 0

	// TODO 3 - Ongkir
	// Ganti 0 dengan 15.000, kecuali gratis bila subtotal >= 150.000
	// atau keranjang tidak berisi baris valid.
	ongkir := 0

	// TODO 4 - Total
	// Ganti 0 dengan subtotal - voucher + ongkir.
	total := 0

	return Hasil{
		subtotal: subtotal,
		voucher:  voucher,
		ongkir:   ongkir,
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
	kasus := []Kasus{
		{"di bawah batas voucher", [][2]int{{25_000, 2}, {20_000, 1}}},
		{"tepat Rp100.000", [][2]int{{100_000, 1}}},
		{"Rp100.001", [][2]int{{100_001, 1}}},
		{"pembulatan setengah ke atas", [][2]int{{100_005, 1}}},
		{"tepat Rp149.999", [][2]int{{149_999, 1}}},
		{"tepat Rp150.000", [][2]int{{150_000, 1}}},
		{"di atas Rp150.000", [][2]int{{50_000, 4}}},
		{"dua baris ditolak", [][2]int{{50_000, 2}, {10_000, 0}, {30_000, -1}}},
		{"keranjang kosong", [][2]int{}},
		{"semua baris ditolak", [][2]int{{-5_000, 2}}},
	}

	for nomor, k := range kasus {
		h := hitung(k.keranjang)

		fmt.Printf(
			"%d. %-30s subtotal %9s | voucher %8s | ongkir %6s | total %9s | ditolak %d\n",
			nomor+1,
			k.nama,
			rupiah(h.subtotal),
			rupiah(h.voucher),
			rupiah(h.ongkir),
			rupiah(h.total),
			h.ditolak,
		)
	}

	fmt.Println()
	fmt.Println("Jalankan go test ./tugas4 untuk memeriksa.")
}
