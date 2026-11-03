# Pekan 5 — Kode Program: Fungsi dan Prosedur

Repositori kode praktikum **Algoritma dan Pemrograman (SI021301)** untuk Pekan 5.  
Sub-CPMK-5 (CPMK-2): Mampu memahami dan menggunakan prosedur serta fungsi.  
Rubrik: **A2 — Praktikum Pemrograman** (Bobot: 5% nilai akhir).

Fokus praktikum pekan ini adalah **refaktorisasi**: memecah alur program kasir monolitik menjadi fungsi-fungsi kecil yang modular, independen, dan dapat diuji secara terisolasi tanpa mengubah perilaku dan keluarannya.

Seluruh materi tersedia dalam tiga bahasa: **Python** (bahasa pengantar utama praktikum), serta **Rust** dan **Go** (pengayaan perbandingan tipe statis dan model eksekusi).

---

## Struktur Berkas Mahasiswa

Setiap konsep memiliki padanan berkas pada ketiga direktori bahasa (`python/`, `rust/`, dan `go/`):

| Materi / Konsep | Python (`code/python/`) | Rust (`code/rust/`) | Go (`code/go/`) |
|---|---|---|---|
| **Pengenalan Fungsi** | `01a_satu_fungsi.py` | `01a_satu_fungsi.rs` | `01a_satu_fungsi/main.go` |
| **Prosedur vs Fungsi** | `01b_prosedur_vs_fungsi.py` | `01b_prosedur_vs_fungsi.rs` | `01b_prosedur_vs_fungsi/main.go` |
| **Diskon Bertingkat** | `01c_diskon_bertingkat.py` | `01c_diskon_bertingkat.rs` | `01c_diskon_bertingkat/main.go` |
| **Komposisi 5 Fungsi** | `01d_komposisi.py` | `01d_komposisi.rs` | `01d_komposisi/main.go` |
| **Analisis Cacat Semantik** | `01e_setengah_jadi.py` | `01e_setengah_jadi.rs` | `01e_setengah_jadi/main.go` |
| **Stack & Mutabilitas** | `01f_stack_mutabilitas.py` | `01f_stack_mutabilitas.rs` | `01f_stack_mutabilitas/main.go` |
| **Acuan Sebelum Refaktor** | `02a_sebelum_refaktor.py` | `02a_sebelum_refaktor.rs` | `02a_sebelum_refaktor/main.go` |
| **Template Praktik** | `latihan_template.py` | `latihan_template.rs` | `latihan_template/main.go` |
| **Template Tugas 5** | `tugas5_template.py` | `tugas5_template.rs` | `tugas5_template/main.go` |
| **Unit Test (10 Kasus)** | `test_diskon.py` | `test_diskon.rs` | `test_diskon/main.go` |
| **Pembanding 10 Kasus** | *(via unit test)* | `main.rs` | `main.go` |

> **Catatan Struktur Go:** Dalam Go, setiap program executable ditempatkan di dalam foldernya masing-masing dengan modul `pekan5_fungsi` (`go.mod`) agar IDE (seperti VS Code / GoLand) tidak mengalami galat *redeclaration error* antar-paket (`package main`).

---

## 1. Menjalankan Kode Python (Bahasa Utama)

Pastikan terminal aktif di direktori `code/python/`:

```bash
cd code/python
```

### Menjalankan Contoh & Konsep
```bash
python3 01a_satu_fungsi.py
python3 01b_prosedur_vs_fungsi.py
python3 01c_diskon_bertingkat.py
python3 01d_komposisi.py
python3 01e_setengah_jadi.py
python3 01f_stack_mutabilitas.py
```

### Menjalankan Template Praktik & Tugas 5
```bash
python3 latihan_template.py      # Memicu NotImplementedError sebelum TODO diisi
python3 tugas5_template.py       # Memicu NotImplementedError sebelum TODO diisi
python3 -m unittest test_diskon  # Pengujian 10 kasus uji otomatis (TDD)
```

### Membuktikan Keberhasilan Refaktor (`diff`)
Setelah Anda mengisi seluruh TODO pada `latihan_template.py`, buktikan bahwa hasil keluaran tidak berubah sama sekali dari program acuan:

```bash
python3 02a_sebelum_refaktor.py > acuan.txt
python3 latihan_template.py > hasil_saya.txt
diff acuan.txt hasil_saya.txt
echo $?                          # 0 menandakan keluaran identik 100%
```

---

## 2. Menjalankan Kode Rust (Pengayaan)

Pastikan terminal aktif di direktori `code/rust/`:

```bash
cd code/rust
```

### Kompilasi dan Eksekusi Berkas
```bash
# Contoh dan demonstrasi
rustc 01a_satu_fungsi.rs -o /tmp/r_01a && /tmp/r_01a
rustc 01b_prosedur_vs_fungsi.rs -o /tmp/r_01b && /tmp/r_01b
rustc 01c_diskon_bertingkat.rs -o /tmp/r_01c && /tmp/r_01c
rustc 01d_komposisi.rs -o /tmp/r_01d && /tmp/r_01d
rustc 01e_setengah_jadi.rs -o /tmp/r_01e && /tmp/r_01e
rustc 01f_stack_mutabilitas.rs -o /tmp/r_01f && /tmp/r_01f
rustc 02a_sebelum_refaktor.rs -o /tmp/r_02a && /tmp/r_02a

# Template praktik & tugas (memicu panic todo! sebelum diisi)
rustc latihan_template.rs -o /tmp/r_latihan && /tmp/r_latihan
rustc tugas5_template.rs -o /tmp/r_tugas5 && /tmp/r_tugas5

# Menjalankan pengujian 10 unit test otomatis
rustc test_diskon.rs -o /tmp/r_test && /tmp/r_test
# Atau menggunakan test harness bawaan rustc:
rustc --test test_diskon.rs -o /tmp/r_harness && /tmp/r_harness

# Menjalankan pembanding 10 kasus uji
rustc main.rs -o /tmp/r_main && /tmp/r_main
```

