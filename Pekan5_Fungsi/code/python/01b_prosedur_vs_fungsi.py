# 01b — Prosedur dan fungsi: bedanya di nilai balik.
# Jalankan: python3 01b_prosedur_vs_fungsi.py
#
# cetak_garis  mengerjakan sesuatu lalu selesai. Tidak ada nilai yang diserahkan.
# hitung_subtotal menghitung sesuatu lalu menyerahkan hasilnya lewat return.
# Keduanya sah. Yang salah adalah mencampur: menghitung sambil mencetak.


def cetak_garis(panjang: int = 32) -> None:
    """Prosedur: menulis garis pemisah. Tidak mengembalikan nilai."""
    print("-" * panjang)


def hitung_subtotal(harga: int, jumlah: int) -> int:
    """Fungsi: mengembalikan subtotal, tidak mencetak apa pun."""
    return harga * jumlah


def hitung_subtotal_campur(harga: int, jumlah: int) -> int:
    """Kesalahan yang sering terjadi: mencetak di dalam, bukan mengembalikan."""
    print(harga * jumlah)
    # tidak ada return — fungsi ini mengembalikan None


def main() -> None:
    cetak_garis()
    print("PROSEDUR — hasilnya tidak bisa dipakai menghitung")
    cetak_garis()

    cetak_garis()
    print("FUNGSI — hasilnya bisa dipakai menghitung")
    cetak_garis()

    benar = hitung_subtotal(15_000, 2)
    print("Memanggil hitung_subtotal_campur — perhatikan angka yang muncul")
    print("tanpa pernah diminta untuk dicetak:")
    campur = hitung_subtotal_campur(15_000, 2)

    print()
    print("benar  :", benar, type(benar).__name__)
    print("campur :", campur, type(campur).__name__)
    print()
    print("benar + 5.000  :", benar + 5_000)
    try:
        print(campur + 5_000)
    except TypeError as galat:
        print("campur + 5.000 :", "TypeError:", galat)


main()
