// 01b — Prosedur dan fungsi: bedanya di nilai balik (versi Rust).
// Jalankan: rustc 01b_prosedur_vs_fungsi.rs -o 01b && ./01b
//
// cetak_garis adalah prosedur: tidak mengembalikan nilai bermakna (tipe unit `()`).
// hitung_subtotal adalah fungsi: mengembalikan hasil perhitungan via nilai balik i64.
// hitung_subtotal_campur melakukan efek samping (cetak) tanpa mengembalikan hasil.

fn cetak_garis(panjang: usize) {
    println!("{}", "-".repeat(panjang));
}

fn hitung_subtotal(harga: i64, jumlah: i64) -> i64 {
    harga * jumlah
}

// Prosedur yang keliru: menghitung sambil mencetak di dalam fungsi
fn hitung_subtotal_campur(harga: i64, jumlah: i64) {
    println!("{}", harga * jumlah);
    // Tidak ada return nilai — fungsi ini mengembalikan ()
}

fn main() {
    cetak_garis(32);
    println!("PROSEDUR — hasilnya tidak bisa dipakai menghitung");
    cetak_garis(32);

    cetak_garis(32);
    println!("FUNGSI — hasilnya bisa dipakai menghitung");
    cetak_garis(32);

    let benar = hitung_subtotal(15_000, 2);
    println!("Memanggil hitung_subtotal_campur — perhatikan angka yang muncul");
    println!("tanpa pernah diminta untuk dicetak:");
    let campur = hitung_subtotal_campur(15_000, 2);

    println!();
    println!("benar  : {} i64", benar);
    println!("campur : {:?} ()", campur);
    println!();
    println!("benar + 5.000  : {}", benar + 5_000);
    println!("campur + 5.000 : [Kompilasi Ditolak] cannot add `i64` to `()`");
    println!("                 Sistem tipe Rust mendeteksi kesalahan ini saat kompilasi!");
}