### Karakteristik Penting di Rust
- **Sistem Tipe & Nilai Balik**: Tipe parameter dan nilai balik wajib dinyatakan eksplisit (`(subtotal: i64) -> i64`).
- **Implicit Return**: Ekspresi di baris terakhir fungsi tanpa titik-koma otomatis menjadi nilai balik.
- **Unit Type `()`**: Prosedur tidak mengembalikan nilai kosong, melainkan tipe unit `()`. Rust menolak operasi aritmatika antara unit `()` dan integer saat tahap kompilasi.
- **Borrowing & Mutability**: Default variabel bersifat *immutable*; mutasi memori memerlukan izin eksplisit melalui `&mut`.

---

## 3. Menjalankan Kode Go (Pengayaan)

Pastikan terminal aktif di direktori `code/go/`:

```bash
cd code/go
```

### Eksekusi Berkas dengan `go run`
Setiap paket dapat dijalankan langsung dengan menunjuk foldernya:

```bash
# Contoh dan demonstrasi
go run ./01a_satu_fungsi
go run ./01b_prosedur_vs_fungsi
go run ./01c_diskon_bertingkat
go run ./01d_komposisi
go run ./01e_setengah_jadi
go run ./01f_stack_mutabilitas
go run ./02a_sebelum_refaktor

# Template praktik & tugas (memicu panic sebelum diisi)
go run ./latihan_template
go run ./tugas5_template

# Menjalankan pengujian 10 unit test otomatis
go run ./test_diskon

# Menjalankan pembanding 10 kasus uji
go run main.go
```

### Karakteristik Penting di Go
- **Struktur Paket**: Setiap program executable memiliki direktori sendiri untuk memisahkan scope `package main`, sehingga bebas dari konflik simbol di IDE.
- **Sintaks Parameter**: Nama variabel ditulis mendahului tipe datanya (`subtotal int`).
- **Explicit Return**: Kata kunci `return` bersifat wajib di setiap cabang alur percabangan.
- **Prosedur Murni**: Fungsi tanpa tipe nilai balik langsung diakhiri tanpa return. Pemanggilan prosedur di dalam ekspresi aritmatika akan ditolak oleh kompilator.
- **Pass-by-Value & Pointer**: Tipe primitif disalin saat dipassing ke fungsi; mutasi objek asal memerlukan pointer (`*[]int`).

---

## Aturan Logika Diskon Bertingkat

Perhitungan diskon menggunakan aturan batas bawah **inklusif** dan pembulatan ke bawah dengan pembagian bulat integer (`// 100` di Python, integer division `/ 100` di Rust & Go):

| Rentang Subtotal Belanja | Tarif Diskon | Rumus Perhitungan |
|---|---|---|
| Di bawah Rp100.000 | 0% | `0` |
| Rp100.000 – Rp249.999 | 5% | `subtotal * 5 / 100` |
| Rp250.000 – Rp499.999 | 10% | `subtotal * 10 / 100` |
| Rp500.000 atau lebih | 15% | `subtotal * 15 / 100` |

### Dua Fenomena Uji Batas Kritis
1. **Pembuangan Pecahan**:
   - Subtotal Rp249.999 × 5% = 12.499,95 → diskon bernilai **Rp12.499** (bukan Rp12.500). Jangan gunakan fungsi pembulatan desimal `round()`.
2. **Lonjakan Tarif Perbatasan**:
   - Belanja Rp249.999 menghasilkan total bayar Rp263.625 (diskon 5%).
   - Belanja Rp250.000 (tambah belanja Rp1) menghasilkan total bayar Rp249.750 (diskon melompat ke 10%), tagihan menjadi **Rp13.875 lebih murah**.

---

## Kriteria Penilaian Tugas 5 (Rubrik A2)

Tugas 5 adalah tugas mandiri individu dengan bobot **5% nilai akhir**:

| Dimensi Penilaian | Bobot | Target Capaian |
|---|---|---|
| **Kebenaran Program** | 35% | Seluruh 10 kasus uji batas (K-01 s.d. K-10) lolos pengujian tanpa galat. |
| **Penerapan Konsep** | 30% | Memisahkan validasi, fungsi diskon, PPN, dan total (SRP; tanpa `print` di fungsi hitung). |
| **Pengujian & Analisis** | 20% | Hasil pengujian 10 kasus dicatat dan fenomena batas dianalisis. |
| **Penjelasan Individu** | 15% | Mampu menelusuri alasan dan alur keputusan logika per baris kode. |

**Format Pengumpulan**: Berkas program Python dan laporan pengujian pelaksanaan tugas.
