// Contoh 4 - perulangan while dengan penanda berhenti
// Pembanding untuk 01d_contoh_sentinel.py. Tidak wajib dikerjakan.
//
// Kompilasi: rustc contoh_4_sentinel.rs -o contoh_4_sentinel
// Jalankan : ./contoh_4_sentinel

fn main() {
    // Pasangan (harga, jumlah), diakhiri harga -1 sebagai penanda berhenti.
    let masukan: [i64; 7] = [15_000, 2, 5_000, 3, 8_000, 1, -1];

    let mut subtotal: i64 = 0;
    let mut i: usize = 0;

    while masukan[i] >= 0 {
        let harga: i64 = masukan[i];
        let jumlah: i64 = masukan[i + 1];
        subtotal += harga * jumlah;
        i += 2;
    }

    println!("Subtotal : Rp{}", subtotal);
}
