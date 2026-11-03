// latihan_template.rs — Praktik (versi Rust).
// Jalankan: rustc latihan_template.rs -o latihan && ./latihan
//
// Tugas Anda: isi empat TODO di bawah sehingga keluaran program ini PERSIS
// sama dengan keluaran 02a_sebelum_refaktor.rs.
//
// Aturan:
//   - Jangan mengubah main() dan jangan mengubah teks yang dicetak.
//   - Setiap fungsi hanya mengerjakan satu hal.
//   - Fungsi yang menghitung tidak boleh mencetak; yang mencetak tidak boleh
//     menghitung sendiri.

const KERANJANG: [(&str, &[(i64, i64)]); 3] = [
    ("A", &[(15_000, 2), (5_000, 3), (8_000, 1)]),
    ("B", &[(125_000, 2), (250_000, 1)]),
    ("C", &[(200_000, 1), (100_000, 1)]),
];

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

// TODO 1 — jumlahkan harga x jumlah untuk seluruh baris keranjang.
fn hitung_subtotal(_isi: &[(i64, i64)]) -> i64 {
    todo!("TODO 1: hitung_subtotal belum diisi")
}

// TODO 2 — kembalikan diskon dalam rupiah.
// Aturan (batas bawah inklusif):
//   < Rp100.000             -> 0%
//   Rp100.000 - Rp249.999   -> 5%
//   Rp250.000 - Rp499.999   -> 10%
//   >= Rp500.000            -> 15%
fn hitung_diskon(_subtotal: i64) -> i64 {
    todo!("TODO 2: hitung_diskon belum diisi")
}

// TODO 3 — PPN 11% dari nilai setelah diskon.
fn hitung_ppn(_dasar: i64) -> i64 {
    todo!("TODO 3: hitung_ppn belum diisi")
}

// TODO 4 — prosedur: cetak satu nota, tidak mengembalikan nilai.
fn cetak_nota(_nama: &str, _isi: &[(i64, i64)]) {
    todo!("TODO 4: cetak_nota belum diisi")
}

fn main() {
    let mut total_terbesar = 0;
    let mut nama_terbesar = "";

    for (nama, isi) in KERANJANG {
        cetak_nota(nama, isi);
        let subtotal = hitung_subtotal(isi);
        let diskon = hitung_diskon(subtotal);
        let mut total = subtotal - diskon;
        total += hitung_ppn(total);

        if total > total_terbesar {
            total_terbesar = total;
            nama_terbesar = nama;
        }
    }

    println!(
        "Transaksi terbesar: {} dengan total {}",
        nama_terbesar,
        rupiah(total_terbesar)
    );
}
