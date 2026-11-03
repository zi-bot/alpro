// main.rs — pembanding Rust untuk kasus yang sama.
// Jalankan: rustc main.rs -o diskon && ./diskon
//
// Yang berbeda dari Python:
//   - tipe parameter dan tipe nilai balik wajib ditulis (subtotal: i64) -> i64
//   - kata kunci return tidak diperlukan pada ekspresi terakhir fungsi
//   - prosedur tetap fungsi, tetapi nilai baliknya () — "unit", kosong

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

fn hitung_diskon(subtotal: i64) -> i64 {
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

fn hitung_total(subtotal: i64, diskon: i64) -> i64 {
    let dasar = subtotal - diskon;
    dasar + dasar * 11 / 100
}

// Prosedur: nilai balik () — tidak ada yang diserahkan ke pemanggil.
fn tampilkan_total(nama: &str, subtotal: i64, diskon: i64, total: i64) {
    println!(
        "{}  subtotal {}  diskon {}  total {}",
        nama,
        rupiah(subtotal),
        rupiah(diskon),
        rupiah(total)
    );
}

fn main() {
    let kasus: [(&str, i64); 10] = [
        ("K-01", 50_000),
        ("K-02", 99_999),
        ("K-03", 100_000),
        ("K-04", 100_001),
        ("K-05", 249_999),
        ("K-06", 250_000),
        ("K-07", 499_999),
        ("K-08", 500_000),
        ("K-09", 500_000),
        ("K-10", 0),
    ];

    for (nama, subtotal) in kasus {
        let diskon = hitung_diskon(subtotal);
        let total = hitung_total(subtotal, diskon);
        tampilkan_total(nama, subtotal, diskon, total);
    }
}
