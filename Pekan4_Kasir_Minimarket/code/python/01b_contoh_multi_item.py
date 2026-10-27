"""
CONTOH 2 - Banyak barang, perulangan + penjumlahan
Jalankan: python3 01b_contoh_multi_item.py

Subtotal adalah jumlah dari harga dikali jumlah untuk setiap baris.
Tipe keranjang: list[tuple[int, int]]
    daftar, tiap isinya pasangan (harga, jumlah) bilangan bulat.
"""

keranjang: list[tuple[int, int]] = [(15_000, 2), (5_000, 3), (8_000, 1)]

subtotal: int = 0
for harga, jumlah in keranjang:
    subtotal = subtotal + harga * jumlah

print(f"Subtotal : Rp{subtotal:,.0f}".replace(",", "."))
