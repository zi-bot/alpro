// Tugas 4 - checkout dan aturan voucher (versi Rust)
// Pembanding untuk tugas4_template.py. Tidak wajib dikerjakan.
//
// Jalankan program: rustc tugas4.rs -o tugas4 && ./tugas4
// Jalankan uji    : rustc --test tugas4.rs -o tugas4_test && ./tugas4_test
//
// Lengkapi 4 TODO di dalam fungsi hitung(), lalu jalankan uji.
// Kunci ada di solusi/tugas4.rs.

#[derive(Debug, Clone, Copy)]
struct Hasil {
    subtotal: i64,
    voucher: i64,
    ongkir: i64,
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

    // TODO 2 - Voucher
    // Ganti 0 dengan voucher 10% bila subtotal > 100.000.
    // Pembulatan setengah ke atas: (subtotal * 10 + 50) / 100
    let voucher: i64 = 0;

    // TODO 3 - Ongkir
    // Ganti 0 dengan 15.000, kecuali gratis bila subtotal >= 150.000
    // atau keranjang tidak berisi baris valid.
    let ongkir: i64 = 0;

    // TODO 4 - Total
    // Ganti 0 dengan subtotal - voucher + ongkir.
    let total: i64 = 0;

    Hasil {
        subtotal,
        voucher,
        ongkir,
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

fn daftar_kasus() -> Vec<(&'static str, Vec<(i64, i64)>)> {
    vec![
        ("di bawah batas voucher", vec![(25_000, 2), (20_000, 1)]),
        ("tepat Rp100.000", vec![(100_000, 1)]),
        ("Rp100.001", vec![(100_001, 1)]),
        ("pembulatan setengah ke atas", vec![(100_005, 1)]),
        ("tepat Rp149.999", vec![(149_999, 1)]),
        ("tepat Rp150.000", vec![(150_000, 1)]),
        ("di atas Rp150.000", vec![(50_000, 4)]),
        (
            "dua baris ditolak",
            vec![(50_000, 2), (10_000, 0), (30_000, -1)],
        ),
        ("keranjang kosong", vec![]),
        ("semua baris ditolak", vec![(-5_000, 2)]),
    ]
}

fn main() {
    for (nomor, (nama, keranjang)) in daftar_kasus().iter().enumerate() {
        let h = hitung(keranjang);

        println!(
            "{}. {:<30} subtotal {:>9} | voucher {:>8} | ongkir {:>6} | total {:>9} | ditolak {}",
            nomor + 1,
            nama,
            rupiah(h.subtotal),
            rupiah(h.voucher),
            rupiah(h.ongkir),
            rupiah(h.total),
            h.ditolak
        );
    }

    println!();
    println!("Jalankan rustc --test tugas4.rs -o tugas4_test && ./tugas4_test untuk memeriksa.");
}

#[cfg(test)]
mod tests {
    use super::*;

    // Uji tabel: setiap baris berisi nama kasus dan hasil yang diharapkan.
    #[test]
    fn tabel_kasus() {
        let harapan: Vec<(i64, i64, i64, i64, i64)> = vec![
            (70_000, 0, 15_000, 85_000, 0),
            (100_000, 0, 15_000, 115_000, 0),
            (100_001, 10_000, 15_000, 105_001, 0),
            (100_005, 10_001, 15_000, 105_004, 0),
            (149_999, 15_000, 15_000, 149_999, 0),
            (150_000, 15_000, 0, 135_000, 0),
            (200_000, 20_000, 0, 180_000, 0),
            (100_000, 0, 15_000, 115_000, 2),
            (0, 0, 0, 0, 0),
            (0, 0, 0, 0, 1),
        ];

        let daftar = daftar_kasus();
        assert_eq!(daftar.len(), harapan.len(), "jumlah kasus tidak sama");

        for ((nama, keranjang), harap) in daftar.iter().zip(harapan.iter()) {
            let h = hitung(keranjang);
            let dapat = (h.subtotal, h.voucher, h.ongkir, h.total, h.ditolak);
            assert_eq!(dapat, *harap, "kasus: {}", nama);
        }
    }
}
