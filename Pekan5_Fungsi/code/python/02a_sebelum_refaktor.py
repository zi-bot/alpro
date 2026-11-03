# 02a — Program sebelum dipecah: semuanya di satu alur.
# Jalankan: python3 02a_sebelum_refaktor.py
#
# Program ini BENAR. Jangan diubah dulu — catat keluarannya, karena keluaran
# itulah bukti bahwa hasil refaktor Anda nanti tidak merusak apa pun.
#
# Yang membuatnya sulit dirawat:
#   - aturan diskon ditulis dua kali (di sini dan di dalam loop pencetakan),
#   - tidak ada bagian yang bisa diuji sendiri,
#   - untuk mengubah satu aturan, tiga tempat harus dicari.

KERANJANG = [
    ("A", [(15_000, 2), (5_000, 3), (8_000, 1)]),
    ("B", [(125_000, 2), (250_000, 1)]),
    ("C", [(200_000, 1), (100_000, 1)]),
]


def main() -> None:
    total_terbesar = 0
    nama_terbesar = ""

    for nama, isi in KERANJANG:
        subtotal = 0
        for harga, jumlah in isi:
            subtotal = subtotal + harga * jumlah

        if subtotal < 100_000:
            diskon = 0
        elif subtotal < 250_000:
            diskon = subtotal * 5 // 100
        elif subtotal < 500_000:
            diskon = subtotal * 10 // 100
        else:
            diskon = subtotal * 15 // 100

        dasar = subtotal - diskon
        ppn = dasar * 11 // 100
        total = dasar + ppn

        print(f"NOTA KERANJANG {nama}")
        print("-" * 34)
        for harga, jumlah in isi:
            print(f"  {'Rp' + format(harga, ',').replace(',', '.'):>11} x {jumlah}"
                  f" = {'Rp' + format(harga * jumlah, ',').replace(',', '.'):>11}")
        print("-" * 34)
        print(f"  {'Subtotal':<20}{'Rp' + format(subtotal, ',').replace(',', '.'):>14}")
        print(f"  {'Diskon':<20}{'Rp' + format(diskon, ',').replace(',', '.'):>14}")
        print(f"  {'PPN 11%':<20}{'Rp' + format(ppn, ',').replace(',', '.'):>14}")
        print(f"  {'TOTAL':<20}{'Rp' + format(total, ',').replace(',', '.'):>14}")
        print()

        if total > total_terbesar:
            total_terbesar = total
            nama_terbesar = nama

    print("Transaksi terbesar:", nama_terbesar,
          "dengan total Rp" + format(total_terbesar, ",").replace(",", "."))


main()
