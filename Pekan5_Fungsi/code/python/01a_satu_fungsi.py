# 01a — Satu fungsi, satu nilai balik.
# Jalankan: python3 01a_satu_fungsi.py
#
# Perhatikan tiga hal pada hitung_subtotal:
#   1. def     — fungsi diberi nama, bukan blok anonim.
#   2. parameter harga dan jumlah — nilai masuk, sudah bertipe.
#   3. return  — nilai keluar; pemanggil yang memutuskan mau diapakan.


def rupiah(nilai: int) -> str:
    """Ubah bilangan bulat menjadi teks rupiah: 53000 -> Rp53.000."""
    return "Rp" + f"{nilai:,}".replace(",", ".")


def hitung_subtotal(harga: int, jumlah: int) -> int:
    """Subtotal satu baris keranjang: harga dikali jumlah."""
    return harga * jumlah


def main() -> None:
    harga = 15_000
    jumlah = 2
    subtotal = hitung_subtotal(harga, jumlah)

    print("harga    :", rupiah(harga))
    print("jumlah   :", jumlah)
    print("subtotal :", rupiah(subtotal))


main()
