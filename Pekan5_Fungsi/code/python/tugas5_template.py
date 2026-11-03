# tugas5_template.py — Tugas 5 (individu, rubrik A2, 5% nilai akhir).
# Jalankan: python3 tugas5_template.py
#
# Kerjakan empat TODO. Jangan mengubah nama fungsi, parameter, atau urutan
# argumen: dosen memanggil fungsi Anda dengan nama itu.
#
# Kontrak yang diminta tugas:
#   hitung_diskon(subtotal) mengembalikan NILAI DISKON (rupiah),
#   tidak membaca input, tidak mencetak apa pun.
#   Subtotal < 0 harus ditolak sebelum perhitungan.

# Kasus uji yang harus lolos (subtotal, total yang benar).
# Angka total dipakai dosen untuk memeriksa hasil Anda.
KASUS_UJI = [
    ("K-01", [(50_000, 1)]),
    ("K-02", [(99_999, 1)]),
    ("K-03", [(100_000, 1)]),
    ("K-04", [(100_001, 1)]),
    ("K-05", [(249_999, 1)]),
    ("K-06", [(250_000, 1)]),
    ("K-07", [(499_999, 1)]),
    ("K-08", [(500_000, 1)]),
    ("K-09", [(125_000, 2), (250_000, 1)]),
    ("K-10", [(0, 0)]),
]


def validasi_subtotal(subtotal: int) -> int:
    """TODO 1 — tolak subtotal negatif.

    Kembalikan subtotal apa adanya bila sah; naikkan ValueError dengan pesan
    'subtotal tidak boleh negatif' bila tidak sah.
    """
    raise NotImplementedError("TODO 1: validasi_subtotal belum diisi")


def hitung_diskon(subtotal: int) -> int:
    """TODO 2 — diskon dalam rupiah, batas bawah inklusif.

    < Rp100.000             -> 0%
    Rp100.000 - Rp249.999   -> 5%
    Rp250.000 - Rp499.999   -> 10%
    >= Rp500.000            -> 15%
    """
    raise NotImplementedError("TODO 2: hitung_diskon belum diisi")


def hitung_total(subtotal: int, diskon: int) -> int:
    """TODO 3 — total bayar: nilai setelah diskon ditambah PPN 11%."""
    raise NotImplementedError("TODO 3: hitung_total belum diisi")


def tampilkan_total(nama: str, subtotal: int, diskon: int, total: int) -> None:
    """TODO 4 — prosedur penyajian. Cetak empat baris, tidak mengembalikan nilai.

    Bentuk yang diminta:
        K-01  subtotal Rp50.000  diskon Rp0  total Rp55.500
    (satu spasi ganda sebagai pemisah antarbagian)
    """
    raise NotImplementedError("TODO 4: tampilkan_total belum diisi")


def rupiah(nilai: int) -> str:
    """Sudah diberikan: 53000 -> 'Rp53.000'."""
    return "Rp" + f"{nilai:,}".replace(",", ".")


def subtotal_dari(isi: list[tuple[int, int]]) -> int:
    """Sudah diberikan: jumlahkan harga x jumlah."""
    subtotal = 0
    for harga, jumlah in isi:
        subtotal = subtotal + harga * jumlah
    return subtotal


def main() -> None:
    for nama, isi in KASUS_UJI:
        subtotal = validasi_subtotal(subtotal_dari(isi))
        diskon = hitung_diskon(subtotal)
        total = hitung_total(subtotal, diskon)
        tampilkan_total(nama, subtotal, diskon, total)


if __name__ == "__main__":
    main()
