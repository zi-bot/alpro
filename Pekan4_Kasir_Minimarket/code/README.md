# Kode Pekan 4

Tiga bahasa, satu kasus: kasir minimarket. Python adalah bahasa utama.
Rust dan Go hanya pembanding, tidak wajib dikerjakan.

```
code/
├── python/     bahasa utama
├── rust/       pembanding
└── go/         pembanding
```

Setiap bahasa punya dua program yang harus dilengkapi: `latihan` dan
`tugas4`. Pemeriksaannya memakai unit test dengan tabel kasus, bukan
blok pemeriksaan di dalam program.

## Python

```bash
cd python
python3 01a_contoh_satu_item.py
python3 01b_contoh_multi_item.py
python3 01c_contoh_bertingkat.py
python3 01d_contoh_sentinel.py
python3 01e_setengah_jadi.py
python3 latihan_template.py
python3 tugas4_template.py
```

Unit test:

```bash
cd python
python3 -m unittest test_latihan
python3 -m unittest test_tugas4
```

Kunci untuk dosen:

```bash
cd python
python3 solusi/latihan_solusi.py
python3 solusi/tugas4_solusi.py
```

Periksa tipe:

```bash
cd python
pip install mypy
mypy .
```

## Rust

```bash
cd rust
rustc contoh_1_satu_item.rs -o contoh_1_satu_item
./contoh_1_satu_item
```

Unit test:

```bash
cd rust
rustc --test latihan.rs -o latihan_test && ./latihan_test
rustc --test tugas4.rs -o tugas4_test && ./tugas4_test
```

Kunci ada di `rust/solusi/`. Uji kuncinya dengan perintah yang sama,
misalnya:

```bash
rustc --test solusi/tugas4.rs -o tugas4_test && ./tugas4_test
```

## Go

```bash
cd go
go run ./contoh_1_satu_item
go run ./latihan
go run ./tugas4
```

Unit test:

```bash
cd go
go test ./latihan
go test ./tugas4
```

Kunci ada di `go/solusi/`. Uji semuanya dengan:

```bash
cd go
go test ./solusi/...
```

## Pasangan berkas

| Python | Rust | Go |
|---|---|---|
| `01a_contoh_satu_item.py` | `contoh_1_satu_item.rs` | `contoh_1_satu_item` |
| `01b_contoh_multi_item.py` | `contoh_2_multi_item.rs` | `contoh_2_multi_item` |
| `01c_contoh_bertingkat.py` | `contoh_3_bertingkat.rs` | `contoh_3_bertingkat` |
| `01d_contoh_sentinel.py` | `contoh_4_sentinel.rs` | `contoh_4_sentinel` |
| `01e_setengah_jadi.py` | — | — |
| `latihan_template.py` | `latihan.rs` | `latihan` |
| `tugas4_template.py` | `tugas4.rs` | `tugas4` |
| `test_latihan.py` | di dalam `latihan.rs` | `latihan/main_test.go` |
| `test_tugas4.py` | di dalam `tugas4.rs` | `tugas4/main_test.go` |

`01e_setengah_jadi.py` hanya ada di Python: tiga fungsi yang sengaja salah
untuk diperbaiki. Kuncinya di `python/solusi/01e_setengah_jadi_solusi.py`.

Python dan Go menaruh test di berkas terpisah. Rust menaruh test di
berkas yang sama di dalam `mod tests`, karena satu berkas Rust adalah
satu unit kompilasi.

## Bentuk uji tabel

Python memakai `unittest` dengan `subTest`:

```python
KASUS = [
    ("di bawah Rp100.000", [(15_000, 2)], {"subtotal": 30_000, ...}),
]

for nama, keranjang, harapan in KASUS:
    with self.subTest(nama):
        self.assertEqual(kasir.hitung(keranjang), harapan)
```

Rust memakai `#[test]` dan `assert_eq!` per baris tabel:

```rust
let harapan: Vec<(i64, i64, i64, i64, i64)> = vec![...];
for (nomor, (keranjang, harap)) in daftar.iter().zip(harapan.iter()).enumerate() {
    let h = hitung(keranjang);
    assert_eq!((h.subtotal, h.diskon, h.pajak, h.total, h.ditolak), *harap);
}
```

Go memakai `testing` dengan `t.Run` per baris tabel:

```go
for _, k := range kasus {
    t.Run(k.nama, func(t *testing.T) {
        h := hitung(k.keranjang)
        dapat := [5]int{h.subtotal, h.diskon, h.pajak, h.total, h.ditolak}
        if dapat != k.harapan {
            t.Errorf("dapat %v, harapan %v", dapat, k.harapan)
        }
    })
}
```

## Perbandingan tipe

| Konsep | Python | Rust | Go |
|---|---|---|---|
| Deklarasi variabel | `subtotal: int = 0` | `let mut subtotal: i64 = 0;` | `subtotal := 0` |
| Deklarasi eksplisit | anotasi, tidak mengikat | `let mut subtotal: i64 = 0;` | `var subtotal int = 0` |
| Kapan tipe diperiksa | saat `mypy` dijalankan (opsional) | saat kompilasi (wajib) | saat kompilasi (wajib) |
| Bilangan bulat | `int`, presisi tak terbatas | `i64`, 64 bit | `int`, 64 bit di mesin 64 bit |
| Daftar pasangan | `list[tuple[int, int]]` | `Vec<(i64, i64)>` | `[][2]int` |
| Kumpulan field | `dict[str, int]` | `struct Hasil` | `struct Hasil` |
| Fungsi | `def hitung(...) -> dict[str, int]` | `fn hitung(...) -> Hasil` | `func hitung(...) Hasil` |
| Pembulatan voucher | `int(0.10 * subtotal + 0.5)` | `(subtotal * 10 + 50) / 100` | `(subtotal*10 + 50) / 100` |
| Salah tipe | program tetap jalan; mypy menandai | tidak bisa dikompilasi | tidak bisa dikompilasi |

## Tiga hal yang paling terasa berbeda

1. **Anotasi Python tidak mengikat.** `subtotal: int = 0` tetap berjalan
   walau nanti diisi teks. Yang menolak adalah `mypy`, bukan Python.
   Di Rust dan Go, kesalahan yang sama menghentikan kompilasi.

2. **Rust dan Go tidak memakai pecahan untuk uang.** Pada `latihan` dan
   `tugas4`, pembulatan dikerjakan dengan bilangan bulat:
   `(subtotal * 10 + 50) / 100`. Hasilnya sama dengan
   `int(0.10 * subtotal + 0.5)` di Python, tetapi bebas dari galat
   pembulatan. Contoh 1 sampai 3 sengaja masih memakai `f64` dan
   `float64` agar terlihat terjemahan langsung dari Python.

3. **Bentuk data harus dinyatakan lebih awal.** Di Python, `dict` boleh
   berisi kunci apa saja. Di Rust dan Go, bentuk data ditulis sebagai
   `struct` sehingga kompilator tahu setiap field dan tipenya.

## Catatan keluaran

`latihan` dan `tugas4` mencetak struk dengan format ribuan, sama seperti
versi Python. Contoh 1 sampai 4 mencetak angka polos (`Rp120000`, bukan
`Rp120.000`) agar berkas tetap pendek dan fokus pada tipe.
