// test_diskon.go — uji otomatis untuk hitungDiskon dan hitungTotal (versi Go).
// Jalankan: go run test_diskon.go

package main

import (
	"errors"
	"fmt"
)

func validasiSubtotal(subtotal int) (int, error) {
	if subtotal < 0 {
		return 0, errors.New("subtotal tidak boleh negatif")
	}
	return subtotal, nil
}

func hitungDiskon(subtotal int) int {
	if subtotal < 100000 {
		return 0
	} else if subtotal < 250000 {
		return subtotal * 5 / 100
	} else if subtotal < 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

func hitungTotal(subtotal int, diskon int) int {
	dasar := subtotal - diskon
	return dasar + dasar*11/100
}

func testBawahBatas() {
	if hitungDiskon(0) != 0 || hitungDiskon(99999) != 0 {
		panic("testBawahBatas gagal")
	}
}

func testTepatDiBatas() {
	if hitungDiskon(100000) != 5000 || hitungDiskon(250000) != 25000 || hitungDiskon(500000) != 75000 {
		panic("testTepatDiBatas gagal")
	}
}

func testSatuRupiahDiAtasBatas() {
	if hitungDiskon(100001) != 5000 || hitungDiskon(250001) != 25000 || hitungDiskon(500001) != 75000 {
		panic("testSatuRupiahDiAtasBatas gagal")
	}
}

func testTepiAtasSetiapTingkat() {
	if hitungDiskon(249999) != 12499 || hitungDiskon(499999) != 49999 {
		panic("testTepiAtasSetiapTingkat gagal")
	}
}

func testPecahanDibuangKeBawah() {
	if hitungDiskon(249999) != 12499 {
		panic("testPecahanDibuangKeBawah gagal")
	}
}

func testDiskonTidakPernahMelebihiSepertiga() {
	for s := 0; s <= 1000000; s += 25000 {
		if hitungDiskon(s) > s/3 {
			panic("testDiskonTidakPernahMelebihiSepertiga gagal")
		}
	}
}

func testSubtotalNegatifDitolak() {
	_, err := validasiSubtotal(-1)
	if err == nil {
		panic("testSubtotalNegatifDitolak gagal: harusnya error")
	}
}

func testPpnSetelahDiskon() {
	if hitungTotal(100000, 5000) != 105450 {
		panic("testPpnSetelahDiskon gagal")
	}
}

func testTanpaDiskon() {
	if hitungTotal(50000, 0) != 55500 {
		panic("testTanpaDiskon gagal")
	}
}

func testTotalBisaTurunSaatSubtotalNaik() {
	total250k := hitungTotal(250000, hitungDiskon(250000))
	total249k := hitungTotal(249999, hitungDiskon(249999))
	if total250k >= total249k {
		panic("testTotalBisaTurunSaatSubtotalNaik gagal: total250k harus lebih kecil dari total249k")
	}
}

func main() {
	testBawahBatas()
	testTepatDiBatas()
	testSatuRupiahDiAtasBatas()
	testTepiAtasSetiapTingkat()
	testPecahanDibuangKeBawah()
	testDiskonTidakPernahMelebihiSepertiga()
	testSubtotalNegatifDitolak()
	testPpnSetelahDiskon()
	testTanpaDiskon()
	testTotalBisaTurunSaatSubtotalNaik()
	fmt.Println("Ran 10 tests ... OK")
}
