# Pekan 4 — Kasir Minimarket

Bahan praktik **Algoritma dan Pemrograman (SI021301)** untuk Pekan 4.
Sub-CPMK-4: menganalisis dan merancang algoritma penyelesaian masalah.
Rubrik: **A1 — Algoritma, flowchart & kuis kasus**. Bobot: 3%.

## Isi folder

| Berkas | Untuk | Keterangan |
|---|---|---|
| `Pekan4_Bahan_Praktik.docx` | Mahasiswa | Materi, contoh bertahap, praktik, tugas akhir sesi |
| `code/README.md` | Mahasiswa | Cara menjalankan dan perbandingan tiga bahasa |
| `code/python/` | Mahasiswa | Bahasa utama: contoh, contoh setengah jadi, dua template, unit test |
| `code/python/solusi/` | Dosen | Kunci kedua template Python |
| `code/rust/` | Mahasiswa | Pembanding, tidak wajib |
| `code/go/` | Mahasiswa | Pembanding, tidak wajib |

## Urutan pengerjaan di kelas

1. Baca Bagian A (materi) sambil dosen menjelaskan.
2. Jalankan `01a`, `01b`, `01c`, lalu `01d`.
3. Isi `latihan_template.py` sampai keempat TODO selesai.
4. Kerjakan `tugas4_template.py` sebagai tugas akhir sesi.

## Cara menjalankan

```bash
cd code/python
python3 01a_contoh_satu_item.py
python3 01b_contoh_multi_item.py
python3 01c_contoh_bertingkat.py
python3 01d_contoh_sentinel.py
python3 01e_setengah_jadi.py
python3 latihan_template.py
python3 tugas4_template.py
```

Menjalankan unit test:

```bash
cd code/python
python3 -m unittest test_latihan
python3 -m unittest test_tugas4
```

Kunci untuk dosen:

```bash
cd code/python
python3 solusi/latihan_solusi.py
python3 solusi/tugas4_solusi.py
```

## Anotasi tipe

Semua kode memakai anotasi tipe pada kemunculan pertama setiap variabel.

```python
subtotal: int = 0
diskon: float = 0.0
keranjang: list[tuple[int, int]] = [(15_000, 2), (5_000, 3)]
```

Anotasi tidak mengubah perilaku program. Python tetap menjalankan
`subtotal: int = 0` sama seperti `subtotal = 0`. Yang berubah adalah
informasi yang dibaca manusia, editor, dan alat pemeriksa tipe.

Aturan yang dipakai di paket ini:

- Rupiah selalu `int`, karena tidak ada pecahan sen.
- Persentase dan hasil bagi memakai `float`.
- Hasil pembulatan disimpan ke variabel `int` terpisah, misalnya
  `diskon_bulat`.
- Variabel yang sama tidak dianotasi dua kali.

Periksa tipe dari folder `code/python`:

```bash
pip install mypy
mypy .
```

Hasil yang diharapkan:

```
Success: no issues found in 13 source files
```

## Cara mengisi TODO

Setiap TODO berupa komentar dan variabel bertipe yang nilainya masih
placeholder, di dalam fungsi `hitung()`:

```python
    # TODO 2 - Voucher
    # Voucher 10% hanya berlaku bila subtotal > 100.000.
    # Bulatkan dengan int(0.10 * subtotal + 0.5).
    voucher: int = 0
```

Tambahkan kode di bawah komentar TODO, jangan mengubah anotasi tipenya:

```python
    if subtotal > 100_000:
        voucher = int(0.10 * subtotal + 0.5)
```

## Cara membaca hasil uji

Unit test memakai tabel kasus. Setiap baris berisi input dan hasil yang
diharapkan, sehingga pesan gagal langsung menunjukkan angkanya:

```
$ python3 -m unittest test_tugas4
F.........
======================================================================
FAIL: test_tabel_kasus (test_tugas4.UjiTugas4.test_tabel_kasus) (nama='Rp100.001')
----------------------------------------------------------------------
AssertionError: {'subtotal': 100001, 'voucher': 10001, ...} !=
                {'subtotal': 100001, 'voucher': 10000, ...}
```

Baris terakhir menunjukkan nilai yang Anda hasilkan dan nilai yang
diharapkan. Perbaiki, lalu jalankan ulang sampai `OK`.

## Aturan yang dinilai

Tugas 4 memakai aturan dari Lampiran RPS:

- Baris valid bila `harga >= 0` dan `jumlah > 0`.
- Voucher 10% hanya bila `subtotal > 100.000`.
- Ongkir Rp15.000, gratis bila `subtotal >= 150.000` sebelum diskon.
- Pembulatan setengah ke atas.
- Keranjang tanpa baris valid: total 0 dan tanpa ongkir.

Perhatikan dua batas yang mudah salah: `>` pada voucher dan `>=` pada ongkir.

## Rust dan Go

Versi Rust dan Go tersedia di `code/rust/` dan `code/go/`. Sifatnya
opsional, hanya untuk membandingkan bahasa bertipe statis. Template
`latihan` dan `tugas4` di kedua bahasa juga belum terisi, dengan unit
test bergaya tabel. Kuncinya ada di folder `solusi` masing-masing.
Rincian perintah ada di `code/README.md`.
