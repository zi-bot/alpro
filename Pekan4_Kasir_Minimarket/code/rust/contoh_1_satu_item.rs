// Contoh 1 - satu barang, satu keputusan (if / else)
// Pembanding untuk 01a_contoh_satu_item.py. Tidak wajib dikerjakan.
//
// Kompilasi: rustc contoh_1_satu_item.rs -o contoh_1_satu_item
// Jalankan : ./contoh_1_satu_item

fn main() {
    // Tipe ditulis eksplisit. Kompilator menolak nilai yang tipenya salah.
    let harga: i64 = 120_000;
    let jumlah: i64 = 1;
    let subtotal: i64 = harga * jumlah;

    // Diskon bertipe f64 karena berasal dari persentase.
    let diskon: f64 = if subtotal > 100_000 {
        0.10 * subtotal as f64
    } else {
        0.0
    };

    println!("Subtotal : Rp{}", subtotal);
    println!("Diskon   : Rp{:.0}", diskon);
}
