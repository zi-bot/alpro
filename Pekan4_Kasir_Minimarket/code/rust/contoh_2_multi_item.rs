// Contoh 2 - banyak barang, perulangan + penjumlahan
// Pembanding untuk 01b_contoh_multi_item.py. Tidak wajib dikerjakan.
//
// Kompilasi: rustc contoh_2_multi_item.rs -o contoh_2_multi_item
// Jalankan : ./contoh_2_multi_item

fn main() {
    // Vec<(i64, i64)> adalah daftar pasangan (harga, jumlah).
    let keranjang: Vec<(i64, i64)> = vec![(15_000, 2), (5_000, 3), (8_000, 1)];

    let mut subtotal: i64 = 0;
    for &(harga, jumlah) in &keranjang {
        subtotal += harga * jumlah;
    }

    println!("Subtotal : Rp{}", subtotal);
}
