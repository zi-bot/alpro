// 01e — Tiga versi setengah jadi: gejala, penyebab, perbaikan (versi Go).
// Jalankan: go run 01e_setengah_jadi.go

package main

import (
	"fmt"
	"strconv"
)

var batas = []int{99999, 100000, 249999, 250000, 499999, 500000}

func seharusnya(subtotal int) int {
	if subtotal < 100000 {
		return 0
	} else if subtotal < 250000 {
		return subtotal * 5 / 100
	} else if subtotal < 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

func v1MencetakBukanMengembalikan(subtotal int) {
	fmt.Printf("  (di dalam fungsi, angka ini tercetak: %d)\n", subtotal*5/100)
}

func v2BatasIkutTergeser(subtotal int) int {
	if subtotal <= 100000 {
		return 0
	} else if subtotal <= 250000 {
		return subtotal * 5 / 100
	} else if subtotal <= 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

func v3UrutanTerbalik(subtotal int) int {
	if subtotal >= 100000 {
		return subtotal * 5 / 100
	} else if subtotal >= 250000 {
		return subtotal * 10 / 100
	} else if subtotal >= 500000 {
		return subtotal * 15 / 100
	}
	return 0
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
	fmt.Println("SALAH 1 — mencetak di dalam fungsi, bukan mengembalikan")
	fmt.Println("----------------------------------------------------------")
	v1MencetakBukanMengembalikan(100000)
	fmt.Println("  nilai yang diterima pemanggil: (tidak ada)")
	fmt.Println("  akibatnya: tidak bisa dipakai menghitung apa pun.")
	fmt.Println()

	type pengujian struct {
		judul  string
		fungsi func(int) int
	}

	daftar := []pengujian{
		{"SALAH 2 — batas memakai <=", v2BatasIkutTergeser},
		{"SALAH 3 — urutan kondisi terbalik", v3UrutanTerbalik},
	}

	for _, p := range daftar {
		fmt.Println(p.judul)
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("  %12s  %11s  %11s  STATUS\n", "SUBTOTAL", "SEHARUSNYA", "HASIL")
		for _, s := range batas {
			benar := seharusnya(s)
			hasil := p.fungsi(s)
			tanda := "ok"
			if hasil != benar {
				tanda = "SALAH"
			}
			fmt.Printf("  %12s  %11s  %11s  %s\n",
				rupiah(s),
				rupiah(benar),
				rupiah(hasil),
				tanda,
			)
		}
		fmt.Println()
	}

	fmt.Println("Yang harus terlihat:")
	fmt.Println("  SALAH 2  Rp100.000, Rp250.000, dan Rp500.000 salah: tepat di batas,")
	fmt.Println("           nilainya masih dihitung dengan tingkat di bawahnya.")
	fmt.Println("  SALAH 3  benar hanya untuk tingkat 5% (Rp99.999, Rp100.000, Rp249.999).")
	fmt.Println("           Mulai Rp250.000 hasilnya selalu 5%, karena cabang pertama")
	fmt.Println("           sudah menang sebelum cabang berikutnya diperiksa.")
}
