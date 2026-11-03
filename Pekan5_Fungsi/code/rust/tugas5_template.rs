// tugas5_template.rs — Tugas 5 (individu, rubrik A2, 5% nilai akhir, versi Rust).
// Jalankan: rustc tugas5_template.rs -o tugas5 && ./tugas5
//
// Kerjakan empat TODO di bawah.

const KASUS_UJI: [(&str, &[(i64, i64)]); 10] = [
    ("K-01", &[(50_000, 1)]),
    ("K-02", &[(99_999, 1)]),
    ("K-03", &[(100_000, 1)]),
    ("K-04", &[(100_001, 1)]),
    ("K-05", &[(249_999, 1)]),
    ("K-06", &[(250_000, 1)]),
    ("K-07", &[(499_999, 1)]),
    ("K-08", &[(500_000, 1)]),
    ("K-09", &[(125_000, 2), (250_000, 1)]),
    ("K-10", &[(0, 0)]),
];

// TODO 1 — tolak subtotal negatif.
// Kembalikan subtotal apa adanya bila sah; panic dengan pesan
// "subtotal tidak boleh negatif" bila tidak sah.
fn validasi_subtotal(_subtotal: i64) -> i64 {
    todo!("TODO 1: validasi_subtotal belum diisi")
}

// TODO 2 — diskon dalam rupiah, batas bawah inklusif.
// < Rp100.000             -> 0%
// Rp100.000 - Rp249.999   -> 5%
// Rp250.000 - Rp499.999   -> 10%
// >= Rp500.000            -> 15%
fn hitung_diskon(_subtotal: i64) -> i64 {
    todo!("TODO 2: hitung_diskon belum diisi")
}

// TODO 3 — total bayar: nilai setelah diskon ditambah PPN 11%.
fn hitung_total(_subtotal: i64, _diskon: i64) -> i64 {
    todo!("TODO 3: hitung_total belum diisi")
}

// TODO 4 — prosedur penyajian. Cetak format:
//   K-01  subtotal Rp50.000  diskon Rp0  total Rp55.500
fn tampilkan_total(_nama: &str, _subtotal: i64, _diskon: i64, _total: i64) {
    todo!("TODO 4: tampilkan_total belum diisi")
}

#[allow(dead_code)]
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

fn subtotal_dari(isi: &[(i64, i64)]) -> i64 {
    let mut subtotal = 0;
    for &(harga, jumlah) in isi {
        subtotal += harga * jumlah;
    }
    subtotal
}

fn main() {
    for (nama, isi) in KASUS_UJI {
        let subtotal = validasi_subtotal(subtotal_dari(isi));
        let diskon = hitung_diskon(subtotal);
        let total = hitung_total(subtotal, diskon);
        tampilkan_total(nama, subtotal, diskon, total);
    }
}
