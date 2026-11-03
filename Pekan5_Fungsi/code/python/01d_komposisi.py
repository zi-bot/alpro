# 01d — Kasus utuh: lima fungsi, satu alur utama.
# Jalankan: python3 01d_komposisi.py
#
# Pembagian tanggung jawab:
#   hitung_subtotal(keranjang)      -> int   menjumlahkan isi keranjang
#   hitung_diskon(subtotal)         -> int   memutuskan tingkat diskon
#   hitung_ppn(dasar)               -> int   pajak 11% dari nilai setelah diskon
#   hitung_total(subtotal, diskon)  -> int   memanggil hitung_ppn, lalu menjumlah
#   tampilkan_hasil(...)            -> None  prosedur: hanya mencetak
#
# Alur utama tidak memuat satu pun operasi aritmatika.


def hitung_subtotal(keranjang: list[tuple[int, int]]) -> int:
    """Jumlahkan harga x jumlah untuk seluruh baris keranjang."""
    subtotal = 0
    for harga, jumlah in keranjang:
        subtotal = subtotal + harga * jumlah
    return subtotal


def hitung_diskon(subtotal: int) -> int:
    """Diskon dalam rupiah. Batas bawah inklusif."""
    if subtotal < 0:
        raise ValueError("subtotal tidak boleh negatif")
    if subtotal < 100_000:
        return 0
    if subtotal < 250_000:
        return subtotal * 5 // 100
    if subtotal < 500_000:
        return subtotal * 10 // 100
    return subtotal * 15 // 100


def hitung_ppn(dasar: int) -> int:
    """PPN 11% dari nilai yang sudah dikurangi diskon."""
    return dasar * 11 // 100


def hitung_total(subtotal: int, diskon: int) -> int:
    """Total bayar: nilai setelah diskon ditambah pajak."""
    dasar = subtotal - diskon
    return dasar + hitung_ppn(dasar)


def tampilkan_hasil(nama: str, keranjang: list[tuple[int, int]]) -> None:
    """Prosedur: menghitung lewat fungsi lain, lalu mencetak. Tidak mengembalikan nilai."""
    subtotal = hitung_subtotal(keranjang)
    diskon = hitung_diskon(subtotal)
    total = hitung_total(subtotal, diskon)

    print(f"NOTA {nama}")
    print("-" * 34)
    for harga, jumlah in keranjang:
        print(f"  {rupiah(harga):>11} x {jumlah} = {rupiah(harga * jumlah):>11}")
    print("-" * 34)
    print(f"  {'Subtotal':<20}{rupiah(subtotal):>14}")
    print(f"  {'Diskon':<20}{rupiah(diskon):>14}")
    print(f"  {'PPN 11%':<20}{rupiah(hitung_ppn(subtotal - diskon)):>14}")
    print(f"  {'TOTAL':<20}{rupiah(total):>14}")
    print()


def rupiah(nilai: int) -> str:
    return "Rp" + f"{nilai:,}".replace(",", ".")


def main() -> None:
    tampilkan_hasil("KERANJANG A", [(15_000, 2), (5_000, 3), (8_000, 1)])
    tampilkan_hasil("KERANJANG B", [(125_000, 2), (250_000, 1)])


main()
