"""
TUGAS 4 - Perancangan checkout dan aturan voucher (kasir minimarket)
Pekan 4 | Sub-CPMK-4 | Rubrik A1 | Bobot 3%

Jalankan program : python3 tugas4_template.py
Jalankan uji     : python3 -m unittest test_tugas4

Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
Selama TODO belum diisi, program tetap berjalan tetapi uji akan gagal.
Jangan mengubah data_uji.py.

Perhatikan anotasi tipe pada setiap variabel. Anotasi tidak mengubah
perilaku program, tetapi menunjukkan tipe yang diharapkan. Semua nilai
rupiah di tugas ini adalah bilangan bulat (int).

Aturan bisnis:
    1. Baris valid bila harga >= 0 dan jumlah > 0.
    2. Baris tidak valid ditolak: tidak dihitung, dicatat jumlahnya.
    3. subtotal = jumlah (harga x jumlah) dari baris valid.
    4. Voucher 10% hanya bila subtotal > 100.000.
    5. Voucher dibulatkan ke rupiah terdekat, setengah ke atas.
    6. Ongkir Rp15.000, gratis bila subtotal >= 150.000 (sebelum diskon).
    7. Keranjang tanpa baris valid: total 0 dan tanpa ongkir.
    8. total = subtotal - voucher + ongkir
"""

import data_uji


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
    # TODO 2 - Voucher
    # Voucher 10% hanya berlaku bila subtotal > 100.000.
    # Bulatkan dengan int(0.10 * subtotal + 0.5).
    # Jangan memakai round().
    # --------------------------------------------------------
    voucher: int = 0

    # --------------------------------------------------------
    # TODO 3 - Ongkir
    # Ongkir Rp15.000, tetapi gratis bila subtotal >= 150.000.
    # Keranjang tanpa baris valid tidak dikenakan ongkir.
    # --------------------------------------------------------
    ongkir: int = 0

    # --------------------------------------------------------
    # TODO 4 - Total
    # total = subtotal - voucher + ongkir
    # --------------------------------------------------------
    total: int = 0

    return {
        "subtotal": subtotal,
        "voucher": voucher,
        "ongkir": ongkir,
        "total": total,
        "ditolak": ditolak,
    }


def main() -> None:
    for nomor, (nama, keranjang) in enumerate(data_uji.KASUS, start=1):
        h = hitung(keranjang)
        baris: str = (
            f"{nomor}. {nama:<30} subtotal {h['subtotal']:>9,} | "
            f"voucher {h['voucher']:>8,} | ongkir {h['ongkir']:>6,} | "
            f"total {h['total']:>9,} | ditolak {h['ditolak']}"
        )
        print(baris.replace(",", "."))

    print()
    print("Jalankan python3 -m unittest test_tugas4 untuk memeriksa.")


if __name__ == "__main__":
    main()
