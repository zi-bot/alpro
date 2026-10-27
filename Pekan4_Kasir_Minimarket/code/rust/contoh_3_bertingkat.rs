// Contoh 3 - diskon bertingkat + pajak (if / else if / else)
// Pembanding untuk 01c_contoh_bertingkat.py. Tidak wajib dikerjakan.
//
// Kompilasi: rustc contoh_3_bertingkat.rs -o contoh_3_bertingkat
// Jalankan : ./contoh_3_bertingkat

fn main() {
    let keranjang: Vec<(i64, i64)> = vec![(60_000, 2), (70_000, 2)];

    let mut subtotal: i64 = 0;
    for &(harga, jumlah) in &keranjang {
        subtotal += harga * jumlah;
    }

    let diskon: f64 = if subtotal < 100_000 {
        0.0
    } else if subtotal < 250_000 {
        0.05 * subtotal as f64
    } else {
        0.10 * subtotal as f64
    };

    let pajak: f64 = 0.11 * (subtotal as f64 - diskon);
    let total: f64 = (subtotal as f64 - diskon) + pajak;

    println!("Subtotal : Rp{}", subtotal);
    println!("Diskon   : Rp{:.0}", diskon);
    println!("Pajak    : Rp{:.0}", pajak);
    println!("Total    : Rp{:.0}", total);
}
