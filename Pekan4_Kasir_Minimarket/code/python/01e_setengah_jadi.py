"""
CONTOH SETENGAH JADI - temukan dan perbaiki
Jalankan: python3 01e_setengah_jadi.py

Tiga fungsi di bawah hampir benar. Masing-masing punya satu kesalahan
yang membuat hasilnya salah pada kasus tertentu.

Tugas: perbaiki setiap fungsi tanpa mengubah nama, parameter, atau tipe
kembaliannya. Jalankan ulang berkas ini sampai hasil dan seharusnya sama.

Berkas ini bukan penilaian. Latihan ini melatih membaca kode orang lain,
keterampilan yang dipakai setiap hari di tim.
"""


def diskon_voucher(subtotal: int) -> int:
    """Voucher 10% bila subtotal > 100.000, dibulatkan setengah ke atas."""
    if subtotal >= 100_000:
        return int(0.10 * subtotal + 0.5)
    return 0


def total_belanja(keranjang: list[tuple[int, int]]) -> int:
    """Jumlahkan harga x jumlah untuk seluruh keranjang."""
    total: int = 0
    for harga, jumlah in keranjang:
        total = harga * jumlah
    return total


def bulatkan(nilai: float) -> int:
    """Bulatkan ke rupiah terdekat, setengah ke atas."""
    return round(nilai)


def main() -> None:
    print("1. diskon_voucher(100_000)")
    print(f"   hasil     : {diskon_voucher(100_000):,}".replace(",", "."))
    print("   seharusnya: 0")
    print()

    print("2. total_belanja([(15_000, 2), (5_000, 3)])")
    print(f"   hasil     : {total_belanja([(15_000, 2), (5_000, 3)]):,}".replace(",", "."))
    print("   seharusnya: 45.000")
    print()

    print("3. bulatkan(10_000.5)")
    print(f"   hasil     : {bulatkan(10_000.5):,}".replace(",", "."))
    print("   seharusnya: 10.001")
    print()

    print("Petunjuk: ketiga kesalahan ada di daftar A.9 dokumen bahan praktik.")


if __name__ == "__main__":
    main()
