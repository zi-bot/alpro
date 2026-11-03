// test_diskon.rs — uji otomatis untuk hitung_diskon dan hitung_total (versi Rust).
// Jalankan: rustc test_diskon.rs -o test_diskon && ./test_diskon
// Atau:     rustc --test test_diskon.rs -o test_diskon_harness && ./test_diskon_harness

fn validasi_subtotal(subtotal: i64) -> Result<i64, &'static str> {
    if subtotal < 0 {
        Err("subtotal tidak boleh negatif")
    } else {
        Ok(subtotal)
    }
}

fn hitung_diskon(subtotal: i64) -> i64 {
    if subtotal < 100_000 {
        0
    } else if subtotal < 250_000 {
        subtotal * 5 / 100
    } else if subtotal < 500_000 {
        subtotal * 10 / 100
    } else {
        subtotal * 15 / 100
    }
}

fn hitung_total(subtotal: i64, diskon: i64) -> i64 {
    let dasar = subtotal - diskon;
    dasar + dasar * 11 / 100
}

fn test_bawah_batas() {
    assert_eq!(hitung_diskon(0), 0);
    assert_eq!(hitung_diskon(99_999), 0);
}

fn test_tepat_di_batas() {
    assert_eq!(hitung_diskon(100_000), 5_000);
    assert_eq!(hitung_diskon(250_000), 25_000);
    assert_eq!(hitung_diskon(500_000), 75_000);
}

fn test_satu_rupiah_di_atas_batas() {
    assert_eq!(hitung_diskon(100_001), 5_000);
    assert_eq!(hitung_diskon(250_001), 25_000);
    assert_eq!(hitung_diskon(500_001), 75_000);
}

fn test_tepi_atas_setiap_tingkat() {
    assert_eq!(hitung_diskon(249_999), 12_499);
    assert_eq!(hitung_diskon(499_999), 49_999);
}

fn test_pecahan_dibuang_ke_bawah() {
    assert_eq!(hitung_diskon(249_999), 12_499);
}

fn test_diskon_tidak_pernah_melebihi_sepertiga() {
    let mut s = 0;
    while s <= 1_000_000 {
        assert!(hitung_diskon(s) <= s / 3);
        s += 25_000;
    }
}

fn test_subtotal_negatif_ditolak() {
    assert!(validasi_subtotal(-1).is_err());
}

fn test_ppn_setelah_diskon() {
    assert_eq!(hitung_total(100_000, 5_000), 105_450);
}

fn test_tanpa_diskon() {
    assert_eq!(hitung_total(50_000, 0), 55_500);
}

fn test_total_bisa_turun_saat_subtotal_naik() {
    assert!(
        hitung_total(250_000, hitung_diskon(250_000))
            < hitung_total(249_999, hitung_diskon(249_999))
    );
}

fn main() {
    test_bawah_batas();
    test_tepat_di_batas();
    test_satu_rupiah_di_atas_batas();
    test_tepi_atas_setiap_tingkat();
    test_pecahan_dibuang_ke_bawah();
    test_diskon_tidak_pernah_melebihi_sepertiga();
    test_subtotal_negatif_ditolak();
    test_ppn_setelah_diskon();
    test_tanpa_diskon();
    test_total_bisa_turun_saat_subtotal_naik();
    println!("Ran 10 tests ... OK");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn t1() {
        test_bawah_batas();
    }
    #[test]
    fn t2() {
        test_tepat_di_batas();
    }
    #[test]
    fn t3() {
        test_satu_rupiah_di_atas_batas();
    }
    #[test]
    fn t4() {
        test_tepi_atas_setiap_tingkat();
    }
    #[test]
    fn t5() {
        test_pecahan_dibuang_ke_bawah();
    }
    #[test]
    fn t6() {
        test_diskon_tidak_pernah_melebihi_sepertiga();
    }
    #[test]
    fn t7() {
        test_subtotal_negatif_ditolak();
    }
    #[test]
    fn t8() {
        test_ppn_setelah_diskon();
    }
    #[test]
    fn t9() {
        test_tanpa_diskon();
    }
    #[test]
    fn t10() {
        test_total_bisa_turun_saat_subtotal_naik();
    }
}
