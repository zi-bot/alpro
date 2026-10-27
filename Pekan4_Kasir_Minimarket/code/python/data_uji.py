"""
DATA UJI - Pekan 4 / Tugas 4
Jangan diubah. Dipakai oleh tugas4_template.py, test_tugas4.py, dan solusinya.

Setiap kasus adalah pasangan (nama, keranjang).
Setiap keranjang berisi pasangan (harga, jumlah).

Anotasi tipe di bawah menjelaskan bentuk data tanpa mengubah isinya:
    list[tuple[str, list[tuple[int, int]]]]
    artinya daftar, tiap isinya pasangan (teks, daftar pasangan bilangan bulat)
"""

KASUS: list[tuple[str, list[tuple[int, int]]]] = [
    ("di bawah batas voucher", [(25_000, 2), (20_000, 1)]),

    ("tepat Rp100.000", [(100_000, 1)]),

    ("Rp100.001", [(100_001, 1)]),

    ("pembulatan setengah ke atas", [(100_005, 1)]),

    ("tepat Rp149.999", [(149_999, 1)]),

    ("tepat Rp150.000", [(150_000, 1)]),

    ("di atas Rp150.000", [(50_000, 4)]),

    ("dua baris ditolak", [(50_000, 2), (10_000, 0), (30_000, -1)]),

    ("keranjang kosong", []),

    ("semua baris ditolak", [(-5_000, 2)]),
]

# Harapan hasil untuk tiap kasus di atas.
# Kunci harus sama dengan kunci yang dihasilkan program Anda.
HARAPAN: list[dict[str, int]] = [
    {"subtotal": 70_000, "voucher": 0, "ongkir": 15_000, "total": 85_000, "ditolak": 0},

    {"subtotal": 100_000, "voucher": 0, "ongkir": 15_000, "total": 115_000, "ditolak": 0},

    {"subtotal": 100_001, "voucher": 10_000, "ongkir": 15_000, "total": 105_001, "ditolak": 0},

    {"subtotal": 100_005, "voucher": 10_001, "ongkir": 15_000, "total": 105_004, "ditolak": 0},

    {"subtotal": 149_999, "voucher": 15_000, "ongkir": 15_000, "total": 149_999, "ditolak": 0},

    {"subtotal": 150_000, "voucher": 15_000, "ongkir": 0, "total": 135_000, "ditolak": 0},

    {"subtotal": 200_000, "voucher": 20_000, "ongkir": 0, "total": 180_000, "ditolak": 0},

    {"subtotal": 100_000, "voucher": 0, "ongkir": 15_000, "total": 115_000, "ditolak": 2},

    {"subtotal": 0, "voucher": 0, "ongkir": 0, "total": 0, "ditolak": 0},

    {"subtotal": 0, "voucher": 0, "ongkir": 0, "total": 0, "ditolak": 1},
]
