# 01c — Diskon bertingkat: satu fungsi, empat tingkat.
# Jalankan: python3 01c_diskon_bertingkat.py
#
# Aturan (batas bawah inklusif, batas atas milik tingkat berikutnya):
#   < Rp100.000                 -> 0%
#   Rp100.000 - Rp249.999       -> 5%
#   Rp250.000 - Rp499.999       -> 10%
#   >= Rp500.000                -> 15%
#
# Perhatikan urutan pemeriksaan: dari batas terkecil ke terbesar, dan setiap
# cabang langsung return. Tidak ada elif yang perlu dibaca sampai habis.


def hitung_diskon(subtotal: int) -> int:
    """Diskon dalam rupiah untuk satu subtotal. Batas bawah inklusif."""
    if subtotal < 0:
        raise ValueError("subtotal tidak boleh negatif")
    if subtotal < 100_000:
        return 0
    if subtotal < 250_000:
        return subtotal * 5 // 100
    if subtotal < 500_000:
        return subtotal * 10 // 100
    return subtotal * 15 // 100


def rupiah(nilai: int) -> str:
    return "Rp" + f"{nilai:,}".replace(",", ".")


def persen(subtotal: int) -> str:
    """Tingkat diskon sebagai teks, untuk dicetak di tabel."""
    if subtotal < 100_000:
        return "0%"
    if subtotal < 250_000:
        return "5%"
    if subtotal < 500_000:
        return "10%"
    return "15%"


def main() -> None:
    # Nilai di sekitar setiap batas. Batas sendiri diuji, begitu juga satu
    # rupiah di bawahnya — di situlah kesalahan paling sering bersembunyi.
    uji = [0, 53_000, 99_999, 100_000, 249_999, 250_000, 499_999, 500_000, 1_250_000]

    print("SUBTOTAL      TINGKAT   DISKON")
    print("-" * 34)
    for subtotal in uji:
        print(f"{rupiah(subtotal):>13}  {persen(subtotal):>6}  {rupiah(hitung_diskon(subtotal)):>10}")

    print()
    print("Dua batas yang paling mudah salah:")
    print("   99.999 ->", rupiah(hitung_diskon(99_999)), " (di bawah batas, tetap 0%)")
    print("  100.000 ->", rupiah(hitung_diskon(100_000)), " (tepat di batas, sudah 5%)")
    print()
    print("Pecahan rupiah dibuang ke bawah, bukan dibulatkan:")
    print("  249.999 x 5% = 12.499,95 ->", rupiah(hitung_diskon(249_999)))


main()
