package main

import "testing"

// Uji tabel: setiap baris berisi nama kasus, keranjang, dan hasil yang diharapkan.
func TestTabelKasus(t *testing.T) {
	kasus := []struct {
		nama      string
		keranjang [][2]int
		harapan   [5]int
	}{
		{"di bawah batas voucher", [][2]int{{25_000, 2}, {20_000, 1}}, [5]int{70_000, 0, 15_000, 85_000, 0}},
		{"tepat Rp100.000", [][2]int{{100_000, 1}}, [5]int{100_000, 0, 15_000, 115_000, 0}},
		{"Rp100.001", [][2]int{{100_001, 1}}, [5]int{100_001, 10_000, 15_000, 105_001, 0}},
		{"pembulatan setengah ke atas", [][2]int{{100_005, 1}}, [5]int{100_005, 10_001, 15_000, 105_004, 0}},
		{"tepat Rp149.999", [][2]int{{149_999, 1}}, [5]int{149_999, 15_000, 15_000, 149_999, 0}},
		{"tepat Rp150.000", [][2]int{{150_000, 1}}, [5]int{150_000, 15_000, 0, 135_000, 0}},
		{"di atas Rp150.000", [][2]int{{50_000, 4}}, [5]int{200_000, 20_000, 0, 180_000, 0}},
		{"dua baris ditolak", [][2]int{{50_000, 2}, {10_000, 0}, {30_000, -1}}, [5]int{100_000, 0, 15_000, 115_000, 2}},
		{"keranjang kosong", [][2]int{}, [5]int{0, 0, 0, 0, 0}},
		{"semua baris ditolak", [][2]int{{-5_000, 2}}, [5]int{0, 0, 0, 0, 1}},
	}

	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			h := hitung(k.keranjang)
			dapat := [5]int{h.subtotal, h.voucher, h.ongkir, h.total, h.ditolak}

			if dapat != k.harapan {
				t.Errorf("dapat %v, harapan %v", dapat, k.harapan)
			}
		})
	}
}
