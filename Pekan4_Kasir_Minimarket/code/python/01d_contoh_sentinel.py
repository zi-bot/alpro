"""
CONTOH TAMBAHAN - Perulangan while dengan penanda berhenti
Jalankan: python3 01d_contoh_sentinel.py

Kasir membaca pasangan (harga, jumlah) sampai menemukan harga negatif.
Data disiapkan di dalam daftar supaya bisa dijalankan tanpa mengetik.
"""

masukan: list[int] = [15_000, 2, 5_000, 3, 8_000, 1, -1]
#                     harga jml  harga jml  harga jml  sentinel

subtotal: int = 0
i: int = 0

while masukan[i] >= 0:
    harga: int = masukan[i]
    jumlah: int = masukan[i + 1]
    subtotal = subtotal + harga * jumlah
    i = i + 2

print(f"Subtotal : Rp{subtotal:,.0f}".replace(",", "."))
