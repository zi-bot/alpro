"""
Uji tabel untuk tugas4_template.hitung().

Jalankan: python3 -m unittest test_tugas4
"""

import unittest

import data_uji
import tugas4_template as kasir


class UjiTugas4(unittest.TestCase):
    def test_tabel_kasus(self) -> None:
        for (nama, keranjang), harapan in zip(data_uji.KASUS, data_uji.HARAPAN):
            with self.subTest(nama):
                self.assertEqual(kasir.hitung(keranjang), harapan)


if __name__ == "__main__":
    unittest.main()
