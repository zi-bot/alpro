# test_diskon.py — uji otomatis untuk hitung_diskon dan hitung_total.
# Jalankan dari folder python/:  python3 -m unittest test_diskon -v
#
# Yang diuji bukan hanya nilai tengah, tetapi setiap batas dan satu rupiah
# di sekitarnya. Di situlah kesalahan off-by-one bersembunyi (lihat 01e).

import unittest

from solusi.tugas5_solusi import hitung_diskon, hitung_total, validasi_subtotal


class UjiDiskon(unittest.TestCase):
    def test_bawah_batas(self):
        self.assertEqual(hitung_diskon(0), 0)
        self.assertEqual(hitung_diskon(99_999), 0)

    def test_tepat_di_batas(self):
        self.assertEqual(hitung_diskon(100_000), 5_000)
        self.assertEqual(hitung_diskon(250_000), 25_000)
        self.assertEqual(hitung_diskon(500_000), 75_000)

    def test_satu_rupiah_di_atas_batas(self):
        self.assertEqual(hitung_diskon(100_001), 5_000)
        self.assertEqual(hitung_diskon(250_001), 25_000)
        self.assertEqual(hitung_diskon(500_001), 75_000)

    def test_tepi_atas_setiap_tingkat(self):
        self.assertEqual(hitung_diskon(249_999), 12_499)
        self.assertEqual(hitung_diskon(499_999), 49_999)

    def test_pecahan_dibuang_ke_bawah(self):
        # 249.999 x 5% = 12.499,95 -> 12.499, bukan 12.500
        self.assertEqual(hitung_diskon(249_999), 12_499)

    def test_diskon_tidak_pernah_melebihi_sepertiga(self):
        for subtotal in range(0, 1_000_001, 25_000):
            self.assertLessEqual(hitung_diskon(subtotal), subtotal // 3)

    def test_subtotal_negatif_ditolak(self):
        with self.assertRaises(ValueError):
            validasi_subtotal(-1)


class UjiTotal(unittest.TestCase):
    def test_ppn_setelah_diskon(self):
        # 100.000 - 5.000 = 95.000; PPN 11% = 10.450; total 105.450
        self.assertEqual(hitung_total(100_000, 5_000), 105_450)

    def test_tanpa_diskon(self):
        # 50.000 + 11% = 55.500
        self.assertEqual(hitung_total(50_000, 0), 55_500)

    def test_total_bisa_turun_saat_subtotal_naik(self):
        # Melewati batas Rp250.000 membuat total justru lebih kecil.
        self.assertLess(hitung_total(250_000, hitung_diskon(250_000)),
                        hitung_total(249_999, hitung_diskon(249_999)))


if __name__ == "__main__":
    unittest.main()
