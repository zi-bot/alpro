// 01a — Satu fungsi, satu nilai balik (versi Rust).
// Jalankan: rustc 01a_satu_fungsi.rs -o 01a && ./01a
//
// Perhatikan tiga hal pada hitung_subtotal:
//   1. fn      — fungsi diberi nama secara eksplisit.
//   2. parameter harga dan jumlah bertipe i64 (tipe statis).
//   3. nilai balik bertipe i64; ekspresi terakhir tanpa titik-koma otomatis di-return.

fn rupiah(nilai: i64) -> String {
    let teks = nilai.abs().to_string();
    let mut keluar = String::new();
    for (i, c) in teks.chars().enumerate() {
        if i > 0 && (teks.len() - i) % 3 == 0 {
            keluar.push('.');
        }
        keluar.push(c);
    }
    let tanda = if nilai < 0 { "-" } else { "" };
    format!("Rp{}{}", tanda, keluar)
}

fn hitung_subtotal(harga: i64, jumlah: i64) -> i64 {
    harga * jumlah
}

fn main() {
    let harga = 15_000;
    let jumlah = 2;
    let subtotal = hitung_subtotal(harga, jumlah);

    println!("harga    : {}", rupiah(harga));
    println!("jumlah   : {}", jumlah);
    println!("subtotal : {}", rupiah(subtotal));
}
