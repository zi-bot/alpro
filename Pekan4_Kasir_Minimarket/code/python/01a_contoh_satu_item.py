"""
CONTOH 1 - Satu barang, satu keputusan (if / else)
Jalankan: python3 01a_contoh_satu_item.py

Aturan: diskon 10% diberikan bila subtotal lebih dari Rp100.000.

Perhatikan anotasi tipe pada setiap variabel:
    nama: tipe = nilai
Anotasi tidak mengubah perilaku program. Fungsinya memberi tahu pembaca
dan alat pemeriksa (mypy, pyright, editor) tipe yang diharapkan.
"""

harga: int = 120_000
jumlah: int = 1
subtotal: int = harga * jumlah
diskon: float = 0.0

if subtotal > 100_000:
    diskon = 0.10 * subtotal
else:
    diskon = 0.0

print(f"Subtotal : Rp{subtotal:,.0f}".replace(",", "."))
print(f"Diskon   : Rp{diskon:,.0f}".replace(",", "."))
