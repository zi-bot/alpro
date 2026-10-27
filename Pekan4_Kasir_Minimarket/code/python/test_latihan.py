"""
Uji tabel untuk latihan_template.hitung().

Jalankan: python3 -m unittest test_latihan
"""

import unittest

import latihan_template as kasir

KASUS: list[tuple[str, list[tuple[int, int]], dict[str, int]]] = [
    (
        "di bawah Rp100.000",
        [(15_000, 2), (5_000, 3), (8_000, 1)],
        {"subtotal": 53_000, "diskon": 0, "pajak": 5_830, "total": 58_830, "ditolak": 0},
    ),
    (
        "batas bawah diskon 5%",
        [(60_000, 2)],
        {"subtotal": 120_000, "diskon": 6_000, "pajak": 12_540, "total": 126_540, "ditolak": 0},
    ),
    (
        "diskon 10%",
        [(70_000, 2), (80_000, 2)],
        {"subtotal": 300_000, "diskon": 30_000, "pajak": 29_700, "total": 299_700, "ditolak": 0},
    ),
    (
        "diskon 15%",
        [(150_000, 4)],
        {"subtotal": 600_000, "diskon": 90_000, "pajak": 56_100, "total": 566_100, "ditolak": 0},
    ),
    (
        "keranjang kosong",
        [],
        {"subtotal": 0, "diskon": 0, "pajak": 0, "total": 0, "ditolak": 0},
    ),
    (
        "dua baris ditolak",
        [(50_000, 2), (10_000, 0), (30_000, -1)],
        {"subtotal": 100_000, "diskon": 5_000, "pajak": 10_450, "total": 105_450, "ditolak": 2},
    ),
]


class UjiLatihan(unittest.TestCase):
    def test_tabel_kasus(self) -> None:
        for nama, keranjang, harapan in KASUS:
            with self.subTest(nama):
                self.assertEqual(kasir.hitung(keranjang), harapan)


if __name__ == "__main__":
    unittest.main()
