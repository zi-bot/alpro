// 01c — Diskon bertingkat: satu fungsi, empat tingkat (versi Rust).
// Jalankan: rustc 01c_diskon_bertingkat.rs -o 01c && ./01c
//
// Aturan (batas bawah inklusif, batas atas milik tingkat berikutnya):
//   < Rp100.000                 -> 0%
//   Rp100.000 - Rp249.999       -> 5%
//   Rp250.000 - Rp499.999       -> 10%
//   >= Rp500.000                -> 15%

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

fn persen(subtotal: i64) -> &'static str {
    if subtotal < 100_000 {
        "0%"
    } else if subtotal < 250_000 {
        "5%"
    } else if subtotal < 500_000 {
        "10%"
    } else {
        "15%"
    }
}

fn main() {
    let uji = [
        0, 53_000, 99_999, 100_000, 249_999, 250_000, 499_999, 500_000, 1_250_000,
    ];

    println!("SUBTOTAL      TINGKAT   DISKON");
    println!("----------------------------------");
    for &subtotal in &uji {
        println!(
            "{:>13}  {:>6}  {:>10}",
            rupiah(subtotal),
            persen(subtotal),
            rupiah(hitung_diskon(subtotal))
        );
    }

    println!();
    println!("Dua batas yang paling mudah salah:");
    println!(
        "   99.999 -> {}  (di bawah batas, tetap 0%)",
        rupiah(hitung_diskon(99_999))
    );
    println!(
        "  100.000 -> {}  (tepat di batas, sudah 5%)",
        rupiah(hitung_diskon(100_000))
    );
    println!();
    println!("Pecahan rupiah dibuang ke bawah, bukan dibulatkan:");
    println!(
        "  249.999 x 5% = 12.499,95 -> {}",
        rupiah(hitung_diskon(249_999))
    );
}
