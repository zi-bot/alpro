// 02a — Program sebelum dipecah: semuanya di satu alur (versi Rust).
// Jalankan: rustc 02a_sebelum_refaktor.rs -o 02a && ./02a
//
// Program ini BENAR. Jangan diubah dulu — catat keluarannya, karena keluaran
// itulah bukti bahwa hasil refaktor Anda nanti tidak merusak apa pun.

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
    let keranjang: [(&str, &[(i64, i64)]); 3] = [
        ("A", &[(15_000, 2), (5_000, 3), (8_000, 1)]),
        ("B", &[(125_000, 2), (250_000, 1)]),
        ("C", &[(200_000, 1), (100_000, 1)]),
    ];

    let mut total_terbesar = 0;
    let mut nama_terbesar = "";

    for (nama, isi) in keranjang {
        let mut subtotal = 0;
        for &(harga, jumlah) in isi {
            subtotal += harga * jumlah;
        }

        let diskon = if subtotal < 100_000 {
            0
        } else if subtotal < 250_000 {
            subtotal * 5 / 100
        } else if subtotal < 500_000 {
            subtotal * 10 / 100
        } else {
            subtotal * 15 / 100
        };

        let dasar = subtotal - diskon;
        let ppn = dasar * 11 / 100;
        let total = dasar + ppn;

        println!("NOTA KERANJANG {}", nama);
        println!("----------------------------------");
        for &(harga, jumlah) in isi {
            println!(
                "  {:>11} x {} = {:>11}",
                rupiah(harga),
                jumlah,
                rupiah(harga * jumlah)
            );
        }
        println!("----------------------------------");
        println!("  {:<20}{:>14}", "Subtotal", rupiah(subtotal));
        println!("  {:<20}{:>14}", "Diskon", rupiah(diskon));
        println!("  {:<20}{:>14}", "PPN 11%", rupiah(ppn));
        println!("  {:<20}{:>14}", "TOTAL", rupiah(total));
        println!();

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
