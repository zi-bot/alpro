// 01f — Call Stack, Mutabilitas, dan Pure Function (versi Rust).
// Jalankan: rustc 01f_stack_mutabilitas.rs -o 01f && ./01f
//
// Demonstrasi konsep fungsi tingkat lanjut:
// 1. Call Stack & Stack Frames (Activation Records) melalui rekursi
// 2. Evaluasi Parameter (Pass-by-value vs Mutable Reference `&mut`):
//    - Go: Pass-by-value murni (nilai variabel disalin utuh).
//    - Rust: Ownership & Borrow Pointer (& untuk baca, &mut untuk izin mutasi ketat).
//    - Python: Call-by-sharing (referensi objek diteruskan; mutabilitas menentukan efek).
// 3. Pure Function vs Efek Samping (Side Effects)

fn contoh_stack_frame(tingkat: usize, kedalaman: usize) {
    println!(
        "  [Frame tingkat {}] fungsi: contoh_stack_frame, kedalaman stack: {}",
        tingkat, kedalaman
    );
    if tingkat > 1 {
        contoh_stack_frame(tingkat - 1, kedalaman + 1);
    }
    println!("  [Pop tingkat {}] selesai dan kembali.", tingkat);
}

fn ubah_immutable(mut x: i64) {
    // Parameter x disalin (pass-by-value); mut hanya berlaku lokal
    x += 100;
    println!("    Di dalam ubah_immutable: x = {}", x);
}

fn ubah_mutable(daftar: &mut Vec<i64>) {
    // Borrowing secara mutable (&mut) memungkinkan perubahan in-place pada memori pemanggil
    daftar.push(999);
    println!("    Di dalam ubah_mutable: daftar = {:?}", daftar);
}

fn hitung_diskon_murni(subtotal: i64) -> i64 {
    // Pure Function: Deterministik, tanpa I/O, tanpa mengubah status luar
    if subtotal < 100_000 {
        0
    } else if subtotal < 250_000 {
        subtotal * 5 / 100
    } else if subtotal < 500_000 {
        subtotal * 10 / 100
    } else {
        subtotal * 15 / 100
    }
}

fn main() {
    println!("=== 1. DEMONSTRASI CALL STACK & STACK FRAMES ===");
    contoh_stack_frame(3, 3);

    println!("\n=== 2. PASS-BY-VALUE VS MUTABLE REFERENCE (&MUT) ===");
    let angka = 50;
    println!("Sebelum ubah_immutable: angka = {}", angka);
    ubah_immutable(angka);
    println!("Setelah ubah_immutable : angka = {} (TIDAK BERUBAH)", angka);

    let mut keranjang_item = vec![15000, 5000];
    println!(
        "Sebelum ubah_mutable  : keranjang_item = {:?}",
        keranjang_item
    );
    ubah_mutable(&mut keranjang_item);
    println!(
        "Setelah ubah_mutable   : keranjang_item = {:?} (BERUBAH: side effect!)",
        keranjang_item
    );

    println!("\n=== 3. PURE FUNCTION (DETERMINISTIK & TANPA SIDE EFFECT) ===");
    let subtotal = 200_000;
    let d1 = hitung_diskon_murni(subtotal);
    let d2 = hitung_diskon_murni(subtotal);
    println!("Panggilan 1 hitung_diskon_murni({}) = {}", subtotal, d1);
    println!(
        "Panggilan 2 hitung_diskon_murni({}) = {} (Selalu konsisten, bebas galat)",
        subtotal, d2
    );
}
