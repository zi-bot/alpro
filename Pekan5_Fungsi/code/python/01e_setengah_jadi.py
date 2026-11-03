# 01e — Tiga versi setengah jadi: gejala, penyebab, perbaikan.
# Jalankan: python3 01e_setengah_jadi.py
#
# Ketiganya "jalan" tanpa galat. Yang salah adalah hasilnya di sekitar batas,
# dan itu tidak akan terlihat kalau yang diuji hanya satu nilai tengah.

BATAS = [99_999, 100_000, 249_999, 250_000, 499_999, 500_000]


def seharusnya(subtotal: int) -> int:
    """Versi benar, dipakai sebagai pembanding."""
    if subtotal < 100_000:
        return 0
    if subtotal < 250_000:
        return subtotal * 5 // 100
    if subtotal < 500_000:
        return subtotal * 10 // 100
    return subtotal * 15 // 100


def v1_mencetak_bukan_mengembalikan(subtotal: int) -> int:
    """Salah 1: hasil dicetak, bukan dikembalikan. Pemanggil menerima None."""
    print(f"  (di dalam fungsi, angka ini tercetak: {subtotal * 5 // 100})")


def v2_batas_ikut_tergeser(subtotal: int) -> int:
    """Salah 2: memakai <= sehingga nilai tepat di batas jatuh ke tingkat bawah."""
    if subtotal <= 100_000:
        return 0
    if subtotal <= 250_000:
        return subtotal * 5 // 100
    if subtotal <= 500_000:
        return subtotal * 10 // 100
    return subtotal * 15 // 100


def v3_urutan_terbalik(subtotal: int) -> int:
    """Salah 3: diperiksa dari batas terkecil dengan >=, cabang pertama selalu menang."""
    if subtotal >= 100_000:
        return subtotal * 5 // 100
    if subtotal >= 250_000:
        return subtotal * 10 // 100
    if subtotal >= 500_000:
        return subtotal * 15 // 100
    return 0


def rupiah(nilai) -> str:
    if nilai is None:
        return "None"
    return "Rp" + f"{nilai:,}".replace(",", ".")


def main() -> None:
    print("SALAH 1 — mencetak di dalam fungsi, bukan mengembalikan")
    print("-" * 58)
    hasil = v1_mencetak_bukan_mengembalikan(100_000)
    print("  nilai yang diterima pemanggil:", hasil)
    print("  akibatnya: tidak bisa dipakai menghitung apa pun.")
    print()

    for judul, fungsi in (("SALAH 2 — batas memakai <=", v2_batas_ikut_tergeser),
                          ("SALAH 3 — urutan kondisi terbalik", v3_urutan_terbalik)):
        print(judul)
        print("-" * 58)
        print(f"  {'SUBTOTAL':>12}  {'SEHARUSNYA':>11}  {'HASIL':>11}  STATUS")
        for s in BATAS:
            benar = seharusnya(s)
            hasil = fungsi(s)
            tanda = "ok" if hasil == benar else "SALAH"
            print(f"  {rupiah(s):>12}  {rupiah(benar):>11}  {rupiah(hasil):>11}  {tanda}")
        print()

    print("Yang harus terlihat:")
    print("  SALAH 2  Rp100.000, Rp250.000, dan Rp500.000 salah: tepat di batas,")
    print("           nilainya masih dihitung dengan tingkat di bawahnya.")
    print("  SALAH 3  benar hanya untuk tingkat 5% (Rp99.999, Rp100.000, Rp249.999).")
    print("           Mulai Rp250.000 hasilnya selalu 5%, karena cabang pertama")
    print("           sudah menang sebelum cabang berikutnya diperiksa.")


main()
