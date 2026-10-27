// Latihan terbimbing - kasir minimarket (versi Rust)
// Pembanding untuk latihan_template.py. Tidak wajib dikerjakan.
//
// Jalankan program: rustc latihan.rs -o latihan && ./latihan
// Jalankan uji    : rustc --test latihan.rs -o latihan_test && ./latihan_test
//
// Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
// Kunci ada di solusi/latihan.rs.

#[derive(Debug, Clone, Copy)]
struct Hasil {
    subtotal: i64,
    diskon: i64,
    pajak: i64,
    total: i64,
    ditolak: i64,
}

fn hitung(keranjang: &[(i64, i64)]) -> Hasil {
    let mut subtotal: i64 = 0;
    let ditolak: i64 = 0;

    for &(harga, jumlah) in keranjang {
        // TODO 1 - Validasi baris
        // Bila jumlah <= 0 atau harga < 0:
        //   1. ubah `let ditolak` di atas menjadi `let mut ditolak`
        //   2. tambahkan 1 ke ditolak
        //   3. lewati baris ini dengan continue
        subtotal += harga * jumlah;
    }

    // TODO 2 - Diskon bertingkat
    // Ganti 0 dengan persen diskon: 0, 5, 10, atau 15.
    // Batas: < 100.000, < 250.000, < 500.000, sisanya.
    let persen: i64 = 0;

    // Pembulatan setengah ke atas: tambah 50 lalu bagi 100.
    let diskon: i64 = (subtotal * persen + 50) / 100;

    // TODO 3 - Pajak
    // Ganti 0 dengan ((subtotal - diskon) * 11 + 50) / 100
    let pajak: i64 = 0;

    // TODO 4 - Total
    // Ganti 0 dengan subtotal - diskon + pajak
    let total: i64 = 0;

    Hasil {
        subtotal,
        diskon,
        pajak,
        total,
        ditolak,
    }
}

// Format ribuan: 58830 menjadi 58.830
fn rupiah(nilai: i64) -> String {
    if nilai < 0 {
        return format!("-{}", rupiah(-nilai));
    }
    let teks = nilai.to_string();
    let panjang = teks.len();
    let mut hasil = String::new();

    for (i, c) in teks.chars().enumerate() {
        if i > 0 && (panjang - i) % 3 == 0 {
            hasil.push('.');
        }
        hasil.push(c);
    }

    hasil
}

fn daftar_keranjang() -> Vec<Vec<(i64, i64)>> {
    vec![
        vec![(15_000, 2), (5_000, 3), (8_000, 1)],
        vec![(60_000, 2)],
        vec![(70_000, 2), (80_000, 2)],
        vec![(150_000, 4)],
        vec![],
        vec![(50_000, 2), (10_000, 0), (30_000, -1)],
    ]
}

fn main() {
    for (nomor, keranjang) in daftar_keranjang().iter().enumerate() {
        let h = hitung(keranjang);

        println!("Keranjang {}", nomor + 1);
        println!("  Subtotal : Rp{}", rupiah(h.subtotal));
        println!("  Diskon   : Rp{}", rupiah(h.diskon));
        println!("  Pajak    : Rp{}", rupiah(h.pajak));
        println!("  Total    : Rp{}", rupiah(h.total));
        println!("  Ditolak  : {} baris", h.ditolak);
    }

    println!();
    println!("Jalankan rustc --test latihan.rs -o latihan_test && ./latihan_test untuk memeriksa.");
}

#[cfg(test)]
mod tests {
    use super::*;

    // Uji tabel: setiap baris berisi keranjang dan hasil yang diharapkan.
    #[test]
    fn tabel_kasus() {
        let harapan: Vec<(i64, i64, i64, i64, i64)> = vec![
            (53_000, 0, 5_830, 58_830, 0),
            (120_000, 6_000, 12_540, 126_540, 0),
            (300_000, 30_000, 29_700, 299_700, 0),
            (600_000, 90_000, 56_100, 566_100, 0),
            (0, 0, 0, 0, 0),
            (100_000, 5_000, 10_450, 105_450, 2),
        ];

        let daftar = daftar_keranjang();
        assert_eq!(daftar.len(), harapan.len(), "jumlah kasus tidak sama");

        for (nomor, (keranjang, harap)) in daftar.iter().zip(harapan.iter()).enumerate() {
            let h = hitung(keranjang);
            let dapat = (h.subtotal, h.diskon, h.pajak, h.total, h.ditolak);
            assert_eq!(dapat, *harap, "keranjang {}", nomor + 1);
        }
    }
}
