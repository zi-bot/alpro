// 01b — Prosedur dan fungsi: bedanya di nilai balik (versi Go).
// Jalankan: go run 01b_prosedur_vs_fungsi.go
//
// cetakGaris adalah prosedur: tidak ada tipe nilai balik di belakang tanda kurung.
// hitungSubtotal adalah fungsi murni: mengembalikan int via return.
// hitungSubtotalCampur melakukan efek samping (cetak) tanpa mengembalikan nilai.

package main

import (
	"fmt"
	"strings"
)

func cetakGaris(panjang int) {
	fmt.Println(strings.Repeat("-", panjang))
}

func hitungSubtotal(harga int, jumlah int) int {
	return harga * jumlah
}

// Prosedur yang keliru: menghitung sambil mencetak di dalam fungsi
func hitungSubtotalCampur(harga int, jumlah int) {
	fmt.Println(harga * jumlah)
	// Tidak ada return nilai
}

func main() {
	cetakGaris(32)
	fmt.Println("PROSEDUR — hasilnya tidak bisa dipakai menghitung")
	cetakGaris(32)

	cetakGaris(32)
	fmt.Println("FUNGSI — hasilnya bisa dipakai menghitung")
	cetakGaris(32)

	benar := hitungSubtotal(15000, 2)
	fmt.Println("Memanggil hitungSubtotalCampur — perhatikan angka yang muncul")
	fmt.Println("tanpa pernah diminta untuk dicetak:")
	hitungSubtotalCampur(15000, 2)

	fmt.Println()
	fmt.Printf("benar  : %d int\n", benar)
	fmt.Println("campur : (tanpa nilai balik)")
	fmt.Println()
	fmt.Println("benar + 5.000  :", benar+5000)
	fmt.Println("campur + 5.000 : [Kompilasi Ditolak] hitungSubtotalCampur(15000, 2) used as value")
	fmt.Println("                 Sistem tipe Go mencegah pemanggilan prosedur di dalam ekspresi aritmatika!")
}
