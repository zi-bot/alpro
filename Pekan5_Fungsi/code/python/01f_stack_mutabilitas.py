"""01f_stack_mutabilitas.py

Demonstrasi konsep teori fungsi tingkat lanjut:
1. Call Stack & Stack Frames (Activation Records) melalui rekursi
2. Mekanisme Evaluasi Parameter (Pass-by-Object-Reference / Call-by-Sharing):
   - Go: Pass-by-value murni (nilai disalin utuh, variabel pemanggil selalu aman).
   - Rust: Ownership & Borrow Pointer (& untuk baca, &mut untuk izin mutasi ketat).
   - Python: Referensi objek diteruskan. Perilaku ditentukan oleh mutabilitas objek:
     * Immutable (int, str, tuple): aman, manipulasi lokal membuat objek baru.
     * Mutable (list, dict, set): berbahaya, mutasi in-place merusak data pemanggil!
3. Pure Function vs Side Effects (Efek Samping)
"""

import inspect


def contoh_stack_frame(tingkat: int) -> None:
    """Mencetak nama fungsi dan kedalaman stack frame saat dieksekusi."""
    stack = inspect.stack()
    nama_fungsi = stack[0].function
    kedalaman = len(stack)
    print(f"  [Frame tingkat {tingkat}] fungsi: {nama_fungsi}, kedalaman stack: {kedalaman}")
    if tingkat > 1:
        contoh_stack_frame(tingkat - 1)
    # Saat fungsi selesai, frame di-pop dari stack
    print(f"  [Pop tingkat {tingkat}] selesai dan kembali.")


def ubah_immutable(x: int) -> None:
    """Bilangan bulat bersifat immutable; manipulasi hanya mengubah referensi lokal."""
    x = x + 100
    print(f"    Di dalam ubah_immutable: x = {x}")


def ubah_mutable(daftar: list[int]) -> None:
    """List bersifat mutable; perubahan in-place mempengaruhi objek pemanggil."""
    daftar.append(999)
    print(f"    Di dalam ubah_mutable: daftar = {daftar}")


def hitung_diskon_murni(subtotal: int) -> int:
    """Pure Function: Deterministik, tanpa I/O, tanpa mengubah status luar."""
    if subtotal < 100_000:
        return 0
    elif subtotal < 250_000:
        return subtotal * 5 // 100
    elif subtotal < 500_000:
        return subtotal * 10 // 100
    else:
        return subtotal * 15 // 100


def main() -> None:
    print("=== 1. DEMONSTRASI CALL STACK & STACK FRAMES ===")
    contoh_stack_frame(3)

    print("\n=== 2. PASS-BY-OBJECT-REFERENCE (IMMUTABLE VS MUTABLE) ===")
    angka = 50
    print(f"Sebelum ubah_immutable: angka = {angka}")
    ubah_immutable(angka)
    print(f"Setelah ubah_immutable : angka = {angka} (TIDAK BERUBAH)")

    keranjang_item = [15000, 5000]
    print(f"Sebelum ubah_mutable  : keranjang_item = {keranjang_item}")
    ubah_mutable(keranjang_item)
    print(f"Setelah ubah_mutable   : keranjang_item = {keranjang_item} (BERUBAH: side effect!)")

    print("\n=== 3. PURE FUNCTION (DETERMINISTIK & TANPA SIDE EFFECT) ===")
    subtotal = 200_000
    d1 = hitung_diskon_murni(subtotal)
    d2 = hitung_diskon_murni(subtotal)
    print(f"Panggilan 1 hitung_diskon_murni({subtotal}) = {d1}")
    print(f"Panggilan 2 hitung_diskon_murni({subtotal}) = {d2} (Selalu konsisten, bebas galat)")


if __name__ == "__main__":
    main()
