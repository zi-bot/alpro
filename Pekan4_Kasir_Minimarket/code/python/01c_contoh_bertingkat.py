"""
CONTOH 3 - Diskon bertingkat + pajak (if / elif / else)
Jalankan: python3 01c_contoh_bertingkat.py

Aturan:
    subtotal < 100.000              -> diskon 0%
    100.000 <= subtotal < 250.000   -> diskon 5%
    subtotal >= 250.000             -> diskon 10%
    pajak 11% dihitung dari (subtotal - diskon)
"""

keranjang: list[tuple[int, int]] = [(60_000, 2), (70_000, 2)]

subtotal: int = 0
for harga, jumlah in keranjang:
    subtotal = subtotal + harga * jumlah

diskon: float = 0.0
if subtotal < 100_000:
    diskon = 0.0
elif subtotal < 250_000:
    diskon = 0.05 * subtotal
else:
    diskon = 0.10 * subtotal

pajak: float = 0.11 * (subtotal - diskon)
total: float = (subtotal - diskon) + pajak

print(f"Subtotal : Rp{subtotal:,.0f}".replace(",", "."))
print(f"Diskon   : Rp{diskon:,.0f}".replace(",", "."))
print(f"Pajak    : Rp{pajak:,.0f}".replace(",", "."))
print(f"Total    : Rp{total:,.0f}".replace(",", "."))
