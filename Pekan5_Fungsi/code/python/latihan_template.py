# latihan_template.py — Praktik (50 menit).
# Jalankan: python3 latihan_template.py
#
# Tugas Anda: isi empat TODO di bawah sehingga keluaran program ini PERSIS
# sama dengan keluaran 02a_sebelum_refaktor.py.
#
# Aturan:
#   - Jangan mengubah main() dan jangan mengubah teks yang dicetak.
#   - Setiap fungsi hanya mengerjakan satu hal.
#   - Fungsi yang menghitung tidak boleh mencetak; yang mencetak tidak boleh
#     menghitung sendiri.
#
# Periksa hasilnya:
#   python3 latihan_template.py > hasil_saya.txt
#   python3 02a_sebelum_refaktor.py > acuan.txt
#   diff acuan.txt hasil_saya.txt        (harus kosong)

KERANJANG = [
    ("A", [(15_000, 2), (5_000, 3), (8_000, 1)]),
    ("B", [(125_000, 2), (250_000, 1)]),
    ("C", [(200_000, 1), (100_000, 1)]),
]


def rupiah(nilai: int) -> str:
    """Sudah diberikan: 53000 -> 'Rp53.000'."""
    return "Rp" + f"{nilai:,}".replace(",", ".")


def hitung_subtotal(isi: list[tuple[int, int]]) -> int:
    """TODO 1 — jumlahkan harga x jumlah untuk seluruh baris keranjang.

    Masukan : daftar pasangan (harga, jumlah)
    Keluaran: satu bilangan bulat, subtotal
    """
    raise NotImplementedError("TODO 1: hitung_subtotal belum diisi")


def hitung_diskon(subtotal: int) -> int:
    """TODO 2 — kembalikan diskon dalam rupiah.

    Aturan (batas bawah inklusif):
        < Rp100.000             -> 0%
        Rp100.000 - Rp249.999   -> 5%
        Rp250.000 - Rp499.999   -> 10%
        >= Rp500.000            -> 15%
    """
    raise NotImplementedError("TODO 2: hitung_diskon belum diisi")


def hitung_ppn(dasar: int) -> int:
    """TODO 3 — PPN 11% dari nilai setelah diskon."""
    raise NotImplementedError("TODO 3: hitung_ppn belum diisi")


def cetak_nota(nama: str, isi: list[tuple[int, int]]) -> None:
    """TODO 4 — prosedur: cetak satu nota, tidak mengembalikan nilai.

    Bentuk keluaran harus sama dengan 02a_sebelum_refaktor.py:
    judul 'NOTA KERANJANG X', garis 34 tanda minus, satu baris per barang,
    lalu Subtotal / Diskon / PPN 11% / TOTAL dengan lebar kolom yang sama.
    """
    raise NotImplementedError("TODO 4: cetak_nota belum diisi")


def main() -> None:
    total_terbesar = 0
    nama_terbesar = ""

    for nama, isi in KERANJANG:
        cetak_nota(nama, isi)
        subtotal = hitung_subtotal(isi)
        total = subtotal - hitung_diskon(subtotal)
        total = total + hitung_ppn(total)
        if total > total_terbesar:
            total_terbesar = total
            nama_terbesar = nama

    print("Transaksi terbesar:", nama_terbesar, "dengan total", rupiah(total_terbesar))


main()
