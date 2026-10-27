"""
PRAKTIK TERBIMBING - Kasir minimarket tahan salah input

Jalankan program : python3 latihan_template.py
Jalankan uji     : python3 -m unittest test_latihan

Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
Selama TODO belum diisi, program tetap berjalan tetapi uji akan gagal.

Ikuti anotasi tipe yang sudah disiapkan:
    int   untuk rupiah bulat
    float untuk hasil persentase sebelum dibulatkan

Aturan:
    - Baris dengan jumlah <= 0 atau harga < 0 DILEWATI dan dihitung ditolak.
    - subtotal < 100.000              -> diskon 0%
    - 100.000 <= subtotal < 250.000   -> diskon 5%
    - 250.000 <= subtotal < 500.000   -> diskon 10%
    - subtotal >= 500.000             -> diskon 15%
    - pajak 11% dihitung dari (subtotal - diskon)
    - diskon dan pajak dibulatkan ke rupiah terdekat, setengah ke atas
    - keranjang tanpa baris valid -> subtotal 0, pajak 0, total 0
"""

DAFTAR_KERANJANG: list[list[tuple[int, int]]] = [
    [(15_000, 2), (5_000, 3), (8_000, 1)],          # 53.000
    [(60_000, 2)],                                  # 120.000
    [(70_000, 2), (80_000, 2)],                     # 300.000
    [(150_000, 4)],                                 # 600.000
    [],                                             # kosong
    [(50_000, 2), (10_000, 0), (30_000, -1)],       # 2 baris ditolak
]


def hitung(keranjang: list[tuple[int, int]]) -> dict[str, int]:
    """Hitung satu keranjang. Lengkapi 4 TODO di bawah."""
    subtotal: int = 0
    ditolak: int = 0

    for harga, jumlah in keranjang:
        # ----------------------------------------------------
        # TODO 1 - Validasi baris
        # Tulis if di bawah ini: bila jumlah <= 0 atau harga < 0,
        # tambahkan 1 ke ditolak, lalu continue.
        # ----------------------------------------------------

        subtotal = subtotal + harga * jumlah

    # --------------------------------------------------------
    # TODO 2 - Diskon bertingkat
    # Tulis if / elif / else di bawah ini yang mengisi diskon.
    # Ada empat kemungkinan: 0%, 5%, 10%, dan 15%.
    # --------------------------------------------------------
    diskon: float = 0.0

    # --------------------------------------------------------
    # TODO 3 - Bulatkan diskon
    # Isi diskon_bulat dengan int(diskon + 0.5).
    # Jangan memakai round().
    # --------------------------------------------------------
    diskon_bulat: int = 0

    # --------------------------------------------------------
    # TODO 4 - Pajak dan total
    # Pajak dihitung dari (subtotal - diskon_bulat), lalu dibulatkan
    # dengan cara yang sama seperti TODO 3.
    # --------------------------------------------------------
    pajak: int = 0
    total: int = 0

    return {
        "subtotal": subtotal,
        "diskon": diskon_bulat,
        "pajak": pajak,
        "total": total,
        "ditolak": ditolak,
    }


def main() -> None:
    for nomor, keranjang in enumerate(DAFTAR_KERANJANG, start=1):
        h = hitung(keranjang)
        print(f"Keranjang {nomor}")
        print(f"  Subtotal : Rp{h['subtotal']:,}".replace(",", "."))
        print(f"  Diskon   : Rp{h['diskon']:,}".replace(",", "."))
        print(f"  Pajak    : Rp{h['pajak']:,}".replace(",", "."))
        print(f"  Total    : Rp{h['total']:,}".replace(",", "."))
        print(f"  Ditolak  : {h['ditolak']} baris")

    print()
    print("Jalankan python3 -m unittest test_latihan untuk memeriksa.")


if __name__ == "__main__":
    main()
