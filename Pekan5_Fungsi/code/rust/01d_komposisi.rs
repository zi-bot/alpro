// 01d — Kasus utuh: lima fungsi, satu alur utama (versi Rust).
// Jalankan: rustc 01d_komposisi.rs -o 01d && ./01d
//
// Pembagian tanggung jawab:
//   hitung_subtotal(keranjang)      -> i64   menjumlahkan isi keranjang
//   hitung_diskon(subtotal)         -> i64   memutuskan tingkat diskon
//   hitung_ppn(dasar)               -> i64   pajak 11% dari nilai setelah diskon
//   hitung_total(subtotal, diskon)  -> i64   memanggil hitung_ppn, lalu menjumlah
//   tampilkan_hasil(...)            -> ()    prosedur: hanya mencetak

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

fn hitung_subtotal(keranjang: &[(i64, i64)]) -> i64 {
    let mut subtotal = 0;
    for &(harga, jumlah) in keranjang {
        subtotal += harga * jumlah;
    }
    subtotal
}

fn hitung_diskon(subtotal: i64) -> i64 {
    if subtotal < 0 {
        panic!("subtotal tidak boleh negatif");
    }
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

fn hitung_ppn(dasar: i64) -> i64 {
    dasar * 11 / 100
}

fn hitung_total(subtotal: i64, diskon: i64) -> i64 {
    let dasar = subtotal - diskon;
    dasar + hitung_ppn(dasar)
}

fn tampilkan_hasil(nama: &str, keranjang: &[(i64, i64)]) {
    let subtotal = hitung_subtotal(keranjang);
    let diskon = hitung_diskon(subtotal);
    let total = hitung_total(subtotal, diskon);

    println!("NOTA {}", nama);
    println!("----------------------------------");
    for &(harga, jumlah) in keranjang {
        println!(
            "  {:>11} x {} = {:>11}",
            rupiah(harga),
            jumlah,
            rupiah(harga * jumlah)
        );
    }
    println!("----------------------------------");
    println!("  {:<20}{:>14}", "Subtotal", rupiah(subtotal));
    println!("  {:<20}{:>14}", "Diskon", rupiah(diskon));
    println!(
        "  {:<20}{:>14}",
        "PPN 11%",
        rupiah(hitung_ppn(subtotal - diskon))
    );
    println!("  {:<20}{:>14}", "TOTAL", rupiah(total));
    println!();
}

fn main() {
    let keranjang_a = [(15_000, 2), (5_000, 3), (8_000, 1)];
    let keranjang_b = [(125_000, 2), (250_000, 1)];

    tampilkan_hasil("KERANJANG A", &keranjang_a);
    tampilkan_hasil("KERANJANG B", &keranjang_b);
}
