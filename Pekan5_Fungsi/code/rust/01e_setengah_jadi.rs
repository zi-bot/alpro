// 01e — Tiga versi setengah jadi: gejala, penyebab, perbaikan (versi Rust).
// Jalankan: rustc 01e_setengah_jadi.rs -o 01e && ./01e

const BATAS: [i64; 6] = [99_999, 100_000, 249_999, 250_000, 499_999, 500_000];

fn seharusnya(subtotal: i64) -> i64 {
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

fn v1_mencetak_bukan_mengembalikan(subtotal: i64) {
    println!(
        "  (di dalam fungsi, angka ini tercetak: {})",
        subtotal * 5 / 100
    );
}

fn v2_batas_ikut_tergeser(subtotal: i64) -> i64 {
    if subtotal <= 100_000 {
        0
    } else if subtotal <= 250_000 {
        subtotal * 5 / 100
    } else if subtotal <= 500_000 {
        subtotal * 10 / 100
    } else {
        subtotal * 15 / 100
    }
}

fn v3_urutan_terbalik(subtotal: i64) -> i64 {
    if subtotal >= 100_000 {
        subtotal * 5 / 100
    } else if subtotal >= 250_000 {
        subtotal * 10 / 100
    } else if subtotal >= 500_000 {
        subtotal * 15 / 100
    } else {
        0
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

fn main() {
    println!("SALAH 1 — mencetak di dalam fungsi, bukan mengembalikan");
    println!("----------------------------------------------------------");
    let hasil = v1_mencetak_bukan_mengembalikan(100_000);
    println!("  nilai yang diterima pemanggil: {:?}", hasil);
    println!("  akibatnya: tidak bisa dipakai menghitung apa pun.");
    println!();

    let pengujian: [(&str, fn(i64) -> i64); 2] = [
        ("SALAH 2 — batas memakai <=", v2_batas_ikut_tergeser),
        ("SALAH 3 — urutan kondisi terbalik", v3_urutan_terbalik),
    ];

    for (judul, fungsi) in pengujian {
        println!("{}", judul);
        println!("----------------------------------------------------------");
        println!(
            "  {:>12}  {:>11}  {:>11}  STATUS",
            "SUBTOTAL", "SEHARUSNYA", "HASIL"
        );
        for &s in &BATAS {
            let benar = seharusnya(s);
            let hasil = fungsi(s);
            let tanda = if hasil == benar { "ok" } else { "SALAH" };
            println!(
                "  {:>12}  {:>11}  {:>11}  {}",
                rupiah(s),
                rupiah(benar),
                rupiah(hasil),
                tanda
            );
        }
        println!();
    }

    println!("Yang harus terlihat:");
    println!("  SALAH 2  Rp100.000, Rp250.000, dan Rp500.000 salah: tepat di batas,");
    println!("           nilainya masih dihitung dengan tingkat di bawahnya.");
    println!("  SALAH 3  benar hanya untuk tingkat 5% (Rp99.999, Rp100.000, Rp249.999).");
    println!("           Mulai Rp250.000 hasilnya selalu 5%, karena cabang pertama");
    println!("           sudah menang sebelum cabang berikutnya diperiksa.");
}
