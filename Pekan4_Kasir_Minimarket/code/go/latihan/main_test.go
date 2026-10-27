package main

import "testing"

// Uji tabel: setiap baris berisi nama, keranjang, dan hasil yang diharapkan.
func TestTabelKasus(t *testing.T) {
	kasus := []struct {
		nama      string
		keranjang [][2]int
		harapan   [5]int
	}{
		{
			"di bawah Rp100.000",
			[][2]int{{15_000, 2}, {5_000, 3}, {8_000, 1}},
			[5]int{53_000, 0, 5_830, 58_830, 0},
		},
		{
			"batas bawah diskon 5%",
			[][2]int{{60_000, 2}},
			[5]int{120_000, 6_000, 12_540, 126_540, 0},
		},
		{
			"diskon 10%",
			[][2]int{{70_000, 2}, {80_000, 2}},
			[5]int{300_000, 30_000, 29_700, 299_700, 0},
		},
		{
			"diskon 15%",
			[][2]int{{150_000, 4}},
			[5]int{600_000, 90_000, 56_100, 566_100, 0},
		},
		{
			"keranjang kosong",
			[][2]int{},
			[5]int{0, 0, 0, 0, 0},
		},
		{
			"dua baris ditolak",
			[][2]int{{50_000, 2}, {10_000, 0}, {30_000, -1}},
			[5]int{100_000, 5_000, 10_450, 105_450, 2},
		},
	}

	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			h := hitung(k.keranjang)
			dapat := [5]int{h.subtotal, h.diskon, h.pajak, h.total, h.ditolak}

			if dapat != k.harapan {
				t.Errorf("dapat %v, harapan %v", dapat, k.harapan)
			}
		})
	}
}
